package opts

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/docker/docker/api"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/pkg/parsers"
)


type ListOpts struct{
	values *[]string
	validator ValidatorFctType
}

type ValidatorFctType func(val string) (string,error)

func newListOptsRef(values *[]string,validator ValidatorFctType) *ListOpts{
	return &ListOpts{
		values: values,
		validator:validator,
	}
}

func HostListVar(values *[]string,name []string,usage string){
	flag.Var(newListOptsRef(values,api.ValidateHost),name,usage)
}