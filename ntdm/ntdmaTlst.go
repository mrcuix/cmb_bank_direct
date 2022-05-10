package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaTlst 查询记账子单元当天交易
type NtdmaTlst struct {
	Base
	Accnbr string `json:"accnbr"`           // 账号， 必填
	Dmanbr string `json:"dmanbr,omitempty"` // 记账子单元编号， 必填 不超过10位
	Rsv50z string `json:"rsv50z,omitempty"` // 续传字段
}

// NtdmaTlstDetailResponse 查询记账子单元信息 返回内容结构体
type NtdmaTlstDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Addusr string `json:"addusr"` // 添加用户
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称

	Trxnbr string `json:"trxnbr"` // 记账流水号
	Ccynbr string `json:"ccynbr"` // 币种
	Dmaccy string `json:"dmaccy"` // 记账子单元币种  10：人民币
	Trxamt string `json:"trxamt"` // 交易金额
	Trxdir string `json:"trxdir"` // 交易方向   D-借方（支出）C-贷方（收入）
	Trxtim string `json:"trxtim"` // 交易时间
	Rpyacc string `json:"rpyacc"` // 收方/付方账号 与交易方向对应 交易方向为D时，这里为收方信息； 交易方向为C时，这里为付方信息。
	Rpynam string `json:"rpynam"` // 收方/付方名称
	Trxtxt string `json:"trxtxt"` // 交易摘要
	Narinn string `json:"narinn"` // 原内部编号 交易识别内部编号，可能为空。
	Mthflg string `json:"mthflg"` // 匹配标志

}

type NtdmaTlstListResponse struct {
	Ntdmtlstz []NtdmaTlstDetailResponse `json:"ntdmtlstz"`
}

type NtdmaTlstBody struct {
	Ntdmtlsty []any `json:"ntdmtlsty,omitempty"`
}

func (p *NtdmaTlst) GetFuncode() string {
	return "NTDMTLST"
}

func (p *NtdmaTlst) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaTlst) GetFunCnName() string {
	return "查询记账子单元当天交易"
}

func (p *NtdmaTlst) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaTlst) GetUserId() string {
	return p.UserId
}

func (p *NtdmaTlst) GetSort() bool {
	return p.Sort
}

func (p *NtdmaTlst) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaTlstBody{
				Ntdmtlsty: list,
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
func (p *NtdmaTlst) Do(URL string) (result common.CommonResponseResult[NtdmaTlstListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaTlstListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	fmt.Println("body Ntdmtlstz 字段空", len(result.Body.Ntdmtlstz) > 0)
	return
}
