package service

import (
	"crawler/lib"
	"fmt"
)

func Download() {
	var (
		url  string
		path string
	)
	go func() {
		fmt.Println("download init")
		for {
			select {
			case url = <-download_url_chan:
				path = lib.GetDownloadPath(url, "html")
				m := lib.Download(url, path)
				link.Cost = m["cost"].(float64)
				link.Status = m["status"].(int)
				link.Size = m["size"].(int)
			case url = <-download_img_chan:
				path = lib.GetDownloadPath(url, "image")
				lib.Download(url,path)
			case url = <-download_vdo_chan:
				fmt.Println("download video url", url)
				path = lib.GetDownloadPath(url, "video")
			}
		}
	}()

}
