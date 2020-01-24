package transtracker_api

import (
	"fmt"
	"github.com/AliceDiNunno/transtracker-api/V1"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func UnableToConnectToDatabaseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(500, gin.H{"error": gin.H{
			"code": 500,
			"detail": "Unable to connect to database" },
		})
		return
	}
}

func StartApi() {
	r := gin.Default()

	db, err := gorm.Open("sqlite3", "/var/transtracker/objects.db")
	_ = db
	if err != nil {
		fmt.Println("Unable to connect to database because of: " + err.Error())
		r.Use(UnableToConnectToDatabaseMiddleware())
	} else { return }

	println("Starting TransTracker API")

	v1 := r.Group("/v1")
	V1.ProductsRoute(v1)

	r.Run(":30100") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}