package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/wangjiandev/product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	DeleteProductByID(int64) error
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

//NewProductRepository 创建productRepository
func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

//InitTable 初始化表
func (u *ProductRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSize{}, &model.ProductSeo{}).Error
}

//FindProductByID 根据ID查找Product信息
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, u.mysqlDb.Preload("ProductImages").Preload("ProductSizes").Preload("ProductSeo").First(product, productID).Error
}

//CreateProduct 创建Product信息
func (u *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

//DeleteProductByID 根据ID删除Product信息
func (u *ProductRepository) DeleteProductByID(productID int64) error {
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Where("id = ?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("image_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//UpdateProduct 更新Product信息
func (u *ProductRepository) UpdateProduct(product *model.Product) error {
	return u.mysqlDb.Model(product).Update(product).Error
}

//FindAll 获取结果集
func (u *ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, u.mysqlDb.Preload("ProductImages").Preload("ProductSizes").Preload("ProductSeo").Find(&productAll).Error
}
