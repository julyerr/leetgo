//Client是一个http客户端,默认客户端(DefaultClient)将使用默认的发送机制的客户端.
// Client的Transport字段一般会含有内部状态(缓存TCP连接),因此Client类型值应尽量被重用而不是创建新的。
// 多个协程并发使用Clients是安全的.

type Client struct{
	// Transport指定执行独立、单次HTTP请求的机制如果Transport为nil,则使用DefaultTransport。
	Transport RoundTripper

	// CheckRedirect指定处理重定向的策略,如果CheckRedirect非nil,client将会在调用重定向之前调用它。
    // 参数req和via是将要执行的请求和已经执行的请求（时间越久的请求优先执行),如果CheckRedirect返回一个错误,
　　 //client的GetGet方法不会发送请求req,而是回之前得到的响应和该错误。
    // 如果CheckRedirect为nil，会采用默认策略：在连续10次请求后停止。
	CheckRedirect func(req *Request,via []*Request) error

	// Jar指定cookie管理器,如果Jar为nil,在请求中不会发送cookie,在回复中cookie也会被忽略。
	Jar CookieJar 

	// Timeout指定Client请求的时间限制,该超时限制包括连接时间、重定向和读取response body时间。
    // 计时器会在Head,Get,Post或Do方法返回后开始计时并在读到response.body后停止计时。
　　// Timeout为零值表示不设置超时。
    // Client的Transport字段必须支持CancelRequest方法,否则Client会在尝试用Head,Get,Post或Do方法执行请求时返回错误。
    // Client的Transport字段默认值（DefaultTransport）支持CancelRequest方法
	Timeout time.Duration 
}

// Do发送http请求并且返回一个http响应,遵守client的策略,如重定向,cookies以及auth等.
// 错误经常是由于策略引起的,当err是nil时,resp总会包含一个非nil的resp.body.
// 当调用者读完resp.body之后应该关闭它,如果resp.body没有关闭,则Client底层RoundTripper将无法重用存在的TCP连接去服务接下来的请求,如果resp.body非nil,则必须对其进行关闭
// 其中Do方法可以对Request进行一系列的设定，而其他的对request设定较少。
func (c *Client) Do(req *Request) (resp *Response,err error)

//利用get方法请求指定的url.Get请求指定的页面信息，并返回实体主体。
func (c *Client) Get(url string) (resp *Response,err error)

//利用head方法请求指定的url，Head只返回页面的首部。
func (c *Client) Head(url string) (resp *Response,err error)

//利用post方法请求指定的URl,如果body也是一个io.Closer,则在请求之后关闭它
func (c *Client) Post(url string,bodyType string,body io.Reader) (resp *Response,err error)

//利用post方法请求指定的url,利用data的key和value作为请求体.
func (c *Client) PostForm(url string,data url.Values) (resp *Response,err error)


