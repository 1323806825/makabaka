package handler

import (
	"awesomeProject/enum"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/resp"
	"net/http"
)

type OrderHandler struct {
	OrderSrv service.OrderSrv
}

func (h *OrderHandler) GetEntity(result model.Order) resp.Order {
	return resp.Order{
		OrderId:     result.OrderId,
		NickName:    result.NickName,
		Mobile:      result.Mobile,
		TotalPrice:  result.TotalPrice,
		PayStatus:   result.PayStatus,
		PayTime:     result.PayTime,
		PayType:     result.PayType,
		OrderStatus: result.OrderStatus,
		ExtraInfo:   result.ExtraInfo,
		UserAddress: result.UserAddress,
		IsDeleted:   result.IsDeleted,
	}
}

func (h *OrderHandler) OrderInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	orderId := c.Param("id")
	if orderId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}
	u := model.Order{
		OrderId: orderId,
	}
	result, err := h.OrderSrv.Get(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}

	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code:      int(enum.OperateOk),
		Msg:       enum.OperateOk.String(),
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{
		"entity": entity,
	})
}

func (h *OrderHandler) AddOrderHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	o := model.Order{}
	err := c.ShouldBindJSON(&o)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}

	r, err := h.OrderSrv.Add(o)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}
	entity.Code = int(enum.OperateOk)
	entity.Msg = enum.OperateOk.String()
	c.JSON(http.StatusOK, gin.H{
		"entity": entity,
	})
}

func (h *OrderHandler) EditOrderHandler(c *gin.Context) {
	o := model.Order{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&o)
	if err != nil || o.OrderId == "" {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}
	b, err := h.OrderSrv.Edit(o)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}
	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
	}
}

func (h *OrderHandler) DeleteOrderHandler(c *gin.Context) {
	id := c.Param("id")
	r := h.OrderSrv.ExistByOrderID(id)
	b, err := h.OrderSrv.Delete(*r)
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}

	if b {
		entity.Code = int(enum.OperateOk)
		entity.Msg = enum.OperateOk.String()
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
	}
}
