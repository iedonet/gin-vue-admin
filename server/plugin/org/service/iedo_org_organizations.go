
package service

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/org/model"
    "github.com/flipped-aurora/gin-vue-admin/server/plugin/org/model/request"
)

var IedoOrgOrganizations = new(iedoOrgOrganizations)

type iedoOrgOrganizations struct {}
// CreateIedoOrgOrganizations 创建组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) CreateIedoOrgOrganizations(ctx context.Context, iedoOrgOrganizations *model.IedoOrgOrganizations) (err error) {
	err = global.GVA_DB.Create(iedoOrgOrganizations).Error
	return err
}

// DeleteIedoOrgOrganizations 删除组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) DeleteIedoOrgOrganizations(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&model.IedoOrgOrganizations{},"id = ?",id).Error
	return err
}

// DeleteIedoOrgOrganizationsByIds 批量删除组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) DeleteIedoOrgOrganizationsByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.IedoOrgOrganizations{},"id in ?",ids).Error
	return err
}

// UpdateIedoOrgOrganizations 更新组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) UpdateIedoOrgOrganizations(ctx context.Context, iedoOrgOrganizations model.IedoOrgOrganizations) (err error) {
	err = global.GVA_DB.Model(&model.IedoOrgOrganizations{}).Where("id = ?",iedoOrgOrganizations.Id).Updates(&iedoOrgOrganizations).Error
	return err
}

// GetIedoOrgOrganizations 根据id获取组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) GetIedoOrgOrganizations(ctx context.Context, id string) (iedoOrgOrganizations model.IedoOrgOrganizations, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&iedoOrgOrganizations).Error
	return
}
// GetIedoOrgOrganizationsInfoList 分页获取组织管理记录
// Author [yourname](https://github.com/yourname)
func (s *iedoOrgOrganizations) GetIedoOrgOrganizationsInfoList(ctx context.Context, info request.IedoOrgOrganizationsSearch) (list []model.IedoOrgOrganizations, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.IedoOrgOrganizations{})
    var iedoOrgOrganizationss []model.IedoOrgOrganizations
    // 如果有条件搜索 下方会自动创建搜索语句
  
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	err = db.Find(&iedoOrgOrganizationss).Error
	return  iedoOrgOrganizationss, total, err
}

func (s *iedoOrgOrganizations)GetIedoOrgOrganizationsPublic(ctx context.Context) {

}
