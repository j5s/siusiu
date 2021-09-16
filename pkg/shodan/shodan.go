package shodan

//BaseURL shodan域名
const BaseURL string = "https://api.shodan.io"

type Client struct {
	apiKey  string
	baseURL string
}

//NewClient 创建client实例
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
