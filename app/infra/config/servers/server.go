package servers

import (
	"go-common/app/infra/config/routers"
	"go-common/library/net/http/umr"
	"log"
	"net/http"
)

type Server struct {
	UMRServer *umr.Engine
}

func NewServer() (*Server, error) {
	server := &Server{
		UMRServer: umr.Default(),
	}

	return server, nil
}

func (server *Server) Init() error {
	routers.NewRouter(server.UMRServer)
	return nil
}

func (server *Server) Start(addr string) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: server.UMRServer,
	}

	log.Printf("Server started on %s", addr)
	err := httpServer.ListenAndServe()

	if err != nil {
		return err
	}

	return nil
}
