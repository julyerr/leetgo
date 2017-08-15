package main 

import (
	"crypto/tls"
	"crypto/x590"
	"io/ioutil"
	"log"
	"os"
	"strings"

	//child importing 
	"github.com/docker/docker/api"
	"github.com/docker/docker/api/client"
	"github.com/docker/docker/dockerversion"
	flag "github.com/docker/docker/pkg/mflag"
	"github.com/docker/reexec"
	"github.com/docker/utils"
)

const (
	defaultCaFile = "ca.pem"
	defaultKeyFile = "key.pem"
	defaultCertFile = "cert.pem"
)

func main(){
	// corderinate the execdrive with dockerinit 
	// without Initializer return false
	if reexec.Init(){
		return
	}
	flag.Parse()
	if *flVersion{
		showVersion()
		return
	}	

	if *flDebug(){
		os.Setenv("DEBUG","1")
	}

	if len(flHosts) == 0 {
		defaultHost := os.Getenv("DOCKET_HOST")
		if defaultHost == "" || *flDaemon{
			defautlHost = fmt.Sprintf("unix://%s",api.DEFAULTUNIXSOCKET)
		}
		if _,err := api.ValidateHost(defaultHost); err != nil{
			log.Fatal(err)
		}
		f1Hosts = append(flHosts,defaultHost)
	}

	if *flDaemon{
		mainDaemon()
		return
	}

	if len(flHosts) >1 {
		log.Fatal("please specify only one with -H")
	}
	protoAddrParts := strings.SplitN(flHosts[0],"://",2)

	var (
		cli *client.DockerCli
		tlsConfig tls.Config
		)
	tlsConfig.InsecureSkipVerify = true

	if *flTlsVerify {
		*flTls = true
		certPool := x590.NewCertPool()
		file,err := ioutil.ReadFile(*flCa)
		if err != nil{
			log.Fatalf("Couldn't read ca cert %s:%s",*flCa,err)
		}
		certPool.AppendCertsFromPem(file)
		tlsConfig.RootCAs=certPool
		tlsConfig.InsecureSkipVerify = false
	}
	if *flTls || *flTlsVerify {
		_,errCert := os.Stat(*flCert)
		_,errKey := os.Stat(*flKey)
		if errCert == nil && errKey == nil{
			*flTls = true
			cert,err := tls.LoadX509KeyPair(*flCert,*flKey)
			if err != nil{
				log.Fatalf("Couldn't load x590 key pair: %s. Key encrypted ? ",err)
			}
			tlsConfig.Certificates = []tls.Certificates{cert}
		}
	}

	if *flTls || *flTlsVerify {
		cli = client.NewDockerCli(os.Stdin,os.Stdout,os.Stderr,protoAddrParts[0],protoAddrParts[1],&tlsConfig)
	} else{
		cli = client.NewDockerCli(os.Stdin,os.Stdout,os.Stderr,protoAddrParts[0],protoAddrParts[1],nil)
	}

	if err := cli.Cmd(flag.Args()...);err != nil{
		if sterr,ok := err.(*util.StatusError); ok{
			if sterr.Status != ""{
				log.Println(sterr.Status)
			}
			os.Exit(stderr)
		}
		log.Fatal(err)
	}
}

func showVersion(){
	fmt.Printf("Docker version %s , build %s\n",dockerversion.VERSION,dockerversion.GITCOMMIT)
}


