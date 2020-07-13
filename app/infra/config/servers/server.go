package servers

import (
	"go-common/app/infra/config/routers"
	"go-common/library/net/http/umr"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
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
	write_timeout, err := strconv.Atoi(os.Getenv("HTTP_SERVER_WRITE_TIMEOUT"))
	if err != nil {
		write_timeout = 8
	}
	read_timeout, err := strconv.Atoi(os.Getenv("HTTP_SERVER_READ_TIMEOUT"))
	if err != nil {
		read_timeout = 5
	}

	httpServer := &http.Server{
		Addr:         addr,
		Handler:      server.UMRServer,
		WriteTimeout: time.Second * write_timeout,
		ReadTimeout:  time.Second * read_timeout,
	}

	log.Printf("Server started on %s", addr)
	err := httpServer.ListenAndServe()

	if err != nil {
		return errors.Wrap(err, "Listen ip and port error.")
	}

	return nil
}
