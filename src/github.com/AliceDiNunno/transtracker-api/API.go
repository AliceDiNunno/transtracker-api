package transtracker_api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func StartApi() {
	r := gin.Default()

	println("Starting TransTracker API")
	db, err := gorm.Open("sqlite3", "/var/transtracker/objects.db")
	_ = db
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	r.Run(":30100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}