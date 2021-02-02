package model

type ProductSeo struct {
	ID int64 `gorm:"primaryKey;not null;autoIncrement" json:"id"`
	SeoTitle string `json:"seo_title"`
	SeoKeywords string `json:"seo_keywords"`
	SeoDescription string `json:"seo_description"`
	SeoCode string `json:"seo_code"`
	SeoProductID int64 `json:"seo_product_id"`
}
