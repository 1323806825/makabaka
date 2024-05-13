package model

type Category struct {
	CategoryId string `json:"categoryId" gorm:"column:category_id"`
	Name       string `json:"name" gorm:"column:name"`
	Desc       string `json:"desc" gorm:"column:desc"`
	Order      string `json:"order" gorm:"column:order"`
	ParentId   string `json:"parentId" gorm:"column:parent_id"`
	IsDeleted  bool   `json:"isDeleted" gorm:"column:is_deleted"`
}

type CategoryResult struct{
	C1CategoryID string `gorm:"c1_category_id"`
	C1Name string `gorm:"c1_name"`
	C1Desc string `gorm:"c1_desc"`
	C1Order int `gorm:"c1_order"`
	C1ParentID string `gorm:"c1_parent_id"`

	C2CategoryID string `gorm:"c2_category_id"`
	C2Name string `gorm:"c2_name"`
	C2Desc string `gorm:"c2_desc"`
	C2Order int `gorm:"c2_order"`
	C2ParentID string `gorm:"c2_parent_id"`

	C3CategoryID string `gorm:"c3_category_id"`
	C3Name string `gorm:"c3_name"`
	C3Desc string `gorm:"c3_desc"`
	C3Order int `gorm:"c3_order"`
	C3ParentID string `gorm:"c3_parent_id"`
}



