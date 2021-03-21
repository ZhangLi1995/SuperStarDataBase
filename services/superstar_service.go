package services

import (
	"superstar/dao"
	"superstar/datasource"
	"superstar/models"
)

type SuperStarService interface {
	GetAll() []models.StarInfo
	Search(country string) []*models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(star *models.StarInfo, columns []string) error
	Create(star *models.StarInfo) error
}

type superStarService struct {
	dao *dao.SuperStarDao
}

func NewSuperStarService() SuperStarService {
	return &superStarService{
		dao: dao.NewSuperStarDao(datasource.InstanceMaster()),
	}
}

func (s *superStarService) GetAll() []models.StarInfo {
	return s.dao.GelAll()
}

func (s *superStarService) Search(country string) []*models.StarInfo {
	return s.dao.Search(country)
}

func (s *superStarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superStarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superStarService) Update(star *models.StarInfo, columns []string) error {
	return s.dao.Update(star, columns)
}

func (s *superStarService) Create(star *models.StarInfo) error {
	return s.dao.Create(star)
}
