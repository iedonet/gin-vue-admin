package model

import (
	"time"
)

// IedoOrgOrganizations 组织管理 结构体
type IedoOrgOrganizations struct {
	Id           int        `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name         string     `json:"name" binding:"required" gorm:"column:name;size:100;not null"`
	Code         string     `json:"code" binding:"required" gorm:"column:code;size:32;unique;not null"`
	Type         string     `json:"type" binding:"required" gorm:"column:type;size:16;not null"`
	ParentId     int        `json:"parentId" gorm:"column:parent_id;default:0"`
	Path         string     `json:"path" gorm:"column:path;size:512"`
	Level        int        `json:"level" gorm:"column:level;default:1"`
	Logo         string     `json:"logo" gorm:"column:logo;size:255"`
	Status       string     `json:"status" gorm:"column:status;size:10;default:'1'"`
	ContactName  string     `json:"contactName" gorm:"column:contact_name;size:50"`
	ContactPhone string     `json:"contactPhone" gorm:"column:contact_phone;size:20"`
	Address      string     `json:"address" gorm:"column:address;size:255"`
	Description  string     `json:"description" gorm:"column:description"`
	Sort         int        `json:"sort" gorm:"column:sort;default:0"`
	CreatedBy    int        `json:"createdBy" gorm:"column:created_by"`
	UpdatedBy    int        `json:"updatedBy" gorm:"column:updated_by"`
	CreateTime   time.Time  `json:"createTime" gorm:"column:create_time;autoCreateTime"`
	UpdateTime   time.Time  `json:"updateTime" gorm:"column:update_time;autoUpdateTime"`
	DeleteReason string     `json:"deleteReason" gorm:"column:delete_reason;size:255"`
	DeleteBy     int        `json:"deleteBy" gorm:"column:delete_by"`
	DeleteTime   *time.Time `json:"deleteTime" gorm:"column:delete_time"`
}

// TableName 组织管理 IedoOrgOrganizations自定义表名 iedo_org_organizations
func (IedoOrgOrganizations) TableName() string {
	return "iedo_org_organizations"
}
