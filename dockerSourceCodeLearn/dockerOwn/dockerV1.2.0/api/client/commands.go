package client

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"text/tabwriter"
	"text/template"
	"time"

	"github.com/docker/docker/api"
	"github.com/docker/docker/archive"
	"github.com/docker/docker/dockerverison"
	"github.com/docker/docker/engine"
	"github.com/docker/docker/nat"
	"github.com/docker/docker/opts"
	"github.com/docker/docker/pkg/log"
	"github.com/docker/docker/pkg/parsers"
	"github.com/docker/docker/pkg/parsers/filters"
	"github.com/docker/docker/pkg/signal"
	"github.com/docker/docker/pkg/term"
	"github.com/docker/docker/pkg/units"
	"github.com/docker/docker/pkg/registry"
	"github.com/docker/docker/pkg/runconfig"
	"github.com/docker/docker/utils"

)

const (
	tarHeaderSize = 512
)

func (cli *DockerCli) CmdHelp(args ...string) error{
	if len(args) > 0 {
		method,exits := cli.getMethod(args[0])
		if ! exits{
			fmt.Fprintf(cli.err,"Error:Command not found :%s\n",args[0])
		} else{
			method("--help")
			return nil
		}
	}
	help:= fmt.Sprintf("Usage : docker [OPTIONS] COMMAND [arg...] \n -H=[unix://%s]: tcp://host:port to bind/connect to or unix://path/to/socket to use\n\nA self-sufficlient runtime for linux containers.\n\nCommands:\n",api.DEFAULTUNIXSOCKET)
	// just range all the command help 
	for _,commands := range [][]string{
		{"attach", "Attach to a running container"},
		{"build", "Build an image from a Dockerfile"},
		{"commit", "Create a new image from a container's changes"},
		{"cp", "Copy files/folders from a container's filesystem to the host path"},
		{"diff", "Inspect changes on a container's filesystem"},
		{"events", "Get real time events from the server"},
		{"export", "Stream the contents of a container as a tar archive"},
		{"history", "Show the history of an image"},
		{"images", "List images"},
		{"import", "Create a new filesystem image from the contents of a tarball"},
		{"info", "Display system-wide information"},
		{"inspect", "Return low-level information on a container"},
		{"kill", "Kill a running container"},
		{"load", "Load an image from a tar archive"},
		{"login", "Register or log in to a Docker registry server"},
		{"logout", "Log out from a Docker registry server"},
		{"logs", "Fetch the logs of a container"},
		{"port", "Lookup the public-facing port that is NAT-ed to PRIVATE_PORT"},
		{"pause", "Pause all processes within a container"},
		{"ps", "List containers"},
		{"pull", "Pull an image or a repository from a Docker registry server"},
		{"push", "Push an image or a repository to a Docker registry server"},
		{"restart", "Restart a running container"},
		{"rm", "Remove one or more containers"},
		{"rmi", "Remove one or more images"},
		{"run", "Run a command in a new container"},
		{"save", "Save an image to a tar archive"},
		{"search", "Search for an image on the Docker Hub"},
		{"start", "Start a stopped container"},
		{"stop", "Stop a running container"},
		{"tag", "Tag an image into a repository"},
		{"top", "Lookup the running processes of a container"},
		{"unpause", "Unpause a paused container"},
		{"version", "Show the Docker version information"},
		{"wait", "Block until a container stops, then print its exit code"},
		}{
			help += fmt.Sprintf("	%-10.10s\n",command[0],command[1])
		}
		fmt.Fprintf(cli.err,"%s\n",help)
		return nil
}

func (cli *DockerCli) CmdPull(args ...string) error{
	cmd := cli.Subcmd("pull","NAME[:TAG]","Pull an image or a repository from the registry")
	tag := cmd.String([]string{"#t","#-tag"},"","Download tagged image in a repository")
	// parse the param again
	if err := cmd.Parse(args);err!= nil{
		return nil
	}
	if cmd.NArg() != 1{
		cmd.Usage()
		return nil
	}	
	var (
		v = url.Values{}
		remote = cmd.Arg(0)
		)
	v.Set("fromImage",remote)
	if *tag == ""{
		v.Set("tag",*tag)
	}
	remote,_=parsers.ParseRepositoryTag(remote)
	hostname,_,err := registry.ResolveRepositoryName(remote)

	if err != nil{
		return err
	}
	cli.LoadConfigFile()
}
