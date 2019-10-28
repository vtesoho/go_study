package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var userList []String

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	userList = []string{}

	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintln("当前总共参与抽奖的用户数：%d\n", count)
}

func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	count1 := len(users)
	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	count2 := len(userList)
	return fmt.Println("当前总共参与抽奖的用户数： %d，成功导入的用户数：%d\n", count2, (count2 - count1))
}

// GET http://localhost:8080/import
func (c *lotteryController) GetLucky() string {
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintln("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		return fmt.Sprintln("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintln("当前已经没有参与用户，请先通过import 导入用户")
	}
}
