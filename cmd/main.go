package main

import (
	"cmb_bank_direct/common"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func genResultCode() string {
	a := []string{"SUC0000", ""}
	return randShuffle(a)

}

func randShuffle(slice []string) string {
	rand.Seed(time.Now().UnixNano()) // 随机种子
	// 洗牌
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice[0]
}

// 构建测试服务
func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {

		var obj struct {
			Data string `json:"DATA" form:"DATA"`
			Uid  string `json:"UID" form:"UID"`
		}

		err := c.ShouldBind(&obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		var result common.Common
		json.Unmarshal([]byte(obj.Data), &result)

		c.JSON(200, gin.H{
			"response": gin.H{
				"body": gin.H{},
				"head": gin.H{
					"funcode":    result.FunCode,
					"userid":     obj.Uid,
					"resultcode": genResultCode(),
				},
			},
		})
	})
	r.Run(":12345")
}
