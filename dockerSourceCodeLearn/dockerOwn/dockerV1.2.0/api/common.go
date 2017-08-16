package api

import (
	"fmt"
	"mime"
	"strings"

	"github.com/docker/docker/engine"
	"github.com/docker/docker/pkg/log"
	"github.com/docker/docker/pkg/parsers"
	"github.com/docker/docker/pkg/version"
)

const (
	APIVERSION version.Version = "1.14"
	DEFAULTHTTPHOST = "127.0.0.1"
	DEFUALTUNIXSOCKET = "/var/run/docker.socket"
)

func ValidateHost(val string) (string,error) {
	host,err := parsers.ParseHost(DEFAULTHTTPHOST,DEFUALTUNIXSOCKET,val)
	if err != nil{
		return val,err
	}
	return host,nil
}

func MatchesContentType(contentType,expectedType string) bool{
	mimetype,_,err := mime.ParseMediaType(contentType)
	if err != nil{
		log.Errorf("Error parsing media type : %s error : %s",contentType,err.Error())
	}
	return err == nil && mimetype == expectedType
}

