package model

type ProductSize struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	SizeName string `json:"size_name"`
	SizeCode string `gorm:"uniqueIndex;not null;size:255" json:"size_code"`
	SizeProductID int64 `json:"size_product_id"`
}
