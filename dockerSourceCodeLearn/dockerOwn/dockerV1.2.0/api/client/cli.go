package client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/template"

	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/docker/pkg/term"
	"github.com/docker/docker/registry"
)


type DockerCli struct{
	proto string
	addr string
	configFile *registry.ConfigFile
	io io.ReadCloser
	out io.Writer
	err io.Writer
	isTerminal bool
	terminalFd uintptr
	tlsConfig *tls.Config
	scheme string
}

func NewDockerCli(io io.ReadCloser,out,err io.Writer,proto,addr string,tlsConfig *tls.Config) *DockerCli{
	var (
		isTerminal = false
		terminalFd uintptr
		scheme = "http"
		)
	if tlsConfig != nil{
		scheme = "https"
	}	

	if in != nil{
		if file,ok := out.(*os.File);ok{
			terminalFd = file.Fd()
			isTerminal = term.IsTerminal(terminalFd)
		}
	}
	if err == nil{
		err = out
	}
	return &DockerCli{
		proto:	proto,
		addr:	addr,
		in:	in,
		out:	out,
		err:	err,
		isTerminal:	isTerminal,
		terminalFd:terminalFd,
		tlsConfig: tlsConfig,
		scheme:scheme,
	}
}

func (cli *DockerCli) Cmd(args ...string) error{
	if len(args) > 0{
		method,exists := cli.getMethod(args[0])
		if ! exists {
			fmt.Println("Error: Command not found:",args[0])
			return cli.CmdHelp(args[1:]...)
		}
		return method(args[1:]...)
	}
	return cli.CmdHelp(args...)
}

func (cli *DockerCli) getMethod(name string) (func(...string) error,bool){
	if len(name) == 0{
		return nil,false
	}
	methodName := "Cmd"+strings.ToUpper(name[:1])+strings.ToLower(name[:1])
	method:=reflect.ValueOf(cli).MethodByName(methodName)
	if !method.IsValid(){
		return nil,false
	}
	return method.Interface().(func(...string) error),true
}

func (cli *DockerCli) SubCmd(name,signature,description string) *flag.FlagSet{
	flags := flag.NewFlagSet(name,flag.ContinueOnError)
	flags.Usage = func(){
		fmt.Fprintf(cli.err,"\nUsage: docker %s %s \n\n%s\n\n",name,signature,description)
		flags.PrintDefaults()
		os.Exit(2)
	}
	return flags
}

func (cli *DockerCli) LoadConfigFile() (err error){
	cli.configFile,err = registry.LoadConfigFile(os.Getenv("HOME"))
	if err != nil{
		fmt.Fprintf(cli.err,"WARNING: %s\n",err)
	}
	return err
}