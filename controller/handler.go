package controller

import (
	"database/sql"
	"go-contacts-api/middleware"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/postgres"
	"github.com/gin-gonic/gin"
)

const SecretKey = "secret"

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())

	store, err := postgres.NewStore(h.db, []byte(SecretKey))
	if err != nil {
		log.Println(err)
	}

	router.Use(sessions.Sessions("mysession", store))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.POST("/logout", h.Logout)
	}

	api := router.Group("/api", middleware.Authentication())
	{
		api.POST("/contact", h.CreateContact)
		api.GET("/contacts", h.GetContacts)
		api.GET("/contact/:id", h.GetContactbyUser)
		api.PUT("/contact/:id", h.UpdateContact)

	}

	return router
}
