package server

import (
	"context"
	"net/http"
	"temp-prj/internal/server/users"

	_ "temp-prj/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	srv *http.Server
}

func New(addr string, usersService users.UserService) *Server {
	srv := &http.Server{
		Addr: addr,
	}
	uh := users.New(usersService)
	r := configureRouter(uh)
	srv.Handler = r

	return &Server{
		srv: srv,
	}
}
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func configureRouter(uh *users.UsersHandler) *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	users := router.Group("/users")
	users.POST("/login", uh.Login)
	users.POST("/register", uh.Register)

	return router
}
