package ntdm

type Base struct {
	UserId     string `json:"-"` // 用户id
	Sort       bool   `json:"-"` // 排序
	ShowCNName bool   `json:"-"` // 显示中文名称
}
