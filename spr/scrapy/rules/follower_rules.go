package rules

import (
	"fmt"
	"github.com/gocolly/colly"
	"gopkg.in/mgo.v2"
	"spider/mdb"
	"spider/utils"
	"strconv"
	"strings"
	"time"
)

func SetFollowCallback(getFollowC *colly.Collector){
	getFollowC.OnResponse(func(r *colly.Response) {
		content := string(r.Body)
		uid := utils.ReParse(`(\d+)`, r.Request.URL.String())
		if strings.Contains(r.Request.URL.String(), "page=1") {
			allPage := utils.ReParse(`/>&nbsp;1/(\d+)页</div>`, content)
			pageNum, _ := strconv.Atoi(allPage)
			for i := 2; i < (pageNum + 1); i++ {
				link := fmt.Sprintf("%s/%s/follow?page=%d", BaseUrl,uid,i)
				getFollowC.Visit(link)
			}
		}
	})
	getFollowC.OnXML(`//a[text()="关注他" or text()="关注她" or text()="取消关注"]/@href`, func(element *colly.XMLElement) {
		followUrl := element.Text
		uid := utils.ReParse(`uid=(\d+)`, followUrl)
		ID := utils.ReParse(`uid=(\d+)`, element.Request.URL.String())
		relationship := mdb.Relationships{}
		relationship.CrawlTime = int32(time.Now().Unix())
		relationship.FanId = ID
		relationship.FollowedId = uid
		relationship.Id_ = ID + "-" + uid
		err := mdb.Insert(dbName, "Relationships", relationship)
		if mgo.IsDup(err) {
			//有重复数据
			fmt.Println("already scrapy")
		}
	})
}

func GetFollowUrl(uid string) string{
	return fmt.Sprintf("%s/%s/follow?page=1", BaseUrl,uid)
}