package main

import (
	"fmt"
	. "github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"math/rand"
	"strings"
	"sync"
	"time"
)

var userList []string

var mu sync.Mutex

func newApp() *Application {
	app := New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()

	userList = make([]string, 0)
	mu = sync.Mutex{}

	// http://localhost:8080
	_ = app.Run(Addr(":9999"))
}

type lotteryController struct {
	Ctx Context
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d\n", count)
}

// POST /import
func (c *lotteryController) PostImport() string {
	mu.Lock()
	defer mu.Unlock()
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	count := len(userList)
	//for _,u:=range users {
	//	u= strings.TrimSpace(u)
	//	if len(u)>0{
	//		userList=append(userList,u)
	//	}
	//
	//}
	userList = append(userList, users...)
	count2 := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d，成功导入用户数: %d\n", count2, (count2 - count))
}

// GET /lucky
func (c *lotteryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()                                // rand内部运算的随机数
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count)) // rand计算得到的随机数
		user := userList[index]                                      // 抽取到一个用户
		userList = append(userList[0:index], userList[index+1:]...)  // 移除这个用户
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		userList = userList[0:0]
		return fmt.Sprintf("当前中奖用户: %s, 剩余用户数: %d\n", user, count-1)
	} else {
		return fmt.Sprintf("已经没有参与用户，请先通过 /import 导入用户 \n")
	}

}
