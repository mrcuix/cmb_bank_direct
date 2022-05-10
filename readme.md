# 招商银行云直连 （交易管家记账）前置机接入 示例
交易管家-记账子单元

go version 1.18.1


#### 配置信息
```go
    var (
        URL = "http://localhost:12345/" // 前置机接入地址
        accnbr = "123456789012345" //15位银行账户
        userid = "N002467314" // 操作用户UID
        bankNo = "01" // 分行号
        dmanbr = "9910000001" // 记账子单元编号
    )
```


#### 新增记账子单元
```go

    // 新增记账子单元
	var book_add ntdm.NtdmaAddx
	book_add.Accnbr = accnbr
	book_add.Dmanbr = dmanbr
	book_add.Dmanam = "子单元003"
	book_add.UserId = userid
	book_add.Sort = true
	book_add.ShowCNName = true
	result, _ := book_add.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)

```


#### 修改记账子单元
```go

    // 修改记账子单元
	var book_edit ntdm.NtdmaMnt
	book_edit.Accnbr = accnbr
	book_edit.Bbknbr = bankNo
	book_edit.Dmanbr = dmanbr
	book_edit.Dmanam = "子单元003"
	book_edit.UserId = userid
	book_edit.Ovrctl = "Y"
	book_edit.Sort = true
	book_edit.ShowCNName = true
	result, _ := book_edit.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)

```

#### 查询记账子单元信息
```go
    // 查询记账子单元信息
	var book_info ntdm.NtdmaLst
	book_info.Accnbr = accnbr
	book_info.UserId = userid
	book_info.Sort = true
	book_info.ShowCNName = true
	result, _ := book_info.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)
```


#### 查询记账子单元当天交易
```go
    // 查询记账子单元当天交易
	var book_today_ls ntdm.NtdmaTlst
	book_today_ls.Accnbr = accnbr
	book_today_ls.UserId = userid
	// book_today_ls.Dmanbr = dmanbr // 填写即仅查询该记账子单元
	book_today_ls.Sort = true
	book_today_ls.ShowCNName = true
	result, _ := book_today_ls.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)
```


#### 查询记账子单元历史交易
```go
    // 查询记账子单元历史交易
	var book_history_ls ntdm.NtdmaThls
	book_history_ls.Accnbr = accnbr
	book_history_ls.Dmanbr = dmanbr 
	book_history_ls.Begdat = "20220410" // 开始日期
	book_history_ls.Enddat = "20220508" // 结束日期 不能大于等于当前日期
	book_history_ls.UserId = userid
	book_history_ls.Sort = true
	book_history_ls.ShowCNName = true
	result, _ := book_history_ls.Do(URL)
	fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)
```


#### 查询记账子单元余额
```go
    // 查询记账子单元余额
    var book_bal ntdm.NtdmaBal
    book_bal.Accnbr = accnbr
    book_bal.Dmanbr = dmanbr 
    book_bal.UserId = userid
    book_bal.Sort = true
    book_bal.ShowCNName = true
    result, _ := book_bal.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
    fmt.Println(result)
```

#### 查询所有记账子单元的某日余额
```go
    // 查询所有记账子单元的某日余额
    var book_all_someday_bal ntdm.NtdmaHbd
    book_all_someday_bal.Accnbr = accnbr
    book_all_someday_bal.Bbknbr = bankNo
    book_all_someday_bal.Qrydat = "20220310" // 查询日期 `yyyymmdd` 格式

    book_all_someday_bal.Dmanbr = dmanbr
    book_all_someday_bal.UserId = userid
    book_all_someday_bal.Sort = true
    book_all_someday_bal.ShowCNName = true
    result, _ := book_all_someday_bal.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
    fmt.Println(result)
```

#### 查询单个记账子单元的历史余额
```go
    // 查询单个记账子单元的历史余额
	var book_all_history_bal ntdm.NtdmaHad
	book_all_history_bal.Accnbr = accnbr
	book_all_history_bal.Bbknbr = bankNo
	book_all_history_bal.Dmanbr = dmanbr
	book_all_history_bal.Begdat = "20220310" // 开始日期
	book_all_history_bal.Enddat = "20220506" // 结束日期

	book_all_history_bal.Dmanbr = dmanbr
	book_all_history_bal.UserId = userid
	book_all_history_bal.Sort = true
	book_all_history_bal.ShowCNName = true
	result, _ := book_all_history_bal.Do(URL)
    fmt.Println("请求状态", result.CheckBlankResponseStatus())
	fmt.Println("返回信息", result.Resultmsg)
	fmt.Println(result)
```