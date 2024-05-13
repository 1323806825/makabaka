package respository

import (
	"errors"
	"github.com/i-coder-robot/gin-demo/model"
	"github.com/jinzhu/gorm"
)

type BannerRepository struct{
	DB *gorm.DB
}

type BannerRepoInterface interface{
	Get(Banner model.Banner)(*model.Banner , error)
	Exist(Banner model.Banner)*model.Banner
	ExistByBannerID(id string) *model.Banner
	Add(Banner model.Banner) (*model.Banner,error)
	Edit(Banner model.Banner) (bool ,error)
	Delete(id string) (bool,error)
}

func (b *BannerRepository)Get(banner model.Banner)(*model.Banner,error){
	if err := b.DB.Where(&banner).Find(&banner).Error; err != nil{
		return nil,err
	}
	return &banner , nil
}

func (b *BannerRepository) Exist(banner model.Banner)*model.Banner{
	if banner.Url != "" && banner.RedirectUrl != ""{
		b.DB.Model(&banner).Where("url=? and redirect_url=?",banner.Url,banner.RedirectUrl)
		return &banner
	}
	return nil
}

func (b *BannerRepository) ExistByBannerID(id string)*model.Banner{
	var Banner model.Banner
	b.DB.Where("banner_id = ?",id).First(&Banner)
	return &Banner
}

func (b *BannerRepository) Add(banner model.Banner) (*model.Banner,error){
	exist := b.Exist(banner)
	if exist!= nil && exist.Url == banner.Url && exist.RedirectUrl == banner.BannerID{
		return nil,errors.New("轮播图已经存在")
	}

	err := b.DB.Create(banner).Error
	if err != nil{
		return nil , errors.New("轮播图添加失败")
	}
	return &banner , nil
}

func (b *BannerRepository)Edit(banner model.Banner) (bool,error){
	if banner.BannerID == ""{
		return false,errors.New("参数错误")
	}
	Banner := model.Banner{}
	err := b.DB.Model(&Banner).Where("banner_id=?",banner.BannerID).Update(map[string]interface{}{
		"url" : banner.Url,
		"RedirectUrl" : banner.RedirectUrl,
		"order" : banner.Order,
	}).Error
	if err != nil{
		return false, err
	}
	return true , nil
}

func (b *BannerRepository) Delete(id string) (bool , error){
	err := b.DB.Where("banner_id=?",id).Delete(&model.Banner{}).Error
	if err != nil{
		return false,err
	}
	return true , nil
}