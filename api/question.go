package api

import (
	"fmt"
	"net/http"
	"qasite/middleware"
	"qasite/model"

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
			"message": "create sucess",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
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
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    qs,
		})

	}

}
