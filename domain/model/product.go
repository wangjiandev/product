package model

type Product struct {
	ID                 int64          `gorm:"primary_key;not_null;auto_increment" json:"id"`
	ProductName        string         `json:"product_name,omitempty"`
	ProductSku         string         `gorm:"unique_index;not_null" json:"product_sku,omitempty"`
	ProductPrice       float64        `json:"product_price,omitempty"`
	ProductDescription string         `json:"product_description,omitempty"`
	ProductImages      []ProductImage `gorm:"ForeignKey:ImageProductId" json:"product_images,omitempty"`
	ProductSizes       []ProductSize  `gorm:"ForeignKey:SizeProductId" json:"product_sizes,omitempty"`
	ProductSeo         ProductSeo     `gorm:"ForeignKey:SeoProductId" json:"product_seo,omitempty"`
}
