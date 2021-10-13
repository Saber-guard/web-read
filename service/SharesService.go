package service

import (
	"web-read/model"
	"web-read/util"
)

type SharesService struct {
}

func (c SharesService) List() []map[string]interface{} {
	var companyList []model.CompanyInfoModel
	DbService.Db.Scopes(util.PaginatorUtil{Page: 1, Size: 100}.Exec).Find(&companyList)
	list := make([]map[string]interface{}, len(companyList))
	for index, item := range companyList {
		list[index] = map[string]interface{}{
			"id":   item.ID,
			"code": item.Code,
			"name": item.Name,
		}
	}
	return list
}
