package model

type ProductImage struct {
	ID             int64  `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ImageName      string `json:"image_name,omitempty"`
	ImageCode      string `gorm:"unique_index;not_null" json:"image_code,omitempty"`
	ImageUrl       string `json:"image_url,omitempty"`
	ImageProductId int64  `json:"image_product_id"`
}
