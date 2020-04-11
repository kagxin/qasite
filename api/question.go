package api

import (
	"github.com/gin-gonic/gin"
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"
)

func (s *Service) CreateQuestion(c *gin.Context) {
	var q model.QuestionReq
	if err := c.ShouldBind(&q); err != nil {
		response.JSON(c, errno.RequestParamErr, nil)
		return
	}
	user := c.MustGet("user").(model.User)
	question := model.Question{
		Title:  q.Title,
		Text:   q.Text,
		UserID: user.ID,
	}
	s.DB.Create(&question)
	response.JSON(c, errno.Success, question)
}

func (s *Service) Question(c *gin.Context) {
	qq := model.QuestionQueryReq{Page: 1, PageSize: 20}
	if err := c.ShouldBind(&qq); err != nil {
		response.JSON(c, errno.RequestParamErr, nil)
	}
	var questions []model.Question
	var count int
	done := make(chan bool, 1)
	go func() {
		s.DB.Model(model.Question{}).Count(&count)
		done <- true
	}()
	s.DB.Offset(qq.PageSize * (qq.Page - 1)).Limit(qq.PageSize).Find(&questions)
	<-done

	response.JSON(c, errno.Success, model.Paginator{
		PageSize: qq.PageSize,
		Page:     qq.Page,
		Total:    count,
		Data:     questions,
	})
}

func (s *Service) QuestionDetail(c *gin.Context) {
	questionId := c.Param("QID")
	var question model.Question
	if s.DB.Where("id=?", questionId).First(&question).RecordNotFound() {
		response.JSON(c, errno.NotFound, nil)
		return
	}
	response.JSON(c, errno.Success, question)
	return
}
