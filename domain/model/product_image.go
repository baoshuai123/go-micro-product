package model

type ProductImage struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ImageName string `json:"image_name"`
	ImageCode string `gorm:"uniqueIndex;not null;size:255" json:"image_code"`
	ImageUrl string `json:"image_url"`
	ImageProductID int64 `json:"image_product_id"`
}
