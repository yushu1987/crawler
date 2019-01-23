package service

import (
	"crawler/database/mysql"
	"crawler/database/redis"
	"crawler/lib"
	"crawler/structs"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	task_url_chan     = make(chan string, 1000)
	download_url_chan = make(chan string, 1000)
	download_img_chan = make(chan string, 1000)
	download_vdo_chan = make(chan string, 1000)
	es_chan           = make(chan *structs.Link, 1)

	link = &structs.Link{}
	db = mysql.LinkDB{}
)

func MainHandler() {
	fmt.Println("main handler init")
	task_url_chan <- "http://www.sina.com.cn"
	download_url_chan <- "http://www.sina.com.cn"

	c := make(chan os.Signal)
	signal.Notify(c,
		syscall.SIGKILL,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)
	for {
		select {
		case url := <-task_url_chan:
			url = lib.FixUrl(url)
			link.Url = url
			CrawlerUrl(url)
		case s := <-c:
			fmt.Println("stop", s)
			os.Exit(99)
		}
	}
}

func CrawlerUrl(url string) {
	parse := lib.NewParse()
	parse.NewDocument(url)

	images := parse.GetImages()
	hrefs := parse.GetHrefs()
	videos := parse.GetVideos()
	title := parse.GetTitle()

	task:=structs.Task {
		Url:url,
		Title:title,
		Image:len(images),
		Href:len(hrefs),
		Video: len(videos),
		CreateTime:time.Now().Format("2006-01-02 15:04:05"),
	}
	id, err:=db.Add(task)
	if err != nil {
		fmt.Println("error",err.Error())
	}

	if title != ""  && id > 0{
		link.Title = title
		link.Id = id
		es_chan <- link
	}

	for _, v := range hrefs {
		md5 := lib.Md5sum(v)
		exist, err := redis.IsExist(md5)
		if err != nil || exist == false {
			lib.Log.Infof("get redis exist failed")
			if i, _ := redis.SetKV(md5, true); i > 0 {
				lib.Log.Infof("set new key")
			}
		} else {
			continue
		}

		download_url_chan <- v
		task_url_chan <- v
	}
	for _, v := range images {
		md5 := lib.Md5sum(v)
		exist, err := redis.IsExist(md5)
		if err != nil || exist == false {
			lib.Log.Infof("get redis exist failed")
			if i, _ := redis.SetKV(md5, true); i > 0 {
				lib.Log.Infof("set new key")
			}
		} else {
			continue
		}
		download_img_chan <- v
	}
	for _, v := range videos {
		md5 := lib.Md5sum(v)
		exist, err := redis.IsExist(md5)
		if err != nil || exist == false {
			lib.Log.Infof("get redis exist failed")
			if i, _ := redis.SetKV(md5, true); i > 0 {
				lib.Log.Infof("set new key")
			}
		} else {
			continue
		}
		download_vdo_chan <- v
	}

}
