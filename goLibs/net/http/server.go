 //该函数是一个http请求多路复用器，它将每一个请求的URL和一个注册模式的列表进行匹配，然后调用和URL最匹配的模式的处理器进行后续操作。
 // 模式是固定的、由根开始的路径，如"/favicon.ico"，或由根开始的子树，如"/images/" （注意结尾的斜杠）。
 ServeMux
 func NewServeMux() *ServeMux //初始化一个新的ServeMux
func (mux *ServeMux) Handle(pattern string, handler Handler) //将handler注册为指定的模式，如果该模式已经有了handler，则会出错panic。
　 func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))//将handler注册为指定的模式
func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) //根据指定的r.Method,r.Host以及r.RUL.Path返回一个用来处理给定请求的handler。该函数总是返回一个非nil的 handler，如果path不是一个规范格式，则handler会重定向到其规范path。Handler总是返回匹配该请求的的已注册模式；在内建重 定向处理器的情况下，pattern会在重定向后进行匹配。如果没有已注册模式可以应用于该请求，本方法将返回一个内建的"404 page not found"处理器和一个空字符串模式。
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)  //该函数用于将最接近请求url模式的handler分配给指定的请求。




 //该结构体定义一些用来运行HTTP Server的参数，如果Server默认为0的话，那么这也是一个有效的配置。
type Server struct {
    Addr           string        // 监听的TCP地址，如果为空字符串会使用":http"
    Handler        Handler       // 调用的处理器，如为nil会调用http.DefaultServeMux
    ReadTimeout    time.Duration // 请求的读取操作在超时前的最大持续时间
    WriteTimeout   time.Duration // 回复的写入操作在超时前的最大持续时间
    MaxHeaderBytes int           // 请求的头域最大长度，如为0则用DefaultMaxHeaderBytes
    TLSConfig      *tls.Config   // 可选的TLS配置，用于ListenAndServeTLS方法
    // TLSNextProto（可选地）指定一个函数来在一个NPN型协议升级出现时接管TLS连接的所有权。
    // 映射的键为商谈的协议名；映射的值为函数，该函数的Handler参数应处理HTTP请求，
    // 并且初始化Handler.ServeHTTP的*Request参数的TLS和RemoteAddr字段（如果未设置）。
    // 连接在函数返回时会自动关闭。
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    // ConnState字段指定一个可选的回调函数，该函数会在一个与客户端的连接改变状态时被调用。
    // 参见ConnState类型和相关常数获取细节。
    ConnState func(net.Conn, ConnState)
    // ErrorLog指定一个可选的日志记录器，用于记录接收连接时的错误和处理器不正常的行为。
    // 如果本字段为nil，日志会通过log包的标准日志记录器写入os.Stderr。
    ErrorLog *log.Logger
    // 内含隐藏或非导出字段
}

 func (srv *Server) ListenAndServe() error//监听TCP网络地址srv.Addr然后调用Serve来处理接下来连接的请求。如果srv.Addr是空的话，则使用“:http”。
func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error//ListenAndServeTLS监听srv.Addr确定的TCP地址，并且会调用Serve方法处理接收到的连接。必须提供证书文件和对应的私钥文 件。如果证书是由权威机构签发的，certFile参数必须是顺序串联的服务端证书和CA证书。如果srv.Addr为空字符串，会使 用":https"。
func (srv *Server) Serve(l net.Listener) error//接受Listener l的连接，创建一个新的服务协程。该服务协程读取请求然后调用srv.Handler来应答。实际上就是实现了对某个端口进行监听，然后创建相应的连接。
func (s *Server) SetKeepAlivesEnabled(v bool)//该函数控制是否http的keep-alives能够使用，默认情况下，keep-alives总是可用的。只有资源非常紧张的环境或者服务端在关闭进程中时，才应该关闭该功能。
