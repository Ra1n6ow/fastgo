package model

import (
	"gorm.io/gorm"

	"github.com/ra1n6ow/fastgo/internal/pkg/rid"
)

// AfterCreate 在 insert 语句触发之后生成 postID.
func (m *Post) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// AfterCreate 在 insert 语句触发之后生成 userID.
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}
