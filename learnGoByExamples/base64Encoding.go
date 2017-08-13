package main 

// encoding/base64
// StdEncoding.EncodeToString StdEncoding.DecodeString
// URLEncoding.EncodeToString URLEncoding.DecodeString
//almost all the encoding , decoding will use the byte as
	// a middle layer 
// string -> byte []byte(str1)
// byte -> string string(byte1)

// interface{} make
 // eg: names := make(interface{},0,10) names = append(names,10) names = append(names,"string")

import b64 "encoding/base64"
import "fmt"

func main(){
	data:="abc123!?$*&()-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)
	sDec,_ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec,_:=b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}
