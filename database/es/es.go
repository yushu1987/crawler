package es

import (
	"context"
	"crawler/lib"
	"crawler/structs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
)

type ES struct {
	Index string
	Typ   string
}

func (e *ES) CreateIndex(index, body string) error {
	ck, err := client.CreateIndex(e.Index).Body(body).Do(context.Background())
	if err != nil {
		lib.Log.Infof("create index %s failed", index)
		return err
	}
	if !ck.Acknowledged {
		lib.Log.Infof("create index %s not acknowledged", index)
		return errors.New("not acknowledged")
	}
	return nil
}

func (e *ES) SearchTitle(title string) []structs.Link{
	//title ="体育"
	searchMap:=elastic.NewMatchQuery("title", title)
	searchResult , err := client.Search().Index(e.Index).Type(e.Typ).Query(searchMap).Pretty(true).Do(context.Background())
	if err != nil {
		lib.Log.Warnf("search keywords :%s failed", title)
		return nil
	}
	var list = []structs.Link{}
	if searchResult.TotalHits() > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var item = structs.Link{}
			err:=json.Unmarshal(*hit.Source, &item)
			if err != nil || item.Title == ""{
				lib.Log.Warnf("get es item failed")
			}else {
				list =append(list, item)
			}

		}
	}
	fmt.Println(list)
	return list
}

func (e *ES) PutObject(obj interface{}, id string) error{
	_, err:=client.Index().Index(e.Index).Type(e.Typ).Id(id).BodyJson(obj).Do(context.Background())
	if err != nil {
		lib.Log.Error("put obj to es failed", err.Error())
		return err
	}
	return nil
}