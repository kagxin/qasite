package api

import (
	"fmt"
	"net/http"
	"qasite/middleware"
	"qasite/model"
	"qasite/utils"

	"github.com/gin-gonic/gin"
)

// UserValidation for create question
type QuestionValidation struct {
	Text  string `form:"text" binding:"required"`
	Title string `form:"title" binding:"required"`
}

// CreateQuestion for create question
func CreateQuestion(c *gin.Context) {
	var questionVal QuestionValidation

	if err := c.ShouldBindJSON(&questionVal); err == nil {
		user, _ := c.Get(middleware.IdentityKey)
		model.DB.Create(&model.Question{
			Title:  questionVal.Title,
			Text:   questionVal.Text,
			UserID: user.(*model.User).ID,
		})
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": "create sucess",
			"data":    gin.H{},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": err.Error(),
			"data":    gin.H{},
		})
	}

}

func ShowQuestion(c *gin.Context) {
	qID, g := c.Params.Get("QID")
	if g {
		var q model.Question
		result := model.DB.Where("id=?", qID).First(&q)
		fmt.Println(result.RowsAffected)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data": gin.H{
				"title":   q.Title,
				"text":    q.Text,
				"id":      q.ID,
				"user_id": q.UserID,
			},
		})
	} else {
		var qs []model.Question
		model.DB.Find(&qs)
		var questions []gin.H
		for _, question := range qs {
			questions = append(questions, gin.H{
				"id":        question.ID,
				"create_at": question.CreatedAt,
				"update_at": question.UpdatedAt,
				"title":     question.Title,
				"text":      question.Text,
				"user_id":   question.UserID,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": "ok",
			"data":    questions,
		})

	}

}
