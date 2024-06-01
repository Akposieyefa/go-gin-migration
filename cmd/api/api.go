package api

import (
	"net/http"

	"github.com/akposiyefa/go-gin-migration/core/routers"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()

	routers.ApiRoutes(router)

	return http.ListenAndServe(s.addr, router)

}
