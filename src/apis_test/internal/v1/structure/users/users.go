package Users

import (
	model "eirc.app/internal/v1/structure"
)

// Company struct is a row record of the user table in the invoice database
// Table struct is database table struct
type Table struct {
	// 使用者編號
	UserID int `gorm:"primaryKey;column:user_id;type:serial;" json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `gorm:"column:UserName;type:VARCHAR;not null;" json:"UserName,omitempty"`
	// 使用者密碼
	Password string `gorm:"column:password;type:VARCHAR;not null;" json:"password,omitempty"`
	//使用者電子郵件
	Email string `gorm:"column:emial;type:VARCHAR;not null;" json:"email,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 使用者編號
	UserID int `json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	//使用者電子郵件
	Email string ` json:"email,omitempty"`
}

// Single return structure file
type Single struct {
	// 使用者編號
	UserID int `json:"user_id,omitempty"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	//使用者電子郵件
	Email string ` json:"email,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 使用者名稱
	UserName string `json:"UserName,omitempty" binding:"required" validate:"required"`
	// 使用者密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	//使用者電子郵件
	Email string `json:"email,omitempty" binding:"required" validate:"required"`
}

// Updated struct is used to update
type Updated struct {
	// 使用者編號
	UserID int `json:"user_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 使用者名稱
	UserName string `json:"UserName,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	//使用者電子郵件
	Email string ` json:"email,omitempty"`
}

// Field is structure file for search
type Field struct {
	// 使用者編號
	UserID int `json:"user_id,omitempty"  binding:"omitempty,uuid4" swaggerignore:"true"`
	// 使用者名稱
	UserName *string `json:"UserName,omitempty" form:"UserName"`
	// 使用者密碼
	Password *int64 `json:"password,omitempty" form:"password"`
	//使用者電子郵件
	Email string ` json:"email,omitempty" from:"email"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	User []*struct {
		// 使用者編號
		UserID int `json:"user_id,omitempty"`
		// 使用者名稱
		UserName string `json:"UserName,omitempty"`
		// 使用者密碼
		Password string `json:"password,omitempty"`
		//使用者電子郵件
		Email string ` json:"email,omitempty"`
	} `json:"users"`
	model.OutPage
}

// TableUserName sets the insert table UserName for this struct type
func (c *Table) TableUserName() string {
	return "users"
}
