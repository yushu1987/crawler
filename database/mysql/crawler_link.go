package mysql

import (
	"crawler/structs"
	"fmt"
	"github.com/jinzhu/gorm"
)

type LinkDB struct {
	table string
}

func (s *LinkDB) SetTable() {
	s.table = "crawler_link"
}

func (s *LinkDB) Add(data structs.Task) (int, error) {
	s.SetTable()
	if err := db.Table(s.table).Create(&data).Error; nil != err {
		fmt.Println(err.Error())
		return 0, err
	}
	return data.Id, nil
}

func (s *LinkDB) GetById(Id int) (structs.Task, error) {
	var data = structs.Task{}
	s.SetTable()
	err := db.Table(s.table).Where("id = ?", Id).Last(&data).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return data, err
	}
	return data, nil
}
