package service

import (
	"crawler/database/es"
	"crawler/structs"
	"fmt"
	"strconv"
)

var (
	document  *structs.Link
	esObj = &es.ES{"crawler", "title"}
)


func PutES() {

	go func() {
		for {
			select {
			case document = <-es_chan:
				fmt.Println("put es title", document.Title,document.Id)
				err:=esObj.PutObject(*document, strconv.Itoa(document.Id))
				if err != nil {
					fmt.Println("err:" , err.Error())
				}
			}
		}
	}()
}

func SearchWord(word string) []structs.Link{
	return esObj.SearchTitle(word)
}
