package model

type ProductSeo struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	SeoTitle       string `json:"seo_title,omitempty"`
	SeoKeywords    string `json:"seo_keywords,omitempty"`
	SeoDescription string `json:"seo_description,omitempty"`
	SeoCode        string `json:"seo_code,omitempty"`
	SeoProductId   int64  `json:"seo_product_id"`
}
