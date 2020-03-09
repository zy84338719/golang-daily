package scrapy

import (
	"fmt"
	"spider/scrapy/rules"
	"spider/utils"
)
var uidLi = utils.GetTargetUidList()

func ScrapyInfomation() {
	getInfoC := rules.GetDefaultCollector()
	getMoreInfoC := rules.GetDefaultCollector()
	rules.SetMoreInfoCallback(getMoreInfoC)

	rules.SetInfoCallback(getInfoC, getMoreInfoC)

	for _, uid := range uidLi {
		url := fmt.Sprintf("%s/%s/info", rules.BaseUrl, uid)
		getInfoC.Visit(url)
	}
	getInfoC.Wait()
	getMoreInfoC.Wait()
}

func ScrapyTweet() {
	getTweetsC := rules.GetDefaultCollector()
	getContentSubC := rules.GetDefaultCollector()
	rules.SetFullContentCallback(getContentSubC)
	getCommentSubC := rules.GetDefaultCollector()
	rules.SetCommentCallback(getCommentSubC)

	rules.SetTweetCallback(getTweetsC, getContentSubC, getCommentSubC)

	for _, uid := range uidLi {
		url := rules.GetTweetUrl(uid)
		getTweetsC.Visit(url)
	}
	getTweetsC.Wait()
	getContentSubC.Wait()
	getCommentSubC.Wait()
}

func ScrapyFollow() {
	getFollowC := rules.GetDefaultCollector()
	rules.SetFollowCallback(getFollowC)
	//read files
	for _, uid := range uidLi {
		url := rules.GetFollowUrl(uid)
		getFollowC.Visit(url)
	}
	getFollowC.Wait()
}

func ScrapyFans() {
	getFansC := rules.GetDefaultCollector()
	rules.SetFansCallback(getFansC)

	for _, uid := range uidLi {
		url := rules.GetFansUrl(uid)
		getFansC.Visit(url)
	}
	getFansC.Wait()
}