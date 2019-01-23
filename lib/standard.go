package lib

import (
	"crawler/structs"
	"strings"
)

func IsCrawled(url string) bool {
	for _, v := range structs.CrawlerType {
		if strings.HasSuffix(url, strings.ToLower(v)) {
			return true
		}
	}
	return false
}

func IsStatus(status int) bool {
	for _, v := range structs.CrawlerStatus {
		if v == status {
			return true
		}
	}
	return false
}

func IsImage(imgUrl string) bool {
	for _, v := range structs.ImageSuffix {
		if b := strings.HasSuffix(imgUrl, v); b == true {
			return true
		}
	}
	return false
}

func IsHtml(str string) bool {
	if strings.HasPrefix(str, "#") || strings.HasSuffix(str, ".exe") || strings.HasSuffix(str, ":void(0);") {
		return false
	} else if strings.HasPrefix(str, "{") && strings.HasSuffix(str, "}") {
		return false
	} else if strings.EqualFold(str, "javascript:;") {
		return false
	} else {
		return true
	}
	return true
}

func Size(s int) float32 {
	return float32(s * 1.0 / 1024)
}

func SamePathUrl(preUrl string, url string, mark int) (newUrl string) {
	last := strings.LastIndex(preUrl, "/")
	if last == 6 {
		newUrl = preUrl + url
	} else {
		if mark == 1 {
			newUrl = preUrl[:last] + url
		} else {
			newPreUrl := preUrl[:last]
			newLast := strings.LastIndex(newPreUrl, "/")
			newUrl = newPreUrl[:newLast] + url
		}
	}
	if i := strings.LastIndex(newUrl, "//"); i > 0 {
		newUrl = newUrl[i+2:]
	}
	return newUrl
}

func FixUrl(url string) string {
	if i := strings.Index(url, "http://"); i == -1 {
		return "http://" + url
	}
	return url
}
