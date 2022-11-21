package controllers

import (
	"github.com/gin-contrib/sessions"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	globals "proj/web/globals"
	// helpers "proj/web/helpers"
)

func FollowersGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(globals.Userkey)
		log.Println("User from followers get handler: ", user)
		// if user != nil {
		// 	log.Println("Inside if")
		// 	c.Redirect(http.StatusAccepted, "/dashboard")
		// 	return
		// }
		
		username := c.Request.URL.Path
		// [len("/followers/"):]
		log.Println("Username: ", username)
		userFollowers := globals.Followers[username]
		log.Println("Followers: ", userFollowers)
		// userFollowers := []string{} 
		// for key, element := range globals.followers {
		// 	// fmt.Println("Key:", key, "=>", "Element:", element)
		// 	if ( key==username ) {
		// 		userFollowers = append(userFollowers, element)
		// 	}
		// }

		c.HTML(http.StatusOK, "followers.html", gin.H{
			"content": "harsha",
			"user":    user,
		})
	}
} 

