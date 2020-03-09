package main

import (
	"fmt"
	"spider/config"
	"spider/login"
	"spider/scrapy"
)

func main() {
	var flag string
	fmt.Printf("是否登录？yes or no")
	fmt.Scanf("%s",&flag)
	if flag=="yes" {
		err:=login.Login()
		if err!=nil{
			fmt.Printf("登录失败",err)
			panic(err)
		}
	}
	if config.Conf.GetBool("SCRAPY_TYPE.Info") {
		scrapy.ScrapyInfomation()
	}
	if config.Conf.GetBool("SCRAPY_TYPE.Follow") {
		scrapy.ScrapyFollow()
	}
	//修复去重问题
	if config.Conf.GetBool("SCRAPY_TYPE.Fans") {
		scrapy.ScrapyFans()
	}
	if config.Conf.GetBool("SCRAPY_TYPE.Tweet.Main") {
		scrapy.ScrapyTweet()
	}
}