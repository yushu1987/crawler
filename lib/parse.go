package lib

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ParseUrl struct {
	Document *goquery.Document
}

func NewParse() *ParseUrl {
	return &ParseUrl{
	}
}

func (p *ParseUrl) NewDocument(url string) {
	document, err := goquery.NewDocument(url)
	if err != nil {
		Log.Warnf("parse Document failed url:%s",url)
		return
	}
	p.Document = document
}

func (p *ParseUrl) GetTitle() string {
	if p.Document == nil {
		return ""
	}
	return p.Document.Find("title").Text()
}

func (p *ParseUrl) Size() int {
	return p.Document.Size()
}

func (p *ParseUrl) GetHrefs() []string {
	var urlMap = map[string]interface{}{}
	if p.Document == nil {
		return []string{}
	}
	p.Document.Find("a").Each(func(i int, aa *goquery.Selection) {
		href, ok := aa.Attr("href")
		href = strings.TrimSpace(href)
		_, exist := urlMap[href]
		if ok && IsHtml(href) && IsCrawled(href) && exist == false {
			if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "./") {
				href = SamePathUrl(href, href, 1)
			} else if strings.HasPrefix(href, "../") {
				href = SamePathUrl(href, href, 2)
			}
			urlMap[href] = true
		}
	})
	return GetMapKeys(urlMap)
}

func (p *ParseUrl) GetImages() []string {
	var urlMap = make(map[string]interface{})
	if p.Document == nil {
		return []string{}
	}
	p.Document.Find("img").Each(func(i int, aa *goquery.Selection) {
		href, ok := aa.Attr("src")
		href = strings.TrimSpace(href)
		_, exist := urlMap[href]
		if ok && exist == false && IsImage(href) && IsHtml(href){
			if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "./") {
				href = SamePathUrl(href, href, 1)
			} else if strings.HasPrefix(href, "../") {
				href = SamePathUrl(href, href, 2)
			}
			urlMap[href] = true
		}
	})
	return GetMapKeys(urlMap)
}

func (p *ParseUrl) GetVideos() []string {
	var urlMap = make(map[string]interface{})
	if p.Document == nil {
		return []string{}
	}
	p.Document.Find("video").Each(func(i int, aa *goquery.Selection) {
		href, ok := aa.Attr("src")
		href = strings.TrimSpace(href)
		_, exist := urlMap[href]
		if ok && exist == false && IsImage(href) && IsHtml(href){
			if strings.HasPrefix(href, "/") || strings.HasPrefix(href, "./") {
				href = SamePathUrl(href, href, 1)
			} else if strings.HasPrefix(href, "../") {
				href = SamePathUrl(href, href, 2)
			}
			urlMap[href] = true
		}
	})
	return GetMapKeys(urlMap)
}
