package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaHbd 查询所有记账子单元的某日余额
type NtdmaHbd struct {
	Base
	Accnbr string `json:"accnbr"` // 账号， 必填
	Bbknbr string `json:"bbknbr"` // 分行号
	Qrydat string `json:"qrydat"` // 查询日期   `yyyymmdd` 格式  支持查询本年和去年的历史数据
	Dmanbr string `json:"dmanbr"` // 续传记账子单元  首次查询空，续传查询传上次查询返回的NTDMAHBDZ2的续传记账子单元dmaNbr

}

type NtdmaHbdDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称
	Ccynbr string `json:"ccynbr"` // 记账子单元币种  10：人民币

	Psedat string `json:"psedat"` // 日期
	Dmabal string `json:"dmabal"` // 余额

}

type NtdmaHbdDetailResponse2 struct {
	Ctnflg string `json:"ctnflg"` // 续传标志 Y：有续传，N：无续传
	Dmanbr string `json:"dmanbr"` // 续传记账子单元  首次查询空，续传查询传上次查询返回的NTDMAHBDZ2的续传记账子单元dmaNbr
}

type NtdmaHbdListResponse struct {
	Ntdmahbdz1 []NtdmaHbdDetailResponse  `json:"ntdmahbdz1"`
	Ntdmahbdz2 []NtdmaHbdDetailResponse2 `json:"ntdmahbdz2"`
}

type NtdmaHbdBody struct {
	Ntdmahbdx1 []any `json:"ntdmahbdx1,omitempty"` // 续传接口
}

func (p *NtdmaHbd) GetFuncode() string {
	return "NtdmaHbd"
}

func (p *NtdmaHbd) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaHbd) GetFunCnName() string {
	return "查询所有记账子单元的某日余额"
}

func (p *NtdmaHbd) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaHbd) GetUserId() string {
	return p.UserId
}

func (p *NtdmaHbd) GetSort() bool {
	return p.Sort
}

func (p *NtdmaHbd) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaHbdBody{
				Ntdmahbdx1: list,
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
func (p *NtdmaHbd) Do(URL string) (result common.CommonResponseResult[NtdmaHbdListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaHbdListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
