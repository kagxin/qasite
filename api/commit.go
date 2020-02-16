package api

import (
	"net/http"
	"qasite/middleware"
	"qasite/model"
	"qasite/utils"
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
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": "ok",
			"data":    gin.H{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPParamError,
			"message": err.Error(),
			"data":    gin.H{},
		})
	}
}

func GetComment(c *gin.Context) {
	QID, ok := c.Params.Get("QID")

	if ok {
		var comments []model.Comment
		qid, _ := strconv.ParseUint(QID, 10, 64)
		model.DB.Where("question_id=?", qid).Find(&comments)
		var commentData []gin.H
		for _, comment := range comments {
			commentData = append(commentData, gin.H{
				"id":          comment.ID,
				"text":        comment.ID,
				"question_id": comment.QuestionID,
				"user_id":     comment.UserID,
				"created_at":  comment.CreatedAt,
				"updated_at":  comment.UpdatedAt,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": "ok",
			"data":    commentData,
		})

	}
}
