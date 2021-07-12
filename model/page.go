package model

type Page struct {
	PageNo   int `form:"page_no" json:"-" gorm:"column:-"`
	PageSize int `form:"page_size" json:"-" gorm:"column:-"`
}

func (p *Page) Offset() int {
	if p.PageNo < 1 {
		p.PageNo = 1
	}

	return (p.PageNo - 1) * p.PageSize
}
