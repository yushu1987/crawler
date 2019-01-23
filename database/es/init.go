package es

import (
	"context"
	"crawler/lib"
	"github.com/olivere/elastic"
	"time"
)

var client *elastic.Client

func InitEsConn() {
	var err error
	//host:=lib.Config.GetString("es.host")
	host := "http://127.0.0.1:9200"
	client,err =elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(host),elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetHealthcheck(false),
		)
	if err != nil {
		panic("es init failed" + err.Error())
	}
	info ,code , err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic ("ping es failed")
	}
	lib.Log.Infof("ElasticSearch returned with code %d and version %s", code, info.Version.Number)
	client.Flush()
}

func CloseEs() {
	lib.Log.Infof("stop es")
	client.Stop()
}



