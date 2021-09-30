package service

import (
	"web-read/model"
	"web-read/util"
)

type SharesService struct {
}

func (c SharesService) List() []model.CompanyInfoModel {
	var companyList []model.CompanyInfoModel
	DbService.Db.Scopes(util.PaginatorUtil{Page: 1, Size: 100}.Exec).Find(&companyList)
	return companyList
}
