package dao

import (
	"github.com/go-xorm/xorm"
	"log"
	"superstar/models"
)

type SuperStarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperStarDao {
	return &SuperStarDao{
		engine:engine,
	}
}

func (d *SuperStarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id:id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
		data.Id = 0
		return data
	}
}

func (d *SuperStarDao) GetAll() []models.StarInfo {
	datalist := make([]models.StarInfo, 0)
	log.Println(datalist)
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Fatal("GetAll err=",err)
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperStarDao) Search(country string) []models.StarInfo {
	datalist := []models.StarInfo{}
	err := d.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Print(err)
		return datalist
	} else {
		return datalist
	}
}

func (d *SuperStarDao) Delete(id int) error {
	data := &models.StarInfo{Id:id, SysStatus:1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *SuperStarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *SuperStarDao) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}