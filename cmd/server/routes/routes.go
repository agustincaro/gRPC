package routes

import (
	"github.com/agustincaro/gRPC-client/cmd/server/handler"
	"github.com/agustincaro/gRPC-client/internal/buyer"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
}

func NewRouter(r *gin.Engine) Router {
	return &router{r: r}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildBuyerRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildBuyerRoutes() {
	repo := buyer.NewRepository()
	service := buyer.NewService(repo)
	handler := handler.NewBuyer(service)
	buyersGroup := r.rg.Group("/buyers")
	buyersGroup.GET("/", handler.GetAll())
}
