package main

import (
	"flag"
	"fmt"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func cmd(cm ...string) {

	openVolCmd := exec.Command(cm[0], cm[1:]...)
	str, err := openVolCmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(str))
	fmt.Println("=======")
	//filter  line breaks
	fmt.Println(strings.Trim(string(str), "\n"))

}

//定时创建数据库
func timeToCreatDb() {
	//for {
	now := time.Now()                                                                  //获取当前时间，放到now里面，要给next用
	next := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location()) //获取下一个凌晨的日期
	t := time.NewTimer(next.Sub(now))                                                  //计算当前时间到凌晨的时间间隔，设置一个定时器
	<-t.C
	fmt.Print("开始准备: \n", time.Now())
	//以下为定时执行的操作
	go func() {
		tick := time.Tick(1 * time.Second)
		for range tick {
			cmd("/usr/bin/osascript", "-e", "set volume output muted 0")
			cmd("/usr/bin/osascript", "-e", "set volume 3")
		}
	}()
	task(*mp3)

	//time.Sleep(60*time.Second)
	//next.Add(24*time.Hour)
	//next.Sub(now)
	//t = time.NewTimer(next.Sub(now))                                                  //计算当前时间到凌晨的时间间隔，设置一个定时器
	//<-t.C
	//}
}

func task(mp3 string) {
	var err error
	var response *http.Response
	if response, err = http.Get(mp3); err != nil {
		log.Fatal(err)
	}

	var dec *minimp3.Decoder
	if dec, err = minimp3.NewDecoder(response.Body); err != nil {
		log.Fatal(err)
	}
	<-dec.Started()

	log.Printf("Convert audio sample rate: %d, channels: %d\n", dec.SampleRate, dec.Channels)

	var context *oto.Context

	if context, err = oto.NewContext(dec.SampleRate, dec.Channels, 2, 4096); err != nil {
		log.Fatal(err)
	}

	var waitForPlayOver = new(sync.WaitGroup)
	waitForPlayOver.Add(1)

	var player = context.NewPlayer()

	go func() {
		defer response.Body.Close()
		for {
			var data = make([]byte, 512)
			_, err = dec.Read(data)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				break
			}
			player.Write(data)
		}
		log.Println("over play.")
		waitForPlayOver.Done()
	}()

	waitForPlayOver.Wait()

	<-time.After(time.Second)
	dec.Close()
	player.Close()
}

var mp3 *string

func main() {
	mp3 = flag.String("yy", "https://raw.githubusercontent.com/zy84338719/-static/main/gfr.mp3", "yy=map3 地址")
	flag.Parse()
	//cmd("/usr/bin/osascript", "-e", "set volume output muted 0")
	//cmd("/usr/bin/osascript", "-e", "set volume 3")

	timeToCreatDb()
	//task(*mp3)
}
