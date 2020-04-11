package model

// Comment
type CommentReq struct {
	Text string `from:"text"`
}

type CommentQueryReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
