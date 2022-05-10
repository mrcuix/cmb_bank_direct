package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaBal 查询记账子单元余额
type NtdmaBal struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Dmanbr string `json:"dmanbr,omitempty"` // 填空表示查询该账号下所有记账子单元信息（包括默认记账子单元）

	Rsv50z string `json:"rsv50z,omitempty"` // 续传字段 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询

}

type NtdmaBalDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Addusr string `json:"addusr"` // 添加用户
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称
	Dmaccy string `json:"dmaccy"` // 记账子单元币种  10：人民币

	Eftdat string `json:"eftdat"` // 生效日期
	Actbal string `json:"actbal"` // 余额

	Stscod string `json:"stscod"` // 状态 A：生效
	Mthaub string `json:"mthaub"` // 当月积数  余额*天数 累加

}

type NtdmaBalListResponse struct {
	Ntdmabalz []NtdmaTlstDetailResponse `json:"ntdmabalz"`
}

type NtdmaBalBody struct {
	Ntdmabalx []any `json:"ntdmabalx,omitempty"` // 续传接口
}

func (p *NtdmaBal) GetFuncode() string {
	return "NTDMABAL"
}

func (p *NtdmaBal) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaBal) GetFunCnName() string {
	return "查询记账子单元余额"
}

func (p *NtdmaBal) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaBal) GetUserId() string {
	return p.UserId
}

func (p *NtdmaBal) GetSort() bool {
	return p.Sort
}

func (p *NtdmaBal) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaBalBody{
				Ntdmabalx: list,
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
func (p *NtdmaBal) Do(URL string) (result common.CommonResponseResult[NtdmaBalListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaBalListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
