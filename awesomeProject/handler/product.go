package handler

import (
	"awesomeProject/enum"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/resp"
	"net/http"
)

type ProductHandler struct {
	ProductSrv service.ProductSrv
}

func (h *ProductHandler) GetEntity(result model.Product) resp.Product {
	return resp.Product{
		ProductId:            result.ProductId,
		ProductName:          result.ProductName,
		ProductIntro:         result.ProductIntro,
		CategoryId:           result.CategoryId,
		ProductCoverImg:      result.ProductCoverImg,
		ProductBanner:        result.ProductBanner,
		OriginalPrice:        result.OriginalPrice,
		SellingPrice:         result.SellingPrice,
		StockNum:             result.StockNum,
		Tag:                  result.Tag,
		SellStatus:           result.SellStatus,
		ProductDetailContent: result.ProductDetailContent,
		IsDeleted:            result.IsDeleted,
	}
}

func (h *ProductHandler) ProductInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}

	u := model.Product{
		ProductId: productId,
	}
	result, err := h.ProductSrv.Get(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}

	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code:      int(enum.OperateOk),
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{
		"entity": entity,
	})
}
func (h *ProductHandler) AddProductHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	p := model.Product{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}

	r, err := h.ProductSrv.Add(p)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.ProductId == "" {
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
	return
}

func (h *ProductHandler) EditProductHandler(c *gin.Context) {
	p := model.Product{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}
	b, err := h.ProductSrv.Edit(p)
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
		return
	}
}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	b, err := h.ProductSrv.Delete(id)
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
