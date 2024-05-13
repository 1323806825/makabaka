package service

import (
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/i-coder-robot/gin-demo/repository"
	uuid "github.com/satori/go.uuid"
)

type BannerService struct {
	Repo repository.BannerRepoInterface
}

type BannerSrv interface {
	Get(Banner model.Banner) (*model.Banner, error)
	Exist(Banner model.Banner) *model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner, error)
	Edit(Banner model.Banner) (bool, error)
	Delete(id string) (bool, error)
}

func (srv *BannerService) Get(Banner model.Banner) (*model.Banner, error) {
	return srv.Get(Banner)
}

func (srv *BannerService) Exist(Banner model.Banner) *model.Banner {
	return srv.Exist(Banner)
}

func (srv *BannerService) ExistByBannerID(id string) *model.Banner {
	return srv.ExistByBannerID(id)
}

func (srv *BannerService) Add(Banner model.Banner) (*model.Banner, error) {
	if Banner.BannerID == "" {
		Banner.BannerID = uuid.NewV4().String()
	}
	return srv.Add(Banner)
}

func (srv *BannerService) Edit(Banner model.Banner) (bool, error) {
	return srv.Edit(Banner)
}

func (srv *BannerService) Delete(id string) (bool, error) {
	return srv.Delete(id)
}
