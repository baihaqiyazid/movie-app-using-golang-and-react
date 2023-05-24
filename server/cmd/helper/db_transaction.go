package helper

import "gorm.io/gorm"

func DBTransaction(tx *gorm.DB) {
	if err := recover(); err != nil {
		tx.Rollback()
	}else {
		tx.Commit()
	}
}