package models

//Resp 响应
type Resp struct {
	StateCode int
	Reason    string
	Result    Result
}

//Result 具体信息
type Result struct {
	CompanyName string
	CompanyType string
	SiteLicense string
	SiteName    string
	MainPage    string
	Owner       string
}
