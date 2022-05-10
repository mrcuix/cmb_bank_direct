package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaAddx 新增记账子单元
type NtdmaAddx struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Dmanbr string `json:"dmanbr"`           // 记账子单元编号， 必填 不超过10位
	Dmanam string `json:"dmanam"`           // 记账子单元名称， 必填
	Ovrctl string `json:"ovrctl,omitempty"` // 额度控制标志， 空：默认 Y   | Y： 允许透支 N：不允许透支
	Bcktyp string `json:"bcktyp,omitempty"` // 退票处理方式， 空：默认 N  | Y:退回原记账子单元 N：退回结算户
	Clstyp string `json:"clstyp,omitempty"` // 余额非零时是否可关闭， Y：可关闭， N：不可关闭 空：默认 Y

}

type NtdmaAddxBody struct {
	Ntdmaaddx []any `json:"ntdmaaddx,omitempty"`
}

func (p *NtdmaAddx) GetFuncode() string {
	return "NTDMAADD"
}

func (p *NtdmaAddx) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaAddx) GetFunCnName() string {
	return "新增记账子单元"
}

func (p *NtdmaAddx) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaAddx) GetUserId() string {
	return p.UserId
}

func (p *NtdmaAddx) GetSort() bool {
	return p.Sort
}

func (p *NtdmaAddx) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaAddxBody{
				Ntdmaaddx: list,
			},
			CommonHead: common.CommonHead{
				FunCode: p.GetFuncode(),
				UserId:  p.GetUserId(),
			},
		},
	}

	str, _ := json.Marshal(reqData)

	return string(str)
}

// Do 执行
func (p *NtdmaAddx) Do(URL string) (result common.CommonResponseResult[any], err error) {
	body, result, err := common.GetResponseResult[any](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
