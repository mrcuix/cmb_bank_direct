package common

type Common struct {
	CommonRequest `json:"request"`
}

type CommonRequest struct {
	Body       any `json:"body"`
	CommonHead `json:"head" from:"head"`
}

type CommonHead struct {
	FunCode string `json:"funcode"`
	UserId  string `json:"userid"`
}
