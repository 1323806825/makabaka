package service

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/repository"
	uuid "github.com/satori/go.uuid"
)

type ProductService struct {
	Repo repository.ProductRepoInterface
}
type ProductSrv interface {
	Get(product model.Product) (*model.Product, error)
	Exist(product model.Product) *model.Product
	ExistByProductID(id string) *model.Product
	Add(product model.Product) (*model.Product, error)
	Edit(product model.Product) (bool, error)
	Delete(id string) (bool, error)
}

func (srv *ProductService) Get(product model.Product) (*model.Product, error) {
	return srv.Repo.Get(product)
}

func (srv *ProductService) Exist(product model.Product) *model.Product {
	return srv.Repo.Exist(product)
}

func (srv *ProductService) ExistByProductID(id string) *model.Product {
	return srv.Repo.ExistByProductID(id)
}

func (srv *ProductService) Add(product model.Product) (*model.Product, error) {
	if product.ProductId == "" {
		product.ProductId = uuid.NewV4().String()
	}
	product.IsDeleted = false
	return srv.Repo.Add(product)
}

func (srv *ProductService) Edit(product model.Product) (bool, error) {
	p := srv.ExistByProductID(product.ProductId)
	if p != nil || p.ProductName == "" {
		return false, errors.New("参数错误")
	}
	return srv.Repo.Edit(product)
}

func (srv *ProductService) Delete(id string) (bool, error) {
	if id == "" {
		return false, errors.New("参数错误")
	}
	p := srv.ExistByProductID(id)
	if p == nil {
		return false, errors.New("参数错误")
	}
	p.IsDeleted = !p.IsDeleted
	return srv.Repo.Delete(*p)
}
