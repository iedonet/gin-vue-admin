package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

var IedoOrgOrganizations = new(iedoOrgOrganizations)

type iedoOrgOrganizations struct {}

// Init 初始化 组织管理 路由信息
func (r *iedoOrgOrganizations) Init(public *gin.RouterGroup, private *gin.RouterGroup) {
	{
	    group := private.Group("iedoOrgOrganizations").Use(middleware.OperationRecord())
		group.POST("createIedoOrgOrganizations", apiIedoOrgOrganizations.CreateIedoOrgOrganizations)   // 新建组织管理
		group.DELETE("deleteIedoOrgOrganizations", apiIedoOrgOrganizations.DeleteIedoOrgOrganizations) // 删除组织管理
		group.DELETE("deleteIedoOrgOrganizationsByIds", apiIedoOrgOrganizations.DeleteIedoOrgOrganizationsByIds) // 批量删除组织管理
		group.PUT("updateIedoOrgOrganizations", apiIedoOrgOrganizations.UpdateIedoOrgOrganizations)    // 更新组织管理
	}
	{
	    group := private.Group("iedoOrgOrganizations")
		group.GET("findIedoOrgOrganizations", apiIedoOrgOrganizations.FindIedoOrgOrganizations)        // 根据ID获取组织管理
		group.GET("getIedoOrgOrganizationsList", apiIedoOrgOrganizations.GetIedoOrgOrganizationsList)  // 获取组织管理列表
	}
	{
	    group := public.Group("iedoOrgOrganizations")
	    group.GET("getIedoOrgOrganizationsPublic", apiIedoOrgOrganizations.GetIedoOrgOrganizationsPublic)  // 组织管理开放接口
	}
}
