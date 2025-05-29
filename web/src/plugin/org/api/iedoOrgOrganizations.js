import service from '@/utils/request'
// @Tags IedoOrgOrganizations
// @Summary 创建组织管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IedoOrgOrganizations true "创建组织管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /iedoOrgOrganizations/createIedoOrgOrganizations [post]
export const createIedoOrgOrganizations = (data) => {
  return service({
    url: '/iedoOrgOrganizations/createIedoOrgOrganizations',
    method: 'post',
    data
  })
}

// @Tags IedoOrgOrganizations
// @Summary 删除组织管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IedoOrgOrganizations true "删除组织管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /iedoOrgOrganizations/deleteIedoOrgOrganizations [delete]
export const deleteIedoOrgOrganizations = (params) => {
  return service({
    url: '/iedoOrgOrganizations/deleteIedoOrgOrganizations',
    method: 'delete',
    params
  })
}

// @Tags IedoOrgOrganizations
// @Summary 批量删除组织管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除组织管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /iedoOrgOrganizations/deleteIedoOrgOrganizations [delete]
export const deleteIedoOrgOrganizationsByIds = (params) => {
  return service({
    url: '/iedoOrgOrganizations/deleteIedoOrgOrganizationsByIds',
    method: 'delete',
    params
  })
}

// @Tags IedoOrgOrganizations
// @Summary 更新组织管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.IedoOrgOrganizations true "更新组织管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /iedoOrgOrganizations/updateIedoOrgOrganizations [put]
export const updateIedoOrgOrganizations = (data) => {
  return service({
    url: '/iedoOrgOrganizations/updateIedoOrgOrganizations',
    method: 'put',
    data
  })
}

// @Tags IedoOrgOrganizations
// @Summary 用id查询组织管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.IedoOrgOrganizations true "用id查询组织管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /iedoOrgOrganizations/findIedoOrgOrganizations [get]
export const findIedoOrgOrganizations = (params) => {
  return service({
    url: '/iedoOrgOrganizations/findIedoOrgOrganizations',
    method: 'get',
    params
  })
}

// @Tags IedoOrgOrganizations
// @Summary 分页获取组织管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取组织管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /iedoOrgOrganizations/getIedoOrgOrganizationsList [get]
export const getIedoOrgOrganizationsList = (params) => {
  return service({
    url: '/iedoOrgOrganizations/getIedoOrgOrganizationsList',
    method: 'get',
    params
  })
}
// @Tags IedoOrgOrganizations
// @Summary 不需要鉴权的组织管理接口
// @Accept application/json
// @Produce application/json
// @Param data query request.IedoOrgOrganizationsSearch true "分页获取组织管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /iedoOrgOrganizations/getIedoOrgOrganizationsPublic [get]
export const getIedoOrgOrganizationsPublic = () => {
  return service({
    url: '/iedoOrgOrganizations/getIedoOrgOrganizationsPublic',
    method: 'get',
  })
}
