package respository

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/jinzhu/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

type ProductRepoInterface interface {
	//List(req *query.ListQuery) (Products []*model.Product, err error)
	//GetTotal(req *query.ListQuery) (total int64, err error)
	Get(Product model.Product) (*model.Product, error)
	Exist(Product model.Product) *model.Product
	ExistByProductId(id string) *model.Product
	Add(Product model.Product) (*model.Product, error)
	Edit(Product model.Product) (bool, error)
	Deleted(Product model.Product) (bool, error)
}

func (repo *ProductRepository) Get(Product model.Product)(*model.Product,error){
	if err := repo.DB.Where(&Product).Find(&Product).Error;err!=nil{
		return nil ,err
	}
	return &Product,nil
}

func (repo *ProductRepository) Exist(Product model.Product) *model.Product{
	if Product.ProductName != ""{
		var temp model.Product
		repo.DB.Where("product_name = ?",Product.ProductName).First(&temp)
		return &temp
	}
	return nil
}

func (repo *ProductRepository) ExistByProduct(id string) *model.Product{
	var p model.Product
	repo.DB.Where("product_id = ?",id).First(&p)
	return &p
}

func (repo *ProductRepository) Add(Product model.Product)(*model.Product, error){
	exist := repo.Exist(Product)
	if exist != nil && exist.ProductName!= ""{
		return &Product,errors.New("商品已经存在")
	}
	err := repo.DB.Create(Product).Error
	if err != nil{
		return nil, errors.New("商品添加失败")
	}
	return &Product,nil
}

func (repo *ProductRepository) Edit(Product model.Product)(bool,error){
	if Product.ProductId == ""{
		return false, errors.New("请输入正确的产品ID")
	}
	p:= model.Product{}
	err := repo.DB.Model(p).Where("product_id=?",Product.ProductId).Updates(map[string]interface{}{
		"product_name":Product.ProductName,
		"Product_intro" : Product.ProductIntro,
		"category_id" : Product.CategoryId,
		"product_cover_img" : Product.ProductCoverImg,
		"product_banner" : Product.ProductBanner,
		"original_price" : Product.OriginalPrice,
		"selling_price" : Product.SellingPrice,
		"stock_num" : Product.StockNum,
		"tag" : Product.Tag,
		"sell_status" : Product.SellStatus,
		"product_detail_content" : Product.ProductDetailContent,
	}).Error
	if err!=nil{
		return false,err
	}
	return true , nil
}

func (repo *ProductRepository) Deleted(p model.Product)(bool,error){
	err := repo.DB.Model(&p).Where("product_id=?",p.ProductId).Update("is_deleted=?",p.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}




