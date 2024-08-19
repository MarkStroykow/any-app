package handler

import (
	"github.com/MarkStroykow/any-app/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}

	api := router.Group("/api")
	{
		cart := api.Group("/cart")
		{
			cart.POST("/", h.createCart)
			cart.GET("/", h.getAllCarts)
			cart.GET("/:id", h.getCartById)
			cart.PUT("/:id", h.updateCart)
			cart.DELETE("/:id", h.deleteCart)

			items := cart.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
