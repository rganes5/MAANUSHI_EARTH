package repository

import (
	"context"
	"fmt"

	domain "github.com/rganes5/maanushi_earth_e-commerce/pkg/domain"
	interfaces "github.com/rganes5/maanushi_earth_e-commerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type cartDatabase struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartDatabase{DB}
}

func (c *cartDatabase) FindCartById(ctx context.Context, id uint) (domain.Cart, error) {
	var cart domain.Cart
	if err := c.DB.Where("user_id=?", id).Find(&cart).Error; err != nil {
		return cart, err
	}
	return cart, nil
}

//	func (c *cartDatabase) FindProductDetailsById(ctx context.Context, id string) (domain.Products, domain.ProductDetails, error) {
//		var ProductDetails domain.ProductDetails
//		var product domain.Products
//		if err := c.DB.Where("product_id=?", id).Find(&ProductDetails).Error; err != nil {
//			return products, ProductDetails, err
//		}
//		return products, ProductDetails, nil
//	}

// func (c *cartDatabase) FindProductDetailsById(ctx context.Context, id string) (domain.Products, domain.ProductDetails, error) {
// 	var productDetails domain.ProductDetails
// 	var product domain.Products

// 	if err := c.DB.Preload("Product").Where("product_id = ?", id).Find(&productDetails).Error; err != nil {
// 		return product, productDetails, err
// 	}

// 	return product, productDetails, nil
// }

func (c *cartDatabase) FindProductDetailsById(ctx context.Context, id string) (domain.ProductDetails, error) {
	var productDetails domain.ProductDetails
	if err := c.DB.Where("product_id = ?", id).Find(&productDetails).Error; err != nil {
		return productDetails, err
	}

	return productDetails, nil
}

func (c *cartDatabase) FindProductById(ctx context.Context, productId string) (domain.Products, error) {
	var product domain.Products
	if err := c.DB.Where("id=?", productId).Find(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (c *cartDatabase) FindDuplicateProduct(ctx context.Context, productId string, cartID uint) (domain.CartItem, error) {
	var duplicateItem domain.CartItem
	if err := c.DB.Where("product_id=? and id=?", productId, cartID).Find(&duplicateItem).Error; err != nil {
		return duplicateItem, err
	}
	return duplicateItem, nil
}

//result := c.DB.Where("product_detail_id=$1 and cart_id=$2", id, cartid).Find(&exsistitem)

func (c *cartDatabase) UpdateCartItem(ctx context.Context, existingItem domain.CartItem) error {
	var grantTotal int
	tx := c.DB.Begin()
	if err := tx.Model(&domain.CartItem{}).Where("id=?", existingItem.ID).UpdateColumns(&existingItem).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.CartItem{}).Where("cart_id=?", existingItem.CartID).Select("SUM(total_price)").Scan(&grantTotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.Cart{}).Where("id=?", existingItem.CartID).UpdateColumn("grand_total", grantTotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (c *cartDatabase) AddNewItem(ctx context.Context, newItem domain.CartItem) error {
	var grantTotal int
	tx := c.DB.Begin()
	if err := tx.Create(&newItem).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.CartItem{}).Where("cart_id=?", newItem.CartID).Select("SUM(total_price)").Scan(&grantTotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&domain.Cart{}).Where("id=?", newItem.CartID).UpdateColumn("grand_total", grantTotal).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("added item is", newItem)
	return nil
}
