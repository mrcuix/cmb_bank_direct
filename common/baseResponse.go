package common

type CommonResponseResult[T any] struct {
	CommonResponse[T] `json:"response"` // 响应
}

type CommonResponse[T any] struct {
	Body               T             `json:"body,omitempty"`
	CommonResponseHead `json:"head"` // 响应报文头
}

type CommonResponseHead struct {
	ResultCode string `json:"resultcode"` // 结果代码
	Reqid      string `json:"reqid"`      // 请求 id
	Resultmsg  string `json:"resultmsg"`  // 返回 信息提示
	FunCode    string `json:"funcode"`    // 函数码
	Rspid      string `json:"rspid"`      // 返回 id
	Userid     string `json:"userid"`     // 用户
	BizCode    string `json:"bizcode"`    //
}

// 检查银行请求是否成功
func (r *CommonResponse[T]) CheckBlankResponseStatus() (result bool) {
	result = r.ResultCode == "SUC0000"
	return
}
