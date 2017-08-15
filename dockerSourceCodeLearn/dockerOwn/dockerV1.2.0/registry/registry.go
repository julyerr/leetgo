package registry

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/docker/docker/pkg/log"
	"github.com/docker/docker/utils"
)

var (
	ErrAlreadyExits	=errors.New("Image already exists")
	ErrInvalidRepositoryName = errors.New("Invalid repository name (ex:\"registry.domain.tld/myrepos\")")
	ErrLoginRequired = errors.New("Authentication is required.")	
)

type TimeoutType uint32

const (
	NoTimeout	TimeoutType = iota
	ReceiveTimeout
	ConnectTimeout
)

func validateRepositoryName(repositoryName string) error{
	var (
		namespace string
		name 	string
		)
	nameParts := strings.SplitN(repositoryName,"/",2)
	if len(nameParts) < 2 {
		namespace = "library"
		name=nameParts[0]
	}else{
		namespace = nameParts[0]
		name = nameParts[1]
	}
	validNameSpace := regexp.MustCompile(`^([a-z0-9_]{4,30}$)`)
	if !validNameSpace.MatchString(namespace) {
		return fmt.Errorf("Invalid namespace name (%s),only [a-z0-9_] are allowed, size between 4 and 30",namespace)
	}
	validRepo := regexp.MustCompile(`^([a-z0-9_.]+)$`)
	if !validRepo.MatchString(name) {
		return fmt.Errorf("Invalid repository name (%s) , only [a-z0-9_.] are allowed",name)
	}
	return nil
}

func ResolveRepositoryName(reposName string) (string,string,error){
	if strings.Contains(reposName,"://"){
		return "","",ErrInvalidRepositoryName
	}
	nameParts := strings.SplitN(reposName,"/",2)
	if len(nameParts) == 1 || (!strings.Contains(nameParts[0],".")) && !strings.Contains(nameParts[0],":") &&
	nameParts[0] != "localhost") {
	// This is a dcoekr Inde repos (ex:samalba/hipache or ubuntu) 
		err := validateRepositoryName(reposName)
		return IndexServerAddress(),reposName,err
	}
	hostname := nameParts[0]
	reposName = nameParts[1]
	if strings.Contains(hostname,"index.docker.io"){
		return "","",fmt.Errorf("Invalid repository name, try \"%s\" instead",respoName)
	}
	if err := validateRepositoryName(reposName);err != nil{
		return "","",err
	}
	return hostname,reposName,nil
}

