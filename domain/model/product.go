package model

type Product struct{
	ID int64 `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	ProductName string `json:"product_name"`
	ProductSku string `gorm:"uniqueIndex;not null;size:255" json:"product_sku"`
	ProductPrice float64 `json:"product_price"`
	ProductDestination string `json:"product_destination"`
	ProductImage []ProductImage `gorm:"foreignKey:ImageProductID" json:"product_image"`
	ProductSize []ProductSize `gorm:"foreignKey:SizeProductID" json:"product_size"`
	ProductSeo ProductSeo `gorm:"foreignKey:SeoProductID" json:"product_seo"`
}

