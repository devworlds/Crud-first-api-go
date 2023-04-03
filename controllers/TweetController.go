package controllers

import (
	entities "api/entities"

	"net/http"

	"github.com/gin-gonic/gin"
)

type tweetController struct {
	tweets []entities.Tweet
}

func NewTweetController() *tweetController {
	return &tweetController{}
}

func (t *tweetController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, t.tweets)
}

func (t *tweetController) Create(ctx *gin.Context) {
	tweet := entities.NewTweet()

	if err := ctx.BindJSON(&tweet); err != nil {
		return
	}

	t.tweets = append(t.tweets, *tweet)

	ctx.JSON(http.StatusCreated, tweet)
}

func (t *tweetController) FindOne(ctx *gin.Context) {
	id := ctx.Param("id")
	founded := false

	for i, v := range t.tweets {
		if v.ID == id {
			ctx.JSON(http.StatusOK, t.tweets[i])
			founded = true
		}
	}

	if !founded {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": "Tweet not found",
		})
	}

}

func (t *tweetController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, v := range t.tweets {
		if v.ID == id {
			t.tweets = append(t.tweets[0:i], t.tweets[i+1:]...)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{
		"error": "Tweet not found",
	})
}

func (t *tweetController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, v := range t.tweets {
		if v.ID == id {
			if err := ctx.BindJSON(&t.tweets[i]); err != nil {
				return
			}
			ctx.JSON(http.StatusOK, t.tweets[i])
		}
	}
}
