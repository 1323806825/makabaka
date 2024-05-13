package model

type Banner struct {
	BannerId    string `json:"banner_id" gorm:"column:banner_id"`
	Url         string `json:"url" gorm:"column:url"`
	RedirectUrl string `json:"redirect_url" gorm:"column:redirect_url"`
	Order       int    `json:"order" gorm:"column:order"`
	CreateUser  string `json:"CreateUser" gorm:"column:create_user"`
	UpdateUser  string `json:"updateUser" gorm:"column:update_user"`
	BannerID    string `json:"bannerID" gorm:"column:banner_id"`
}
