package respository

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/jinzhu/gorm"
)

type OrderRespository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	Get(Order model.Order) (*model.Order, error)
	Exist(Order model.Order) *model.Order
	ExistByOrder(id string) *model.Order
	Add(Order model.Order) (*model.Order, error)
	Edit(Order model.Order) (bool, error)
	Delete(Order model.Order) (bool, error)
}

func (o *OrderRespository) Get(order model.Order) (*model.Order, error) {
	if err := o.DB.Where(&order).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderRespository) Exist(order model.Order) *model.Order {
	if order.OrderId != "" {
		o.DB.Model(&order).Where("order_id=?", order.OrderId)
		return &order
	}
	return nil
}

func (o *OrderRespository) ExistByOrder(id string) *model.Order {
	var order model.Order
	o.DB.Where("order_id=?", order.OrderId).First(&order)
	return &order
}

func (o *OrderRespository) Add(order model.Order) (*model.Order, error) {
	err := o.DB.Create(order).Error
	if err != nil {
		return &order, errors.New("订单添加失败")
	}
	return &order, nil
}

func (o *OrderRespository) Edit(order model.Order) (bool, error) {
	if order.OrderId == "" {
		return false, errors.New("传入参数错误")
	}
	t := &model.Order{}
	err := o.DB.Model(t).Where("order_id=?", order.OrderId).Updates(map[string]interface{}{
		"nickname":     order.NickName,
		"mobile":       order.Mobile,
		"pay_status":   order.PayStatus,
		"order_status": order.OrderStatus,
		"extra_info":   order.ExtraInfo,
		"user_address": order.UserAddress,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *OrderRespository) Delete(order *model.Order) (bool, error) {
	err := o.DB.Model(&order).Where("order_id=?", order.OrderId).Update("is_deleted", order.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
