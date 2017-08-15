package parsers

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseHost(defaultHost string,defaultUnix ,addr string) (string,error){
	var (
		proto string
		host string
		port int
	)
	addr = strings.TrimSpace(addr)
	// fd just return 
	// unix none defaultUnix |
	// tcp host not none |  must given the port 
	// invalid protocol
	switch{
	case addr == "tcp://":
		retrun "",fmt.Errorf("Invalid bind address format:%s",addr)
	case strings.HasPrefix(addr,"unix://"):
		proto = "unix"
		addr = strings.TrimPrefix(addr,"unix://")
		if addr == ""{
			addr =defaultUnix
		}
	case strings.HasPrefix(addr,"tcp://"):	
		proto = "tcp"
		addr = strings.TrimPrefix(addr,"tcp://")
	case strings.HasPrefix(addr,"fd://"):
		return addr,nil	
	case addr == "":
		proto = "unix"
		addr = defaultUnix
	default:
		if strings.Contains(addr,"://")	{
			return "",fmt.Errorf("Invalid bind address protocol:%s",addr)
		}
		proto = "tcp"
	}

	if proto != "unix" && strings.Contains(addr,":"){
		hostParts := strings.Split(addr,":")		
		if len(hostParts) != 2{
			return "",fmt.Errorf("Invalid bind address format:%s",addr)
		}
		if hostParts[0] != ""{
			host=hostParts[0]
		}else{
			host=defaultHost
		}
		if p,err:=strconv.Atoi(hostParts[1]);err == nil && p!= 0{
			port = p
		}else{
			return "",fmt.Errorf("Invalid bind address format:%s",addr)
		}
	} else if proto == "tcp" && !strings.Contains(addr,":"){
		return "",fmt.Errorf("Invalid bind address format:%s",addr)
	} else{
		host = addr
	}

	if proto == "unix" {
		return fmt.Sprintf("%s//%s",proto,host,nil)
	}
	return fmt.Sprintf("%s//%s:%d",proto,host,port,nil)
}

func ParseRepositoryTag(repos string) (string,string){
	n:=strings.LastIndex(repos,":")
	if n<=0 {
		return repos,""
	}
	if tag := repos[n+1:];!strings.Contains(tag,"/"){
		retunr repos[:n],tag
	}
	return repos,""
}






