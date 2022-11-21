package controllers

import (
	"github.com/gin-contrib/sessions"

	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"

	"time"
	"log"
)

func TweetGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		// if user != nil {
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }

		
		c.HTML(http.StatusOK, "composeTweet.html", gin.H{
			"user":    user,
		})
	}
} 

func TweetPostHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		// if user != nil {
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }

		tweetMsg := c.PostForm("tweetMsg")
		
		tweet := globals.Tweet{
			Time:  time.Now(),    
			Text: tweetMsg ,     
		}

		globals.Tweets[user.(string)] = append(globals.Tweets[user.(string)], tweet)

		log.Println("Posted tweet: ", globals.Tweets[user.(string)])
		
		// c.HTML(http.StatusOK, "index.html", gin.H{
		// 	"user":    user,
		// })
		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	}
} 

