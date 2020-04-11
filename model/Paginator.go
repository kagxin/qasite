package model

type Paginator struct {
	Page     int         `json:"page"` //当前页码
	PageSize int         `json:"page_size"`
	Total    int         `json:"Total"`
	Data     interface{} `json:"data"`
}
