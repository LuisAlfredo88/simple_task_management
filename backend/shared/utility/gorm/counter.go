package gorm

import (
	"gorm.io/gorm"
)

type QueryCount struct {
	TotalRecords uint64
}

func GetCount(db *gorm.DB, query *gorm.DB) int64 {
	var queryData QueryCount

	err := db.Raw("SELECT COUNT(*) AS total_records FROM  (?) ",
		query,
	).First(&queryData).Error

	if err != nil {
		return 0
	}

	return int64(queryData.TotalRecords)
}
