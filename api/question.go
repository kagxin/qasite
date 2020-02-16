package api

import (
	"net/http"
	"qasite/middleware"
	"qasite/model"
	"qasite/utils"
	"time"

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
	var title, text, username string
	var questionId, userId uint
	var createdAt, updatedAt time.Time

	if g {
		row := model.DB.Raw("select q.title, q.text, q.id, q.user_id, user.username, q.created_at, q.updated_at from question q left join user on user.id=q.user_id where q.id=?", qID).Row()
		if err := row.Scan(&title, &text, &questionId, &userId, &username, &createdAt, &updatedAt); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"data": gin.H{
					"title":      title,
					"text":       text,
					"id":         questionId,
					"user_id":    userId,
					"username":   username,
					"created_at": createdAt,
					"updated_at": updatedAt,
				},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    utils.HTTPNotFind,
				"message": "not find question",
				"data":    gin.H{},
			})
		}

	} else {
		var qs []model.Question
		model.DB.Find(&qs)
		rows, _ := model.DB.Raw("select q.title, q.text, q.id, q.user_id, user.username " +
			"from question q left join user on user.id=q.user_id").Rows()
		var questions []gin.H
		for rows.Next() {
			rows.Scan(&title, &text, &questionId, &userId, &username)
			questions = append(questions, gin.H{
				"title":      title,
				"text":       text,
				"id":         questionId,
				"user_id":    userId,
				"username":   username,
				"created_at": createdAt,
				"updated_at": updatedAt,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    utils.HTTPOK,
			"message": "ok",
			"data":    questions,
		})

	}

}
