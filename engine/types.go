package engine

// 请求结构体
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

// 返回结果类型
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// 获取空返回结果
func NilParseResult([]byte) ParseResult {
	return ParseResult{}
}

// 返回结果结构
type Item struct {
	Title   string   `json:"title"`
	Price   string   `json:"price"`
	Details []string `json:"details"`
}
