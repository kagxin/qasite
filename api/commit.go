package api

import (
	"net/http"
	"qasite/middleware"
	"qasite/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Comment struct {
	Text string `form:"text" binding:"required"`
}

func CreateComment(c *gin.Context) {

	var comment Comment
	QID, _ := c.Params.Get("QID")
	if err := c.ShouldBindJSON(&comment); err == nil {
		user, _ := c.Get(middleware.IdentityKey)
		qid, _ := strconv.ParseUint(QID, 10, 64)
		model.DB.Create(&model.Comment{
			Text:       comment.Text,
			QuestionID: uint(qid),
			UserID:     user.(*model.User).ID,
		})
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	} else {
		c.JSON(http.StatusOK, err.Error())
	}
}

func GetComment(c *gin.Context) {
	QID, ok := c.Params.Get("QID")

	if ok {
		var comments []model.Comment
		qid, _ := strconv.ParseUint(QID, 10, 64)
		model.DB.Where("question_id=?", qid).Find(&comments)
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    comments,
		})

	}
}
