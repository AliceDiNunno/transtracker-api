package transtracker_api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func StartApi() {
	r := gin.Default()

	db, err := gorm.Open("sqlite3", "test.db")
	_ = db
	if err != nil {
		panic("failed to connect database")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":30100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}