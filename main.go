package main

import (
	"net/http"

	_ "docs"
	"handlers"
	"storage"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

var (
	db     *storage.DBConnected = &storage.DBConnected{DB: new(gorm.DB)}
	conStr string               = "host=localhost user=postgres password=postgres dbname=test_db port=5432 sslmode=disable"
	router *gin.Engine          = gin.Default()
)

// @title           Studysyncr API
// @version         1.0
// @description     API for Studysyncr practice project
func main() {
	db.Init(conStr)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("notes/:user/:id", handlers.GetNote(db))
	router.GET("notes/:user", handlers.GetAllNotes(db))
	router.POST("notes/:user", handlers.PostNote(db))
	router.DELETE("notes/:user/:id", handlers.DeleteNote(db))
	router.PATCH("notes/:user/:id", handlers.PatchNote(db))
	router.Run("localhost:8080")
}
