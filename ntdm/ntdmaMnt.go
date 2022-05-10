package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaMnt 修改记账子单元
type NtdmaMnt struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Bbknbr string `json:"bbknbr"`           // 分行号
	Dmanbr string `json:"dmanbr"`           // 记账子单元编号， 必填 不超过10位
	Dmanam string `json:"dmanam"`           // 记账子单元名称， 必填
	Ovrctl string `json:"ovrctl,omitempty"` // 额度控制标志， 空：默认 Y   | Y： 允许透支 N：不允许透支
	Bcktyp string `json:"bcktyp,omitempty"` // 退票处理方式， 空：默认 N  | Y:退回原记账子单元 N：退回结算户
	Clstyp string `json:"clstyp,omitempty"` // 余额非零时是否可关闭， Y：可关闭， N：不可关闭 空：默认 Y

}

type NtdmaMntBody struct {
	Ntdmamntx1 []any `json:"ntdmamntx1,omitempty"`
}

func (p *NtdmaMnt) GetFuncode() string {
	return "NTDMAMNT"
}

func (p *NtdmaMnt) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaMnt) GetFunCnName() string {
	return "修改记账子单元"
}

func (p *NtdmaMnt) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaMnt) GetUserId() string {
	return p.UserId
}

func (p *NtdmaMnt) GetSort() bool {
	return p.Sort
}

func (p *NtdmaMnt) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaMntBody{
				Ntdmamntx1: list,
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
func (p *NtdmaMnt) Do(URL string) (result common.CommonResponseResult[any], err error) {
	body, result, err := common.GetResponseResult[any](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
