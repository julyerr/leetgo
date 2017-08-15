package main 

import (
	"bufio"
	"fmt"
	"strings"
	"bytes"
	"io"
)

type Reader struct{
	buf []byte
	rd io.Reader
	r,w int
	err error
	lastByte int
	lastRuneSize int
}

type Writer struct{
	err error
	buf []byte
	n int 
	wr io.Writer
}

func main(){
	inputReadBuf := strings.NewReader("1234567890")
	reader := bufio.NewReader(inputReadBuf)

	buf:= bytes.NewBuffer(make([]byte,0))
	writer := bufio.NewWriter(buf)

	//Peek is just the slice citing
	b,err := reader.Peek(5)
	if err != nil{
		fmt.Printf("Read data error")
		return
	}

	b[0] = 'A'
	b,_ =  reader.Peek(5)
	writer.Write(b)
	writer.Flush()
	fmt.Println("buf(changed):",buf,"\ninputReadBuf(Not changed):",inputReadBuf)

	for {
		b1:=make([]byte,3)
		n1,_:=reader.Read(b1)
		if n1 <= 0{
			break
		}
		fmt.Println(n1,string(b1))
	}

	inputReadBuf2 := strings.NewReader("1234567")
	reader2 := bufio.NewReader(inputReadBuf2)
	b2,_:=reader2.ReadByte()
	fmt.Println(string(b2))
	reader2.UnreadByte()
	b2,_=reader2.ReadByte()
	fmt.Println(string(b2))

	inputReadBuf3 := strings.NewReader("中文1341513463")
	reader3 := bufio.NewReader(inputReadBuf3)
	b3,size,_:=reader3.ReadRune()
	fmt.Println(string(b3),size)
	reader3.UnReadRune()
	b3,size,_=reader3.ReadRune()
	fmt.Println(string(b3),size)
	b33,_=reader3.ReadByte()
	fmt.Println(string(b33))
	err3:=reader3.UnReadRune()
	if err3 != nil{
		fmt.Println(Err)
	}

	inputReadBuf4:=strings.NewReader("中文13251352")
	reader4 := bufio.NewReader(inputReadBuf4)
	// if not starting reading , buffered 0
	fmt.Println(reader4.Buffered())
	reader4.ReadByte()
	fmt.Println(reader4.Buffered())
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())
	reader4.ReadRune()
	fmt.Println(reader4.Buffered())

	inputReadBuf5 := strings.NewReader("中文123 456 789")
	reader5 := bufio.NewReader(inputReadBuf5)
	for {
		// just citing rather than copying
		b5,err := reader5.ReadSlice(' ')
		fmt.Println(string(b5))
		if err == io.EOF{
			break
		}
	}

	inputReadBuf6 := strings.NewReader("中文123\n456\n789")
	reader6 := bufio.NewReader(inputReadBuf6)
	for{
		l,p,err := reader6.ReadLine()
		fmt.Println(string(l),p,err)
		if err == io.EOF{
			break
		}
	}

	inputReadBuf7 := strings.NewReader("中文123;456;789")
	reader7 := bufio.NewReader(inputReadBuf7)
	for {
		line,err := reader7.ReadByte(';')
		fmt.Println(string(line))
		if err != nil {
			break
		}
	}

	inputReadBuf8 := strings.NewReader("中文123;456;789")
	reader8 := bufio.NewReader(inputReadBuf8)
	for {
		line,err := reader8.ReadString(';')
		fmt.Println(line)
		if err != nil{
			break
		}
	}

	b10 := bytes.NewBuffer(make([]byte,30))
	writer10.WriteString("35131436346")
	// if not flush then output will be 0
	fmt.Println(writer10.Available(),writer10.Buffered(),b10)
	writer10.Flush()
	fmt.Println(writer10.Available(),writer10.Buffered(),b10)

	b11 := bytes.NewBuffer(make([]byte,1024))
	writer11 := bufio.NewWriter(b11)
	writer11.WriteString("ABC")
	writer11.WriteByte(byte('M'))
	writer11.WriteRune(rune('号'))
	writer11.WriteRune(rune('么'))
	writer11.Write([]byte("1315364136"))
	writer11.Flush()
	fmt.Println(b11)

}	

