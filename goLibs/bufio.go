import "bufio"

type Reader struct{
	buf []byte
	rd io.Reader
	r,w int //r,w position
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

type ReadWriter struct{
	*Reader 
	*Writer
}

func NewReader(rd io.Reader) *Reader
func (b *Reader) Peek(n int) ([]byte,error) //return citing slice , not moving pos
func (b *Reader) Read(p []byte) (n int ,err error) 
func (b *Reader) ReadByte() (c byte,err error)
func (b *Reader) UnreadByte() error
func (b *Reader) ReadRune() (r rune,size int,err error)
func (b *Reader) UnreadRune() error
func (b *Reader) ReadSlice(delim byte) (line []byte,err error)
func (b *Reader) ReadLine(delim byte) (line []byte,err error)
func (b *Reader) ReadBytes(delim byte) (line string,err error)
func (b *Reader) ReadString(delim byte) (line string,err error)
func (b *Reader) Buffered() int //the space haven't read

func NewWriter(w io.Writer) *Writer
func (b *Writer) WriteString(s string) (int error)
func (b *Writer) WriteByte(c byte) error
func (b *Writer) WriteRune(r rune) error
func (b *Writer) Write(p []byte) (nn int,err error)
func (b *Writer) Flush() error

func (b *Reader) WriteTo(w io.Writer) (n int64,err error) 
func (b *Writer) ReadFrom(r io.Reader) (n int64,error)
