package utils

import (
	"bufio"
	"github.com/prometheus/common/log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var ExecPath = getExecutePath()

func ReParse(pattern string, content string) string {
	str := regexp.MustCompile(pattern).FindAllStringSubmatch(content, -1)
	if str != nil {
		if len(str[0]) == 1 {
			return str[0][0]
		}
		return str[0][1]
	}
	return ""
}

func ReParseMayLi(pattern string, content string) [][]string {
	str := regexp.MustCompile(pattern).FindAllStringSubmatch(content, -1)
	return str
}

func ConvTime(timeStr string) string {
	now_time := time.Now()
	if strings.Contains(timeStr, "分钟前") {
		min, _ := strconv.Atoi(ReParse(`^(\d+)分钟`, timeStr))
		createdTimep := now_time.Add(-time.Duration(min) * time.Minute)
		return createdTimep.Format("2006-01-02 15:04")
	}
	if strings.Contains(timeStr, "小时前") {
		hour, _ := strconv.Atoi(ReParse(`^(\d+)小时`, timeStr))
		createdTimep := now_time.Add(-time.Duration(hour) * time.Hour)
		return createdTimep.Format("2006-01-02 15:04")
	}
	if strings.Contains(timeStr, "今天") {
		return strings.Replace(timeStr, "今天", now_time.Format("2006-01-02"), -1)
	}
	if strings.Contains(timeStr, "月") {
		rp := strings.NewReplacer("月", "-", "日", "")
		return rp.Replace(timeStr)
	}
	return timeStr
}

//目标账号
func GetTargetUidList() []string {
	var uidLi []string
	file, err := os.Open(ExecPath + "/account/target.txt")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		lineText = strings.TrimSpace(lineText)
		lineText = strings.Replace(lineText, "\uFEFF", "", -1)
		uidLi = append(uidLi, lineText)
	}
	return uidLi
}
//获得执行环境地址
func getExecutePath() string {
	path:=os.Args
	log.Info("执行地址,path=",path[0])
	return filepath.Dir(path[0])
}