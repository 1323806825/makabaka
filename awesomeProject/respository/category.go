package respository

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/jinzhu/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

type CategoryRepoInterface interface {
	Get(id string) ([]*model.CategoryResult, error)
	Exist(Category model.Category) *model.Category
	ExistByCategoryID(id string) *model.Category
	Add(Category model.Category) (*model.Category, error)
	Edit(Category model.Category) (bool, error)
	Delete(Category model.Category) (bool, error)
}

func (c *CategoryRepository) Get(id string) ([]*model.CategoryResult, error) {
	var list []*model.CategoryResult
	err := c.DB.Raw("select c1.category_id as c1_category_id,c1,name as c1_name,"+
		"c2.desc as c1_desc,c1.order as c1_order,c2.category_id as c2_category_id,"+
		"c2.name as c2_name,c2.order as c2_order,"+
		"c3.category_id as c3_category_d , c3.name as c3_name,"+
		"c3.order as c3_order from "+
		"category c1 join category c2 on c1.category_id = c2.parent_id "+
		"join category c3 on c2.category_id = c3.parent_id"+
		"where c3.category_id = ?", id).Find(&list).Error
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (c *CategoryRepository) Exist(Category model.Category) *model.Category {
	var category model.Category
	if category.Name != "" {
		c.DB.Model(&category).Where("name=?", Category.Name)
		return &category
	}
	return nil
}

func (c *CategoryRepository) ExistByCategory(id string) *model.Category {
	var category model.Category
	c.DB.Where("category_id = ?", id).First(&category)
	return &category
}

func (c *CategoryRepository) Add(Category model.Category) (*model.Category, error) {
	err := c.DB.Create(Category).Error
	if err != nil {
		return nil, errors.New("分类添加失败")
	}
	return &Category, nil
}

func (c *CategoryRepository) Edit(category model.Category) (bool, error) {
	if category.CategoryID == "" {
		return false, errors.New("参数错误")
	}
	temp := &model.Category{
		CategoryID: category.CategoryID,
	}
	err := c.DB.Model(temp).Where("category_id=?", category.CategoryID).Updates(map[string]interface{}{
		"name":      category.Name,
		"order":     category.Order,
		"parent_id": category.CategoryID,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CategoryRepository) Delete(category model.Category) (bool, error) {
	err := c.DB.Model(&category).Where("category_id=?", category.CategoryID).Update("is_deleted", category.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
