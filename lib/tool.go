package lib

import (
	"crypto/md5"
	"fmt"
	"shorturl/lib"
	"strings"
)

func GetMapKeys(m map[string]interface{})[]string {
	var list =[]string{}
	for k, _:=range m {
		list = append(list, k)
	}
	return list
}


func GetDownloadPath(url ,typ string) string{
	li := strings.LastIndex(url, ".")
	suffix := url[li+1:]
	filename :=lib.Md5sum(url[:li])
	if typ == "html" {
		return fmt.Sprintf("download/html/%s.%s", filename , suffix)
	}else if typ == "image" {
		return fmt.Sprintf("download/image/%s.%s", filename, suffix)
	}else if typ == "video" {
		return fmt.Sprintf("download/video/%s.%s", filename, suffix)
	}
	return ""
}

func Md5sum(s string) string {
	data := []byte(s)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
