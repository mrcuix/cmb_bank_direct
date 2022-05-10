package main

import (
	"cmb_bank_direct/ntdm"
	"fmt"
)

var (
	URL = "http://localhost:12345/" // 前置机接入地址
)

func main() {

	accnbr := "123456789012345"
	userid := "N002467314"
	bankNo := "01" // 分行号

	dmanbr := "9910000001"

	fmt.Println("账号", accnbr)
	fmt.Println("用户id", userid)
	fmt.Println("分行号", bankNo)
	fmt.Println("子单元", dmanbr)
	fmt.Println("------------通用信息------------")

	// // 新增记账子单元
	// var book_add ntdm.NtdmaAddx
	// book_add.Accnbr = accnbr
	// book_add.Dmanbr = dmanbr
	// book_add.Dmanam = "子单元003"
	// book_add.UserId = userid
	// book_add.Sort = true
	// book_add.ShowCNName = true
	// result, _ := book_add.Do(URL)
	// fmt.Println(result)
	// // fmt.Println("body包含字段", result.BodyHasField())

	// // 修改记账子单元
	// var book_edit ntdm.NtdmaMnt
	// book_edit.Accnbr = accnbr
	// book_edit.Bbknbr = bankNo
	// book_edit.Dmanbr = dmanbr
	// book_edit.Dmanam = "子单元003"
	// book_edit.UserId = userid
	// book_edit.Ovrctl = "Y"
	// book_edit.Sort = true
	// book_edit.ShowCNName = true
	// result, _ := book_edit.Do(URL)
	// fmt.Println(result)
	// // fmt.Println("body包含字段", result.BodyHasField())

	// // 记账子单元 查询记账子单元信息
	// var book_info ntdm.NtdmaLst
	// book_info.Accnbr = accnbr
	// book_info.UserId = userid
	// book_info.Sort = true
	// book_info.ShowCNName = true
	// // book_info.Do(URL)
	// result, _ := book_info.Do(URL)
	// fmt.Println(result)

	// fmt.Println("body包含字段", result.BodyHasField())
	// t := result.Body.(response.NtdmalstzResponse)
	// t := result.Body.NtdmalstzResponse
	// fmt.Println(t)

	// // 查询记账子单元当天交易
	// var book_today_ls ntdm.NtdmaTlst
	// book_today_ls.Accnbr = accnbr
	// book_today_ls.UserId = userid
	// // book_today_ls.Dmanbr = dmanbr
	// book_today_ls.Sort = true
	// book_today_ls.ShowCNName = true
	// result, _ := book_today_ls.Do(URL)
	// fmt.Println(result)

	// fmt.Println("body包含字段", result.BodyHasField())

	// 查询记账子单元历史交易
	var book_history_ls ntdm.NtdmaThls
	book_history_ls.Accnbr = accnbr
	book_history_ls.Dmanbr = dmanbr
	book_history_ls.Begdat = "20220410"
	book_history_ls.Enddat = "20220508"
	book_history_ls.UserId = userid
	book_history_ls.Sort = true
	book_history_ls.ShowCNName = true
	result, _ := book_history_ls.Do(URL)
	fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)

	// fmt.Println("body包含字段", result.BodyHasField())

	// // 查询记账子单元余额
	// var book_bal ntdm.NtdmaBal
	// book_bal.Accnbr = accnbr
	// book_bal.Dmanbr = dmanbr
	// book_bal.UserId = userid
	// book_bal.Sort = true
	// book_bal.ShowCNName = true
	// result, _ := book_bal.Do(URL)
	// fmt.Println(result)
	// // fmt.Println("body包含字段", result.BodyHasField())

	// // 查询所有记账子单元的某日余额
	// var book_all_someday_bal ntdm.NtdmaHbd
	// book_all_someday_bal.Accnbr = accnbr
	// book_all_someday_bal.Bbknbr = bankNo
	// book_all_someday_bal.Qrydat = "20220310"

	// book_all_someday_bal.Dmanbr = dmanbr
	// book_all_someday_bal.UserId = userid
	// book_all_someday_bal.Sort = true
	// book_all_someday_bal.ShowCNName = true
	// result, _ := book_all_someday_bal.Do(URL)
	// fmt.Println(result)
	// // fmt.Println("body包含字段", result.BodyHasField())

	// // 查询单个记账子单元的历史余额
	// var book_all_history_bal ntdm.NtdmaHad
	// book_all_history_bal.Accnbr = accnbr
	// book_all_history_bal.Bbknbr = bankNo
	// book_all_history_bal.Dmanbr = dmanbr
	// book_all_history_bal.Begdat = "20220310"
	// book_all_history_bal.Enddat = "20220506"

	// book_all_history_bal.Dmanbr = dmanbr
	// book_all_history_bal.UserId = userid
	// book_all_history_bal.Sort = true
	// book_all_history_bal.ShowCNName = true
	// result, _ := book_all_history_bal.Do(URL)
	// fmt.Println(result)
	// // fmt.Println("body包含字段", result.BodyHasField())

}
