package models

import "time"

type AboutUs struct {
	ID          string    `gorm:"column:id;type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string    `gorm:"column:title;type:varchar(255)"`
	Description string    `gorm:"column:description;type:text"`
	Image       string    `gorm:"column:image;type:text"`
	ListOrder   int16     `gorm:"column:list_order;type:int2"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp"`
	DeletedAt   time.Time `gorm:"column:deleted_at;type:timestamp"`
}

type TableColumn struct {
	ColumnName string `gorm:"column:column_name"`
	DataType   string `gorm:"column:data_type"`
}
