package purchase

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	// 請購編號
	PurchaseID string `gorm:"primaryKey;column:purchase_id;uuid_generate_v4()type:UUID;" json:"purchase_id,omitempty"`
	// 申請人
	PurchaseUser string `gorm:"column:purchase_user;type:VARCHAR;" json:"purchase_user,omitempty"`
	// 公司名稱
	Company string `gorm:"column:purchase_company;type:VARCHAR;" json:"purchase_company,omitempty"`
	// 部門名稱
	Department string `gorm:"column:purchase_department;type:VARCHAR;" json:"purchase_department,omitempty"`
	// 品名
	Item string `gorm:"column:purchase_item;type:VARCHAR;" json:"purchase_item,omitempty"`
	// 請購數量
	Prquantity int64 `gorm:"column:purchase_prquantity;type:int64;" json:"purchase_prquantity,omitempty"`
	// 價格
	Price int64 `gorm:"column:purchase_price;type:INT4;" json:"purchase_price,omitempty"`
	// 流水號
	Code string `gorm:"->;column:purchase_code;type:VARCHAR;" json:"purchase_code,omitempty"`
	// 是否刪除
	IsDeleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	CreatedBy string `gorm:"column:created_by;type:VARCHAR;" json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `gorm:"column:updated_by;type:VARCHAR;" json:"updated_by,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 請購編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 申請人
	PurchaseUser string `json:"purchase_user,omitempty"`
	// 公司名稱
	Company string `json:"purchase_company,omitempty"`
	// 部門名稱
	Department string `json:"purchase_department,omitempty"`
	// 品名
	Item string `json:"purchase_item,omitempty"`
	// 請購數量
	Prquantity int64 `json:"purchase_prquantity,omitempty"`
	// 價格
	Price int64 ` json:"purchase_price,omitempty"`
	// 流水號
	Code string `json:"purchase_code,omitempty"`
	// 是否刪除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Single return structure file
type Single struct {
	// 請購編號
	PurchaseID string `json:"purchase_id,omitempty"`
	// 申請人
	PurchaseUser string `json:"purchase_user,omitempty"`
	// 公司名稱
	Company string `json:"purchase_company,omitempty"`
	// 部門名稱
	Department string `json:"purchase_department,omitempty"`
	// 品名
	Item string `json:"purchase_item,omitempty"`
	// 請購數量
	Prquantity int64 `json:"purchase_prquantity,omitempty"`
	// 價格
	Price int64 ` json:"purchase_price,omitempty"`
	// 流水號
	Code string `json:"purchase_code,omitempty"`
	// 創建時間
	CreatedAt time.Time `json:"created_at"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新時間
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 申請人
	PurchaseUser string `json:"purchase_user" binding:"required" validate:"required"`
	// 公司名稱
	Company string `json:"purchase_company,omitempty" binding:"required" validate:"required"`
	// 部門名稱
	Department string `json:"purchase_department" binding:"required" validate:"required"`
	// 品名
	Item string `json:"purchase_item" binding:"required" validate:"required"`
	// 請購數量
	Prquantity int64 `json:"purchase_prquantity" binding:"required" validate:"required"`
	// 價格
	Price int64 `json:"purchase_price,omitempty" binding:"required" validate:"required"`
	// 創建者
	CreatedBy string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 請購編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 申請人
	PurchaseUser *string `json:"purchase_user,omitempty" form:"purchase_user"`
	// 公司名稱
	Company *string `json:"purchase_company,omitempty" form:"purchase_company"`
	// 部門名稱
	Department *string `json:"purchase_department,omitempty" form:"purchase_department"`
	// 品名
	Item *string `json:"purchase_item,omitempty" form:"item"`
	// 請購數量
	Prquantity *int64 `json:"purchase_prquantity,omitempty" form:"purchase_prquantity"`
	// 價格
	Price *int64 `json:"purchase_price,omitempty" form:"purchase_price"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Purchase []*struct {
		// 請購編號
		PurchaseID string `json:"purchase_id,omitempty"`
		// 申請人
		PurchaseUser string `json:"purchase_user,omitempty"`
		// 公司名稱
		Company string `json:"purchase_company,omitempty"`
		// 部門名稱
		Department string `json:"purchase_department,omitempty"`
		// 品名
		Item string `json:"purchase_item,omitempty"`
		// 請購數量
		Prquantity int64 `json:"purchase_prquantity,omitempty"`
		// 價格
		Price int64 `json:"purchase_price,omitempty"`
		// 流水號
		Code string `json:"purchase_code,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
	} `json:"purchase"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 請購編號
	PurchaseID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 申請人
	PurchaseUser *string `json:"purchase_user,omitempty"`
	// 公司名稱
	Company string `json:"purchase_company,omitempty"`
	// 部門名稱
	Department *string `json:"purchase_department,omitempty"`
	// 品名
	Item *string `json:"purchase_item,omitempty"`
	// 請購數量
	Prquantity *int64 `json:"purchase_prquantity,omitempty"`
	// 價格
	Price int64 `json:"purchase_price,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" swaggerignore:"true"`
	// 是否刪除
	IsDeleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Tablepurchase_department sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "purchase"
}
