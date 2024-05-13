package handler

import (
	"awesomeProject/enum"
	"awesomeProject/service"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/resp"
	"net/http"
)

type UserHandler struct {
	UserSrv service.UserSrv
}

func (h *UserHandler) GetEntity(result model.User) resp.User {
	return resp.User{
		UserId:    result.UserId,
		NickName:  result.NickName,
		Mobile:    result.Mobile,
		Address:   result.Address,
		IsDeleted: result.IsDeleted,
		IsLocked:  result.IsLocked,
	}
}

func (h *UserHandler) UserInfoHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}

	u := model.User{
		UserId: userId,
	}

	result, err := h.UserSrv.Get(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"entity": entity,
		})
		return
	}

	r := h.GetEntity(*result)
	entity = resp.Entity{
		Code:      http.StatusOK,
		Msg:       "OK",
		Total:     0,
		TotalPage: 0,
		Data:      r,
	}
	c.JSON(http.StatusOK, gin.H{
		"entity": entity,
	})
}

//func (h *UserHandler) UserListHandler(c *gin.Context) {
//	var q query.ListQuery
//	entity := resp.Entity{
//		Code:      int(enum.OperateFail),
//		Msg:       enum.OperateFail.String(),
//		Total:     0,
//		TotalPage: 1,
//		Data:      nil,
//	}
//	err := c.ShouldBindQuery(&q)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"entity": entity,
//		})
//		return
//	}
//
//}

func (h *UserHandler) AddUserHandler(c *gin.Context) {
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	u := model.User{}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}

	r, err := h.UserSrv.Add(u)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.UserId == "" {
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

func (h *UserHandler) EditUserHandler(c *gin.Context) {
	u := model.User{}
	entity := resp.Entity{
		Code:  int(enum.OperateFail),
		Msg:   enum.OperateFail.String(),
		Total: 0,
		Data:  nil,
	}
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"entity": entity,
		})
		return
	}
	b, err := h.UserSrv.Edit(u)
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

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")

	b, err := h.UserSrv.Delete(id)
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
