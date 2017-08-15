package registry

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/docker/docker/utils"
)

const CONFIGFILE = ".dockercfg"

const INDEXSERVER = "https://index.docker.io/v1/"

var (
	ErrConfigFileMissing = errors.New("The Auth config file is missing")
)

type AuthConfig struct{
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty`
	Auth string `json:"auth"`
	Email string "json:email"
	ServerAddress string `json:"serveraddress,omitempty`
}

type ConfigFile {
	Configs map[string]AuthConfig `json:"configs,omitempty"`
	rootPath string
}

func IndexServerAddress() string{
	return INDEXSERVER
}