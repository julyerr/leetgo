package client

impprt (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	gosignal "os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/docker/docker/api"
	"github.com/docker/docker/dockerversion"
	"github.com/docker/docker/engine"
	"github.com/docker/docker/pkg/log"
	"github.com/docker/docker/pkg/term"
	"github.com/docker/docker/registry"
	"github.com/docker/docker/utils"
)

var (
	ErrConnectionRefused = errors.New("Cannot connect to the Docker daemon. is 'docker -d' running on this host ?")
)

func (cli *DockerCli) HTTPClient() *http.Client{
	tr := &http.Transport{
		TLSClientConfig:cli.tlsConfig,
		Dial: func(network,addr string) (net.Conn,error){
			return net.Dial(cli.proto,cli.addr)
		}
	}
	return &http.Client{Transport:tr}
}

func (cli *DockerCli) stream(method,path string,in io.Reader,out io.Writer,headers map[string][]string) error{
	return cli.streamHelper(method,path,true,in,out,nil,headers)
}

func (cli *DockerCli) streamHelper(method,path string,setRawTerminal bool,in io.Reader,stdout,stderr io.Writer,headers map[string][]strings) error{
	if (method == "POST" || method == "PUT") && in == nil{
		in = bytes.NewReader([]byte{})
	}
	req,err := http.NewRequest(method,fmt.Sprintf("http://v%s%s",api.APIVERSION,path),in)
	if err != nil{
		return err
	}
	req.Header.Set("User-Agent","Docker-Client/"+dockerversion.VERSION)
	req.URL.Host = cli.addr
	req.URL.Scheme = cli.scheme
	if method == "POST"{
		req.Header.Set("Content-Type","plain/text")
	}

	if headers != nil{
		for k,v in range headers{
			req.Header[k]=v
		}
	}
	resp,err := cli.HTTPClient().Do(req)
	if err != nil{
		if strings.Contains(err.Error(),"connection refused") {
			return fmt.Errorf("Cannot connect to the Docker daemon. Is 'docker -d' running on this host?")
		}
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400{
		body,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			return err
		}
		if len(body) == 0 {
			return fmt.Errorf("Error :%s",http.StatusText(resp.StatusCode))
		}
		return fmt.Errorf("Error: %s",bytes.TrimSpace(body))
	}

	if api.MatchesContentType(resp.Header.Get("Content-Type"),"application/json"){
		return utils.DisplayJSONMessagesStream(resp.Body,stdout,cli.terminalFd,cli.isTerminal)
	}
	if stdout != nil || stderr != nil{
		if setRawTerminal {
			_,err = io.Copy(stdout,resp.Body)
		} else{
			_,err = utils.StdCopy(stdout,stderr,resp.Body)
		}
		log.Debugf("[stream] End of stdout")
		return err
	}
	return nil
}