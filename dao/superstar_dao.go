package dao

import (
	"log"
	"superstar/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type SuperStarDao struct {
	engine *xorm.Engine
}

func NewSuperStarDao(engine *xorm.Engine) *SuperStarDao {
	return &SuperStarDao{
		engine: engine,
	}
}

func (d *SuperStarDao) Get(id int) *models.StarInfo {
	result := &models.StarInfo{Id: id}
	ok, err := d.engine.Where("sys_status=1").Get(result)
	if err != nil {
		log.Fatal(err)
	} else if !ok {
		return nil
	}
	return result
}

func (d *SuperStarDao) GelAll() []models.StarInfo {
	result := make([]models.StarInfo, 0)
	err := d.engine.Desc("id").Where("sys_status=?", 1).Find(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *SuperStarDao) Search(country string) []*models.StarInfo {
	result := make([]*models.StarInfo, 0)
	err := d.engine.Where("country=? AND sys_status=1", country).Desc("id").Find(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (d *SuperStarDao) Create(star *models.StarInfo) error {
	_, err := d.engine.Insert(star)
	return err
}

func (d *SuperStarDao) Update(star *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(star.Id).MustCols(columns...).Update(star)
	return err
}

func (d *SuperStarDao) Delete(id int) error {
	data := &models.StarInfo{
		Id:        id,
		SysStatus: 1,
	}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}
