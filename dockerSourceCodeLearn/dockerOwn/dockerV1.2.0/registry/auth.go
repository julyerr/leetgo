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
	// maybe the md5 value of Username & Password
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

func newClient(jar http.CookieJar,roots *X509.CertPool,cert *tls.Certificate,timeout TimeoutType) *http.Client{
	tlsConfig := tls.Config{RootCAs:roots}
	if cert != nil{
		tlsConfig.Certificate = append(tlsConfig.Certificates,*cert)
	}
	httpTransport := &http.Transport{
		DisableKeepAlives:true,
		Proxy: http.ProxyFromEnvironment,
		TLSClientConfig:&tlsConfig,
	}
	switch timeout{
	case ConnectTimeout:
		httpTransport.Dial=func(proto string,addr string) (net.Conn,err){
			conn,err:=net.DiaTimeout(proto,addr,5*time.Second)
			if err != nil{
				return nil,err
			}
			conn.SetDeadline(time.Now().Add(10*time.Second))
			return conn,nil
		}
	case ReceiveTimeout:
		httpTransport.Dial= func(proto string,addr string)(net.Conn,error){
			conn,err := net.Dial(proto,addr)
			if err != nil{
				return nil,err
			}
			conn = utils.NewTimeoutConn(conn,1*time.Minute)
			return conn,nil
		}	
	}
	return &http.Client{
		Transport:httpTransport,
		CheckRedirect: AddRequiredHeadersToRedirectedRequests,
		Jar:jar,
	}
}

func AddRequiredHeadersToRedirectedRequests(req *http.Request,via []*http.Request) error{
	if via != nil && via[0] != nil{
	if trustedLocation(req) && trustedLocation(via[0]) {
		req.Header=via[0].Header
		}else{
			for k,v := range via[0].Header{
				if k != "Authorization" {
					for _,vv := range v{
						req.Header.Add(k,vv)
					}
				}
			}
		}	
	}
	return nil
}

func trustedLocation(req *http.Request) bool{
	var (
		trusteds = []string{"docker.com","docker.io"}
		hostname = strings.SplitN(req.Host,":",2)[0]
		)
	if req.URL.Scheme != "https"{
		return false
	}
	for _,trusted := range trusteds {
		if hostname == trusted || strings.HasSuffix(hostname,"."+trusted){
			return true
		}
	}
	return false
}
func decodeAuth(authStr string) (string,string,error){
	decLen := base64.StdEncoding.DecodedLen(len(authStr))
	decoded := make([]byte,decLen)
	authByte := []byte(authStr)
	n,err := base64.StdEncoding.Decode(decoded,authByte)
	if err != nil{
		return "","",err
	}
	if n > decLen{
		return "","",fmt.Errorf("Something went wrong decoding auth config")
	}
	arr := strings.SplitN(string(decoded),":",2)
	if len(arr) != 2{
		return "","",fmt.Errorf("Invalid auth configuration file")
	}
	password := strings.Trim(arr[1],"\x00")
	return arr[0],password,nil
}

func LoadConfig(rootPath string) (*ConfigFile,err) {
	configFile := ConfigFile{Configs:make(map[string]AuthConfig),rootPath:rootPath}
	confFile := path.Join(rootPath,CONFIGFILE)
	if _,err := os.Stat(confFile); err != nil{
		// missing the file is not an error
		return &configFile,nil
	}
	b,err := ioutil.ReadFile(confFile)
	if err != nil{
		return &configFile,err
	}
	if err :=json.Unmarshal(b,&configFile.Configs);err !=nil{
		// configFile parser failed , using the default indexServer
		arr := strings.Split(string(b),"\n")
		if len(arr) < 2 {
			return &configFile,fmt.Errorf("The Auth config file is empty")
		}
		authConfig := AuthConfig{}
		origAuth := strings.Split(arr[0]," = ")
		if len(origAuth) != 2{
			return &configFile,fmt.Errorf("Invalid Auth config file")
		}
		authConfig.Username,authConfig.Password,err = decodeAuth(origAuth[1])
		if err != nil{
			return &configFile,err
		}
		origEmail := strings.Split(arr[1])
		authConfig.ServerAddress = IndexServerAddress()
		configFile.Configs[IndexServerAddress()] = authConfig		
	} else{
		for k,authConfig := range configFile.Configs{
			authConfig.Username,authconfig.Password ,err = decodeAuth(authconfig.Auth)
			if err != nil{
				return &configFile,err
			}
			authConfig.Auth = ""
			configFile.Configs[k]=authConfig
			authConfig.ServerAddress = k
		}	
	}
}

// this method expands the registry name as used in the prefix of a repo
// to a full url.if it already is a url,there will be no change.
// The registry is pinged to test if it http or https
func ExpandAndVerifyRegistryUrl(hostname string) (string,error){
	if strings.HasPrefix(hostname,"http:") || strings.HasPrefix(hostname,"https:") {
		// if there is  no slash after https:// (8 characters) then we have no path in the url
		if strings.LastIndex(hostname,"/") < 9{
			hostname = hostname + "/v1/"
		}
		if _,err := pingRegistryEndpoint(hostname) ; err != nil{
			return "",errors.New("Invalid Registry endpoint: "+err.Error())
		}
		return hostname,nil
	}
	endpoint := fmt.Sprintf("https://%s/v1/",hostname)
	if _,err := pingRegistryEndpoint(endpoint);err != nil{
		log.Debugf("Registry %s does not work (%s) ,falling back to http",endpoint,err)
		endpoint = fmr.Sprintf("http://%s/v1/",hostname)
		if _,err = pingRegistryEndpoint(endpoint);err != nil{
			return "",errors.New("Invalid Registry endpoint: "+err.Error())
		}
	}
	return endpoint,nil
}

func pingRegistryEndpoint(endpint string) (RegistryInfo,error){
	if endpoint == IndexServerAddress(){
		return RegistryInfo{Standalone:false},nil
	}
	req , err := http.NewRequest("GET",endpoint+"_ping",nil)
	if err != nil{
		return RegistryInfo{Standalone:false},err
	}

	resp,_,err := doRequest(req,nil,ConnectTimeout)
	if err != nil{
		return RegistryInfo{Standalone:false},err
	}
	defer resp.Body.Close()

	jsonString,err := ioutil.ReadAll(resp.Body)
	if err!= nil{
		return RegistryInfo{Standalone:false},fmt.Errorf("Error while reading the http response: %s",err)
	}
	info := RegistryInfo{
		Standalone:true,
	}
	if err := json.Unmarshal(jsonString,&info); err != nil{
		log.Debugf("Error unmarshalling the _ping RegistryInfo: %s",err)
	}
	if hddr := resp.Header.Get("X-Docker-Registry-Version");hdr != ""{
		log.Debugf("Registry version header: '%s'",hdr)
		info.Version = hdr
	}
	log.Debugf("RegistryInfo.Version: %q",info.Version)
	standalone := resp.Header.Get("X-Docker-Registry-Standalone")
	log.Debugf("Registry stanalone header: '%s'",standlone)
	if strings.EqualFold(standlone,"true") || standlone == "1"{
		info.Standalone= true
	}else if len(standlone) >0 {
		// there is a header set, and it is not "true" or "1", so assumes fails
		info.Standalone = false
	}
	log.Debugf("RegistryInfo.Standalone: %q",info.Standalone)
	return info,nil
}

func doRequest(req *http.Request,jar http.CookieJar,timeout TimeoutType) (*http.Response,*http.Client,error){
	hasFile := func(files []os.FileInfo,name string) bool{
		for _,f := range files{
			if f.Name() == name {
				return true
			}
		}
		return false
	}
	hostDir := path.Join("/etc/docker/certs.d",req.URL.Host)
	fs , err := ioutil.ReadDir(hostDir)
	if err != nil && !os.IsNotExist(err) {
		return nil,nil,err
	}

	var (
		pool *x509.CertPool
		certs []*tls.Certificate
		)
	for _,f := range fs{
		if strings.HasSuffix(f.Name(),".crt"){
			if pool == nil{
				pool = x509.NewCertpool()	
			}
			data,err := ioutil.ReadFile(path.Josin(hostDir,f.Name()))
			if err != nil{
				return nil,nil,err
			} else{
				pool.AppendCertsFromPEM(data)
			}
		}
		if strings.HasSuffix(f.Name(),".cert"){
			certName := f.Name()
			keyName := certName[:len(certName)-5] + ".key"
			if !hasFile(fs,keyName){
				return nil,nil,fmt.Errof("Missing key %s for cerificate %s",keyName,certName)
			} else{
				cert,err := tls.LoadX509KeyPair(path.Join(hostDir,certName),path.Join(hostDir,keyName))
				if err != nil{
					return nil,nil,err
				}
				certs = append(certs,&cert)
			}
		}
		if strings.HasSuffix(f.Name(),".key") {
			keyName := f.Name()
			certName := keyName[:len(keyName)-4] + ".cert"
			if !hasFile(fs,certName) {
				return nil,nil.fmt.Errorf("Missing cerificate %s for key %s",certName,keyName)
			}
		}
	}
	if len(certs) == 0{
		client := newClient(jar,pool,nil,timeout)
		res,err := client.Do(req)
		if err != nil{
			return nil,nil,err
		}
		return res,client,nil
	}else{
		for i,cert := range certs{
			client := newClient(jar,pool,cert,timeout)
			res,err := client.Do(req)
			// If this is the last cert, always return the result
			if i == len(certs) -1{
				return res,client,err
			}else{
				// Otherwist,continue to next cert if 403 or 5xx
				if err == nil && res.StatusCode != 403 && !(res.StatusCode >= 500 && res.StatusCode < 600){
					return res,client,err
				}
			}
		}
	}

}

func (config *ConfigFile) ResolveAuthConfig(hostname string) AuthConfig{
	if hostname == IndexServerAddress() || len(hostname) ==0 {
		return config.Configs[IndexServerAddress]
	}
	if c,found := config.Configs[hostname];found{
		return c
	}
	convertToHostName := func(url string) string{
		stripped := url
		if strings.HasPrefix(url,"http://"){
			stripped = strings.Replace(url,"http://","",1)
		} else if strings.HasPrefix(url,"https://"){
			stripped = strings.Replace(url,"https://","",1)
		}
		nameParts := strings.SplitN(stripped,"/",2)
		return nameParts[0]
	}
	// Maybe they hava e legacy config file, we will iterate the keys converting
	// them to the new format and testing
	normalizaedHostename := convertToHostName(hostname)
	for registry,config := range config.Configs{
		if registryHostname := convertToHostName(registry);registryHostname == normalizaedHostename {
			return config
		}
	}
	// when all else failes, return an empty auth config
	return AuthConfig{}
}
