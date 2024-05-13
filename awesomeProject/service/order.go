package service

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/repository"
)

type OrderService struct {
	Repo repository.OrderRepoInterface
}

type OrderSrv interface {
	Exist(Order model.Order) *model.Order
	Get(Order model.Order) (*model.Order, error)
	ExistByOrderID(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(order model.Order) (bool, error)
}

func (srv *OrderService) Exist(Order model.Order) *model.Order {
	return srv.Repo.Exist(Order)
}

func (srv *OrderService) Get(order model.Order) (*model.Order, error) {
	return srv.Repo.Get(order)
}

func (srv *OrderService) ExistByOrderID(id string) *model.Order {
	panic("implement me")
}

func (srv *OrderService) Add(Order model.Order) (*model.Order, error) {
	return srv.Repo.Add(Order)
}

func (srv *OrderService) Edit(order model.Order) (bool, error) {
	o := srv.Repo.ExistByOrderID(order.OrderId)
	if o == nil || o.Mobile == "" {
		return false, errors.New("参数错误")
	}
	return srv.Repo.Edit(order)
}

func (srv *OrderService) Delete(order model.Order) (bool, error) {
	order.IsDeleted = !order.IsDeleted
	return srv.Repo.Delete(order)
}
