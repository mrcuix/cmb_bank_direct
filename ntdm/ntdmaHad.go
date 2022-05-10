package ntdm

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
)

// NtdmaHad 查询单个记账子单元的历史余额
type NtdmaHad struct {
	Base
	Accnbr string `json:"accnbr"` // 账号， 必填
	Bbknbr string `json:"bbknbr"` // 分行号
	Begdat string `json:"begdat"` // 开始日期   `yyyymmdd` 格式  支持查询本年和去年的历史数据
	Enddat string `json:"enddat"` // 开始日期   `yyyymmdd` 格式  一次查询间隔不能超过3个月
	Dmanbr string `json:"dmanbr"` // 记账子单元编号

}

type NtdmaHadDetailResponse struct {
	Accnbr string `json:"accnbr"` // 主账号
	Dmanbr string `json:"dmanbr"` // 记账子单元编号
	Dmanam string `json:"dmanam"` // 记账子单元名称
	Ccynbr string `json:"ccynbr"` // 记账子单元币种  10：人民币

	Psedat string `json:"psedat"` // 日期
	Dmabal string `json:"dmabal"` // 余额

}

type NtdmaHadDetailResponse2 struct {
	Ctnflg string `json:"ctnflg"` // 续传标志 Y：有续传，N：无续传
	Dmanbr string `json:"dmanbr"` // 续传记账子单元  首次查询空，续传查询传上次查询返回的NtdmaHadZ2的续传记账子单元dmaNbr
}

type NtdmaHadListResponse struct {
	Ntdmahbdz1 []NtdmaHbdDetailResponse `json:"ntdmahbdz1"`
}

type NtdmaHadBody struct {
	Ntdmahadx1 []any `json:"ntdmahadx1,omitempty"` // 详细信息
}

func (p *NtdmaHad) GetFuncode() string {
	return "NTDMAHAD"
}

func (p *NtdmaHad) GetShowCNName() bool {
	return p.ShowCNName
}

func (p *NtdmaHad) GetFunCnName() string {
	return "查询单个记账子单元的历史余额"
}

func (p *NtdmaHad) SetUserId(val string) {
	p.UserId = val
}

func (p *NtdmaHad) GetUserId() string {
	return p.UserId
}

func (p *NtdmaHad) GetSort() bool {
	return p.Sort
}

func (p *NtdmaHad) ToJson() string {

	var (
		list []any
	)

	list = append(list, *p)

	reqData := common.Common{
		CommonRequest: common.CommonRequest{
			Body: NtdmaHadBody{
				Ntdmahadx1: list,
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
func (p *NtdmaHad) Do(URL string) (result common.CommonResponseResult[NtdmaHadListResponse], err error) {
	body, result, err := common.GetResponseResult[NtdmaHadListResponse](p, URL)
	fmt.Println("返回", body)
	fmt.Println("状态", result.CheckBlankResponseStatus())
	return
}
