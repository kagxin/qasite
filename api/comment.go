package api

import (
	"qasite/errno"
	"qasite/model"
	"qasite/utils/response"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateComment(c *gin.Context) {
	var commentReq model.CommentReq
	if err := c.ShouldBind(&commentReq); err != nil {
		response.JSON(c, errno.RequestParamErr, nil)
	}
	questionId := c.Param("QID")
	if s.Mysql.DB.Where("id=?", questionId).First(&model.Question{}).RecordNotFound() {
		response.JSON(c, errno.NotFound, nil)
		return
	}
	user := c.MustGet("user").(model.User)
	comment := model.Comment{
		UserID: user.ID,
		Text:   commentReq.Text,
	}
	s.Mysql.DB.Create(&comment)
	response.JSON(c, errno.Success, comment)
	return
}

func (s *Service) Comment(c *gin.Context) {

	questionId := c.Param("QID")
	if s.Mysql.DB.Where("id=?", questionId).First(&model.Question{}).RecordNotFound() {
		response.JSON(c, errno.NotFound, nil)
		return
	}
	qq := model.CommentQueryReq{Page: 1, PageSize: 20}
	if err := c.ShouldBind(&qq); err != nil {
		response.JSON(c, errno.RequestParamErr, nil)
	}
	var comments []model.Comment
	var count int
	done := make(chan bool, 1)
	go func() {
		s.Mysql.DB.Model(model.Comment{}).Count(&count)
		done <- true
	}()
	s.Mysql.DB.Offset(qq.PageSize * (qq.Page - 1)).Limit(qq.PageSize).Find(&comments)
	<-done

	response.JSON(c, errno.Success, model.Paginator{
		PageSize: qq.PageSize,
		Page:     qq.Page,
		Total:    count,
		Data:     comments,
	})
}
