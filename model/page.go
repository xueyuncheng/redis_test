package model

type Page struct {
	PageNo   int `form:"page_no" json:"page_no"`
	PageSize int `form:"page_size" json:"page_size"`
}

func (p *Page) Offset() int {
	if p.PageNo < 1 {
		p.PageNo = 1
	}

	return (p.PageNo - 1) * p.PageSize
}
