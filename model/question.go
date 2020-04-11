package model

// QuestionReq
type QuestionReq struct {
	Title string `form:"title"` // 问题的title
	Text  string `from:"text"`  // 问题正文
}

type QuestionQueryReq struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}
