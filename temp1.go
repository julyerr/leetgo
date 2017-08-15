package main 

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main(){
	service := "7777"
	tcpAddr,err := net.ResolveTCPAddr("ipv4",service)
	checkError(err)
	listener,err:=net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	for{
		conn,err := listener.Accept()
		if err != nil{
			continue
		}
		go handlerClient(conn)
	}
}

func handlerClient(conn net.Conn){
	defer conn.Close()
	daytime:=time.LocalTime().String()
	conn.Write([]byte(daytime))
}

func checkError(err os.Error){
	if err != nil{
		fmt.Fprintf(os.Stderr,"fatal error:%s",err)
		os.Exit(1)
	}
}