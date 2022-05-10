package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaLst 查询记账子单元信息
type NtdmaLst struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Dmanbr string `json:"dmanbr,omitempty"` // 记账子单元编号， 必填 不超过10位
	Rsv50z string `json:"rsv50z,omitempty"` // 续传字段

}

// NtdmaLstDetailResponse 查询记账子单元信息 返回内容结构体
type NtdmaLstDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Addusr string `json:"addusr"` // 添加用户
	Bcktyp string `json:"bcktyp"` // 退票处理方式  Y:退回原记账子单元账户 N：退回结算户
	Clstyp string `json:"clstyp"` // 余额非零时是否可关闭  Y：可关闭
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称
	Dmaccy string `json:"dmaccy"` // 记账子单元币种  10：人民币
	Eftdat string `json:"eftdat"` // 生效日期
	Ovrctl string `json:"ovrctl"` // 额度控制标志  Y：允许透支N：不允许透支
	Rltchk string `json:"rltchk"` // 付方校验标志  Y：校验付款方账号 N：不校验
	Stscod string `json:"stscod"` // 状态

}

type NtdmaLstListResponse struct {
	Ntdmalstz []NtdmaLstDetailResponse `json:"ntdmalstz"`
}

type NtdmaLstBody struct {
	Ntdmalstx []any `json:"ntdmalstx,omitempty"`
}

func (p *NtdmaLst) GetFuncode() string {
	return "NTDMALST"
}

func (p *NtdmaLst) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaLst) GetFunCnName() string {
	return "查询记账子单元信息"
}

func (p *NtdmaLst) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaLst) GetUserId() string {
	return p.UserId
}
func (p *NtdmaLst) GetSort() bool {
	return p.Sort
}

func (p *NtdmaLst) ToJson() string {

	var (
		list []interface{}
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaLstBody{
				Ntdmalstx: list,
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

// 执行
func (p *NtdmaLst) Do(URL string) (result common.CommonResponseResult[NtdmaLstListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaLstListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
