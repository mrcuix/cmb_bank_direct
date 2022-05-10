package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaThls 查询记账子单元历史交易
type NtdmaThls struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Dmanbr string `json:"dmanbr,omitempty"` // 填空表示查询该账号下所有记账子单元信息（包括默认记账子单元）
	Begdat string `json:"begdat,omitempty"` // 起始日期 ，必填， 只能查询近 13 个月的历史交易
	Enddat string `json:"enddat,omitempty"` // 小于当天，必填，大于或等于起始日期，与起始日期相距100天内
	Ctnkey string `json:"ctnkey,omitempty"` // 续传字段 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放到下次请求报文中继续查询

}

type NtdmaThlsDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Addusr string `json:"addusr"` // 添加用户
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称
	// Dmaccy string `json:"dmaccy"` // 记账子单元币种  10：人民币

	Trxnbr string `json:"trxnbr"` // 记账流水号
	Ccynbr string `json:"ccynbr"` // 币种

	Trxamt string `json:"trxamt"` // 交易金额
	Trxdir string `json:"trxdir"` // 交易方向 D-借方（支出） C-贷方（收入）
	Trxdat string `json:"trxdat"` // 交易金额
	Trxtim string `json:"trxtim"` // 交易时间
	Autflg string `json:"autflg"` // 记账方式 交易关联记账子单元的方式： 1：自动 2：手工 3：内部转账
	Rpyacc string `json:"rpyacc"` // 收方/付方账号
	Rpynam string `json:"rpynam"` // 收方/付方名称
	Trxtxt string `json:"trxtxt"` // 交易摘要

	Narinn string `json:"narinn"` // 原内部编号

}

type NtdmaThlsDetailResponse2 struct {
	Accnbr string `json:"accnbr"` // 主账号
	Dmanbr string `json:"dmanbr"` // 记账子单元编号 填空表示查询该账号下所有记账子单元信息（包括默认记账子单元）
	Dmanam string `json:"dmanam"` // 记账子单元名称

	Begdat string `json:"begdat"` // 起始日期 只能查询近 13 个月的历史交易
	Enddat string `json:"enddat"` // 结束日期 小于当天，大于或等于起始日期，与起始日期相距100天内
	Ctnkey string `json:"ctnkey"` // 续传字段 第一次传空，当返回接口有本接口并且该接口的续传字段有值，说明需要续传，取出续传字段放

}

type NtdmaThlsListResponse struct {
	Ntdmthlsz []NtdmaThlsDetailResponse  `json:"ntdmthlsz"`
	Ntdmthlsy []NtdmaThlsDetailResponse2 `json:"ntdmthlsy"`
}

type NtdmaThlsBody struct {
	Ntdmthlsy []any `json:"ntdmthlsy,omitempty"` // 续传接口
	Ntdmthlsz []any `json:"ntdmthlsz,omitempty"` // 详情信息
}

func (p *NtdmaThls) GetFuncode() string {
	return "NTDMTHLS"
}
func (p *NtdmaThls) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaThls) GetFunCnName() string {
	return "查询记账子单元历史交易"
}

func (p *NtdmaThls) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaThls) GetUserId() string {
	return p.UserId
}

func (p *NtdmaThls) GetSort() bool {
	return p.Sort
}

func (p *NtdmaThls) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaThlsBody{
				Ntdmthlsy: list,
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
func (p *NtdmaThls) Do(URL string) (result common.CommonResponseResult[NtdmaThlsListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaThlsListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
