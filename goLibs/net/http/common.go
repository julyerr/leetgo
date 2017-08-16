
//代表客户端给服务器端发送的一个请求.该字段在服务器端和客户端使用时区别很大.
type Request struct{
	// Method指定HTTP方法(GET,POST,PUT等)默认使用GET方法。
	Method string

	// URL在服务端表示被请求的URI(uniform resource identifier,统一资源标识符)，在客户端表示要访问的URL。
    // 在服务端,URL字段是解析请求行的URI（保存在RequestURI字段）得到的,对大多数请求来说,除了Path和RawQuery之外的字段都是空字符串。
    // 在客户端,URL的Host字段指定了要连接的服务器,而Request的Host字段指定要发送的HTTP请求的Host头的值。
    URL *url.URL

    // 接收到的请求的协议版本。client的Request总是使用HTTP/1.1
    Proto string
    ProtoMajor int 
    ProtoMinor int 

    // Header字段用来表示HTTP请求的头域。如果header（多行键值对格式）为：
    //    accept-encoding: gzip, deflate
    //    Accept-Language: en-us
    //    Connection: keep-alive

    // HTTP规定header的键名（头名）是区分大小写的，请求的解析器通过规范化头域的键名来实现这点,即首字母大写,其他字母小写,通过"-"进行分割。
    // 在客户端的请求，可能会被自动添加或重写Header中的特定的头，参见Request.Write方法。
    Header Header

    // Body是请求的主体.对于客户端请求来说,一个nil body意味着没有body,http Client的Transport字段负责调用Body的Close方法。
    // 在服务端，Body字段总是非nil的；但在没有主体时，读取Body会立刻返回EOF.Server会关闭请求主体，而ServeHTTP处理器不需要关闭Body字段。
    Body *io.ReadCloser

    // 在客户端,如果Body非nil而该字段为0,表示不知道Body的长度。
    ContentLength int64

    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
    // 本字段一般会被忽略。当发送或接受请求时，会自动添加或移除"chunked"传输编码。
    TransferEncoding []string

    // Close在服务端指定是否在回复请求后关闭连接，在客户端指定是否在发送请求后关闭连接。
    Close bool

    // 对于服务器端请求,Host指定URL指向的主机,可能的格式是host:port.对于客户请求,Host用来重写请求的Host头,如过该字段为""，Request.Write方法会使用URL.Host来进行赋值。
    Host string

     // Form是解析好的表单数据，包括URL字段的query参数和POST或PUT的表单数据.本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
    Form url.Values

      // MultipartForm是解析好的多部件表单，包括上传的文件.本字段只有在调用ParseMultipartForm后才有效。http客户端中会忽略MultipartForm并且使用Body替代
    MultipartForm *multipart.Form

     // 几乎没有HTTP客户端、服务端或代理支持HTTP trailer。
    Trailer Header

    // RemoteAddr允许HTTP服务器和其他软件记录该请求的来源地址,该字段经常用于日志.本字段不是ReadRequest函数填写的，也没有定义格式。
    // 本包的HTTP服务器会在调用处理器之前设置RemoteAddr为"IP:port"格式的地址.客户端会忽略请求中的RemoteAddr字段。
    RemoteAddr string

      // RequestURI是客户端发送到服务端的请求中未修改的URI(参见RFC 2616,Section 5.1),如果在http请求中设置该字段便会报错.
    RequestURI string

    // TLS字段允许HTTP服务器和其他软件记录接收到该请求的TLS连接的信息,本字段不是ReadRequest函数填写的。
    // 对启用了TLS的连接，本包的HTTP服务器会在调用处理器之前设置TLS字段，否则将设TLS为nil。
    // 客户端会忽略请求中的TLS字段。
    TLS *tls.ConnectionState
}

 //利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口
// ，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostForm方法以及Transport.RoundTrip方法关闭。
func NewRequest(method,urlStr string,body io.Reader) (*Request,error)

//从b中读取和解析一个请求.
func ReadRequest(b *bufio.Reader) (req *Request,err error)

// 给request添加cookie,AddCookie向请求中添加一个cookie.按照RFC 6265 section 5.4的规则,AddCookie不会添加超过一个Cookie头字段.
// 这表示所有的cookie都写在同一行,用分号分隔（cookie内部用逗号分隔属性）
func (r *Request) AddCookie(c *Cookie)


//返回request中指定名name的cookie，如果没有发现，返回ErrNoCookie
func (r *Request) Cookies(name string) *Cookie

//返回该请求的所有cookies
func (r *Request) Cookies() []*Cookie

//利用提供的用户名和密码给http基本权限提供具有一定权限的header。当使用http基本授权时，用户名和密码是不加密的
func (r *Request) SetBasicAuth(username, password string) 

//如果在request中发送，该函数返回客户端的user-Agent
func (r *Request) UserAgent() string

//对于指定格式的key，FormFile返回符合条件的第一个文件，如果有必要的话，该函数会调用ParseMultipartForm和ParseForm。
func (r *Request) FormFile(key string) (multipart.File,*multipart.FileHeader,error)

//返回key获取的队列中第一个值。在查询过程中post和put中的主题参数优先级高于url中的value。
// 为了访问相同key的多个值，调用ParseForm然后直接检查RequestForm。
  func (r *Request) FormValue(key string) string

//如果这是一个有多部分组成的post请求，该函数将会返回一个MIME 多部分reader，否则的话将会返回一个nil和error。
// 使用本函数代替ParseMultipartForm可以将请求body当做流stream来处理。
func (r *Request) MultipartReader() (*multipart.Reader,error)

//解析URL中的查询字符串，并将解析结果更新到r.Form字段。
// 对于POST或PUT请求，ParseForm还会将body当作表单解析，并将结果既更新到r.PostForm也更新到r.Form
// 解析结果中，POST或PUT请求主体要优先于URL查询字符串（同名变量，主体的值在查询字符串的值前面）。
// 如果请求的主体的大小没有被MaxBytesReader函数设定限制，其大小默认限制为开头10MB。ParseMultipartForm会自动调用ParseForm。重复调用本方法是无意义的。
func (r *Request) ParseForm() error

//ParseMultipartForm将请求的主体作为multipart/form-data解析。
// 请求的整个主体都会被解析，得到的文件记录最多 maxMemery字节保存在内存，其余部分保存在硬盘的temp文件里。如果必要，ParseMultipartForm会自行调用 ParseForm。重复调用本方法是无意义的。
func (r *Request) ParseMultipartForm(maxMemory int64) error 

//返回post或者put请求body指定元素的第一个值，其中url中的参数被忽略。
func (r *Request) PostFormValue(key string) string

//检测在request中使用的http协议是否至少是major.minor
   func (r *Request) ProtoAtLeast(major, minor int) bool

//如果request中有refer，那么refer返回相应的url。
   // Referer在request中是拼错的，这个错误从http初期就已经存在了。该值也可以从Headermap中利用Header["Referer"]获取；在使用过程中利用Referer这个方法而不是map的形式的好处是在编译过程中可以检查方法的错误，而无法检查map中key的错误。   
 func (r *Request) Referer() string   

//Write方法以有线格式将HTTP/1.1请求写入w（用于将请求写入下层TCPConn等）。
 // 本方法会考虑请求的如下字段：Host URL Method (defaults to "GET") Header ContentLength TransferEncoding Body
// 　如果存在Body，ContentLength字段<= 0且TransferEncoding字段未显式设置为["identity"]，Write方法会显式添加"Transfer-Encoding: chunked"到请求的头域。Body字段会在发送完请求后关闭。 
 func (r *Request) Write(w io.Writer) error

//该函数与Write方法类似，但是该方法写的request是按照http代理的格式去写。
 // 尤其是，按照RFC 2616 Section 5.1.2，WriteProxy会使用绝对URI（包括协议和主机名）来初始化请求的第1行（Request-URI行）。无论何种情况，WriteProxy都会使用r.Host或r.URL.Host设置Host头
 func (r *Request) WriteProxy(w io.Writer) error





 //指对于一个http请求的响应response
type Response struct {
    Status     string // 例如"200 OK"
    StatusCode int    // 例如200
    Proto      string // 例如"HTTP/1.0"
    ProtoMajor int    // 主协议号：例如1
    ProtoMinor int    // 副协议号：例如0
    // Header保管header的key values，如果response中有多个header中具有相同的key，那么Header中保存为该键对应用逗号分隔串联起来的这些头的值// 被本结构体中的其他字段复制保管的头（如ContentLength）会从Header中删掉。Header中的键都是规范化的，参见CanonicalHeaderKey函数
    Header Header
    // Body代表response的主体。http的client和Transport确保这个body永远非nil，即使response没有body或body长度为0。调用者也需要关闭这个body
    // 如果服务端采用"chunked"传输编码发送的回复，Body字段会自动进行解码。
    Body io.ReadCloser
    // ContentLength记录相关内容的长度。
    // 其值为-1表示长度未知（采用chunked传输编码）
    // 除非对应的Request.Method是"HEAD"，其值>=0表示可以从Body读取的字节数
    ContentLength int64
    // TransferEncoding按从最外到最里的顺序列出传输编码，空切片表示"identity"编码。
    TransferEncoding []string
    // Close记录头域是否指定应在读取完主体后关闭连接。（即Connection头）
    // 该值是给客户端的建议，Response.Write方法的ReadResponse函数都不会关闭连接。
    Close bool
    // Trailer字段保存和头域相同格式的trailer键值对，和Header字段相同类型
    Trailer Header
    // Request是用来获取此回复的请求，Request的Body字段是nil（因为已经被用掉了）这个字段是被Client类型发出请求并获得回复后填充的
    Request *Request
    // TLS包含接收到该回复的TLS连接的信息。 对未加密的回复，本字段为nil。返回的指针是被（同一TLS连接接收到的）回复共享的，不应被修改。
    TLS *tls.ConnectionState
}

func (r *Response) Cookies() []*Cookie//解析cookie并返回在header中利用set-Cookie设定的cookie值。
func (r *Response) Location() (*url.URL, error)//返回response中Location的header值的url。如果该值存在的话，则对于请求问题可以解决相对重定向的问题，如果该值为nil，则返回ErrNOLocation的错误。
func (r *Response) ProtoAtLeast(major, minor int) bool//判定在response中使用的http协议是否至少是major.minor的形式。
func (r *Response) Write(w io.Writer) error//将response中信息按照线性格式写入w中。



//该接口被http handler用来构建一个http response
type ResponseWriter interface {
    // Header返回一个Header类型值，该值会被WriteHeader方法发送.在调用WriteHeader或Write方法后再改变header值是不起作用的。
    Header() Header
    // WriteHeader该方法发送HTTP回复的头域和状态码。如果没有被显式调用，第一次调用Write时会触发隐式调用WriteHeader(http.StatusOK)
    // 因此，显示调用WriterHeader主要用于发送错误状态码。
    WriteHeader(int)
    // Write向连接中写入数据，该数据作为HTTP response的一部分。如果被调用时还没有调用WriteHeader，本方法会先调用WriteHeader(http.StatusOK)
    // 如果Header中没有"Content-Type"键，本方法会使用包函数DetectContentType检查数据的前512字节，将返回值作为该键的值。
    Write([]byte) (int, error)
}







type RoundTripper//该函数是一个执行简单http事务的接口，该接口在被多协程并发使用时必须是安全的。
type RoundTripper interface {
    // RoundTrip执行单次HTTP事务，返回request的response，RoundTrip不应试图解析该回复。
    // 尤其要注意，只要RoundTrip获得了一个回复，不管该回复的HTTP状态码如何，它必须将返回值err设置为nil。
    // 非nil的返回值err应该留给获取回复失败的情况。类似的，RoundTrip不能试图管理高层协议，如重定向、认证或者cookie。
    // RoundTrip除了从请求的主体读取并关闭主体之外，不能够对请求做任何修改，包括（请求的）错误。
    // RoundTrip函数接收的请求的URL和Header字段必须保证是初始化了的。
    RoundTrip(*Request) (*Response, error)
}

func NewFileTransport(fs FileSystem) RoundTripper //该函数返回一个RoundTripper接口，服务指定的文件系统。 返回的RoundTripper接口会忽略接收的请求中的URL主机及其他绝大多数属性。该函数的典型应用是给Transport类型的值注册"file"协议






  //该结构体实现了RoundTripper接口，支持HTTP，HTTPS以及HTTP代理，TranSport也能缓存连接供将来使用。
type Transport struct {
    // Proxy指定一个对给定请求返回代理的函数。如果该函数返回了非nil的错误值，请求的执行就会中断并返回该错误。
    // 如果Proxy为nil或返回nil的*URL值，将不使用代理。
    Proxy func(*Request) (*url.URL, error)
    // Dial指定创建未加密TCP连接的dial函数。如果Dial为nil，会使用net.Dial。
    Dial func(network, addr string) (net.Conn, error)
　　// DialTls利用一个可选的dial函数来为非代理的https请求创建一个TLS连接。如果DialTLS为nil的话，那么使用Dial和TLSClientConfig。
　　//如果DialTLS被设定，那么Dial钩子不被用于HTTPS请求和TLSClientConfig并且TLSHandshakeTimeout被忽略。返回的net.conn默认已经经过了TLS握手协议。
　　DialTLS func(network, addr string) (net.Conn, error) 
    // TLSClientConfig指定用于tls.Client的TLS配置信息。如果该字段为nil，会使用默认的配置信息。
    TLSClientConfig *tls.Config
    // TLSHandshakeTimeout指定等待TLS握手完成的最长时间。零值表示不设置超时。
    TLSHandshakeTimeout time.Duration
    // 如果DisableKeepAlives为真，不同HTTP请求之间TCP连接的重用将被阻止。
    DisableKeepAlives bool
    // 如果DisableCompression为真，会禁止Transport在请求中没有Accept-Encoding头时，
    // 主动添加"Accept-Encoding: gzip"头，以获取压缩数据。
    // 如果Transport自己请求gzip并得到了压缩后的回复，它会主动解压缩回复的主体。
    // 但如果用户显式的请求gzip压缩数据，Transport是不会主动解压缩的。
    DisableCompression bool
    // 如果MaxIdleConnsPerHost不为0，会控制每个主机下的最大闲置连接数目。
    // 如果MaxIdleConnsPerHost为0，会使用DefaultMaxIdleConnsPerHost。
    MaxIdleConnsPerHost int
    // ResponseHeaderTimeout指定在发送完请求（包括其可能的主体）之后，
    // 等待接收服务端的回复的头域的最大时间。零值表示不设置超时。
    // 该时间不包括读取回复主体的时间。
ResponseHeaderTimeout time.Duration
}

func (t *Transport) CancelRequest(req *Request)//通过关闭连接来取消传送中的请求。
func (t *Transport) CloseIdleConnections()//关闭所有之前请求但目前处于空闲状态的连接。该方法并不中断任何正在使用的连接。
func (t *Transport) RegisterProtocol(scheme string, rt RoundTripper)//RegisterProtocol注册一个新的名为scheme的协议。t会将使用scheme协议的请求转交给rt。rt有责任模拟HTTP请求的语义。RegisterProtocol可以被其他包用于提供"ftp"或"file"等协议的实现。
func (t *Transport) RoundTrip(req *Request) (resp *Response, err error)//该函数实现了RoundTripper接口，对于高层http客户端支持，例如处理cookies以及重定向，查看Get，Post以及Client类型。






type CloseNotifier//该接口被ResponseWriter用来实时检测底层连接是否已经断开.如果客户端已经断开连接,该机制可以在服务端响应之前取消二者之间的长连接.
type CloseNotifier interface {
        // 当客户端断开连接时,CloseNotifier接受一个通知
        CloseNotify() <-chan bool
}


 //表示客户端连接服务端的状态
// type ConnState int
const (
    // StateNew代表一个新的连接，将要立刻发送请求。
    // 连接从这个状态开始，然后转变为StateAlive或StateClosed。
    StateNew ConnState = iota
    // StateActive代表一个已经读取了请求数据1到多个字节的连接。
    // 用于StateAlive的Server.ConnState回调函数在将连接交付给处理器之前被触发，
    // 等到请求被处理完后，Server.ConnState回调函数再次被触发。
    // 在请求被处理后，连接状态改变为StateClosed、StateHijacked或StateIdle。
    StateActive
    // StateIdle代表一个已经处理完了请求、处在闲置状态、等待新请求的连接。
    // 连接状态可以从StateIdle改变为StateActive或StateClosed。
    StateIdle
    // 代表一个被劫持的连接。这是一个终止状态，不会转变为StateClosed。
    StateHijacked
    // StateClosed代表一个关闭的连接。
    // 这是一个终止状态。被劫持的连接不会转变为StateClosed。
    StateClosed
)





type Cookie//常用SetCooker用来给http的请求或者http的response设置cookie
type Cookie struct {

        Name       string  //名字
        Value      string  //值
        Path       string   //路径
        Domain     string   
        Expires    time.Time //过期时间
        RawExpires string

        // MaxAge=0 意味着 没有'Max-Age'属性指定.
        // MaxAge<0 意味着 立即删除cookie
        // MaxAge>0 意味着设定了MaxAge属性,并且其单位是秒
        MaxAge   int
        Secure   bool
        HttpOnly bool
        Raw      string
        Unparsed []string // 未解析的属性值对
}

func (c *Cookie) String() string//该函数返回cookie的序列化结果。如果只设置了Name和Value字段，序列化结果可用于HTTP请求的Cookie头或者HTTP回复的Set-Cookie头；如果设置了其他字段，序列化结果只能用于HTTP回复的Set-Cookie头。





//在http请求中，CookieJar管理存储和使用cookies.Cookiejar的实现必须被多协程并发使用时是安全的.
type CookieJar interface {
        // SetCookies 处理从url接收到的cookie,是否存储这个cookies取决于jar的策略和实现
        SetCookies(u *url.URL, cookies []*Cookie)

        // Cookies 返回发送到指定url的cookies
        Cookies(u *url.URL) []*Cookie
}







type Header map[string][]string
func (h Header) Add(key, value string)//将key,value组成键值对添加到header中
func (h Header) Del(key string)  //header中删除对应的key-value对
func (h Header) Get(key string) string //获取指定key的第一个value,如果key没有对应的value,则返回"",为了获取一个key的多个值,用CanonicalHeaderKey来获取标准规范格式.
func (h Header) Set(key, value string) //给一个key设定为相应的value.
func (h Header) Write(w io.Writer) error//利用线格式来写header
func (h Header) WriteSubset(w io.Writer, exclude map[string]bool) error//利用线模式写header,如果exclude不为nil,则在exclude中包含的exclude[key] == true不被写入.






type Dir//使用一个局限于指定目录树的本地文件系统实现一个文件系统.一个空目录被当做当前目录"."
type Dir string
    func (d Dir) Open(name string) (File, error)
type File //File是通过FileSystem的Open方法返回的,并且能够被FileServer实现.该方法与*os.File行为表现一样.
type File interface {
        io.Closer
        io.Reader
        Readdir(count int) ([]os.FileInfo, error)
        Seek(offset int64, whence int) (int64, error)
        Stat() (os.FileInfo, error)
}
type FileSystem//实现了对一系列指定文件的访问,其中文件路径之间通过分隔符进行分割.
type FileSystem interface {
        Open(name string) (File, error)
}
 

type Flusher //responsewriters允许http控制器将缓存数据刷新入client.然而如果client是通过http代理连接服务器,这个缓存数据也可能是在整个response结束后才能到达客户端.
type Flusher interface {
        // Flush将任何缓存数据发送到client
        Flush()
}



type Hijacker
type Hijacker interface {
        // Hijack让调用者接管连接,在调用Hijack()后,http server库将不再对该连接进行处理,对于该连接的管理和关闭责任将由调用者接管.
        Hijack() (net.Conn, *bufio.ReadWriter, error) //conn表示连接对象，bufrw代表该连接的读写缓存对象。
}





type Handler //实现Handler接口的对象可以注册到HTTP服务端，为指定的路径或者子树提供服务。ServeHTTP应该将回复的header和数据写入ResponseWriter接口然后返回。返回意味着该请求已经结束，HTTP服务端可以转移向该连接上的下一个请求。
//如果ServeHTTP崩溃panic,那么ServeHTTP的调用者假定这个panic的影响与活动请求是隔离的,二者互不影响.调用者恢复panic,将stack trace记录到错误日志中,然后挂起这个连接.
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
　func FileServer(root FileSystem) Handler // FileServer返回一个使用FileSystem接口提供文件访问服务的HTTP处理器。可以使用httpDir来使用操作系统的FileSystem接口实现。其主要用来实现静态文件的展示。
func NotFoundHandler() Handler //返回一个简单的请求处理器,该处理器对任何请求都会返回"404 page not found"
func RedirectHandler(url string, code int) Handler//使用给定的状态码将它接受到的任何请求都重定向到给定的url
func StripPrefix(prefix string, h Handler) Handler//将请求url.path中移出指定的前缀,然后将省下的请求交给handler h来处理,对于那些不是以指定前缀开始的路径请求,该函数返回一个http 404 not found 的错误.
func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler //具有超时限制的handler,该函数返回的新Handler调用h中的ServerHTTP来处理每次请求,但是如果一次调用超出时间限制,那么就会返回给请求者一个503服务请求不可达的消息,并且在ResponseWriter返回超时错误.

type HandlerFunc//HandlerFunc type是一个适配器，通过类型转换我们可以将普通的函数作为HTTP处理器使用。如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Handler接口。
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) //ServeHttp调用f(w,r)


