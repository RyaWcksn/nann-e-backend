package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nann-e-backend/api/handler"
	"github.com/nann-e-backend/api/usecase"
	"github.com/nann-e-backend/config"
	database "github.com/nann-e-backend/pkgs/db"
	"github.com/nann-e-backend/server/middleware"
	ai "github.com/nann-e-backend/store/AI"
)

type Server struct {
	cfg     *config.Config
	usecase usecase.IUsecase
	handler handler.IHandler
}

var addr string
var SVR *Server
var db *sql.DB
var signalChan chan (os.Signal) = make(chan os.Signal, 1)

func (s *Server) initServer() {
	addr = ":9000"
	cfg := s.cfg
	if len(cfg.Server.HTTPAddress) > 0 {
		if _, err := strconv.Atoi(cfg.Server.HTTPAddress); err == nil {
			addr = fmt.Sprintf(":%v", cfg.Server.HTTPAddress)
		} else {
			addr = cfg.Server.HTTPAddress
		}
	}
}

func (s *Server) Register() {

	s.initServer()

	// MYSQL
	dbConn := database.NewDatabaseConnection(*s.cfg)
	if dbConn == nil {
		log.Fatal("Expecting DB connection but received nil")
	}

	db = dbConn.DBConnect()
	if db == nil {
		log.Fatal("Expecting DB connection but received nil")
	}

	ai := ai.NewAi(db)

	// Register service
	s.usecase = usecase.NewUsecase(ai)

	// Register handler
	s.handler = handler.NewHandler(s.usecase)
}

func NewService(cfg *config.Config) *Server {
	if SVR != nil {
		return SVR
	}
	SVR = &Server{
		cfg: cfg,
	}

	SVR.Register()

	return SVR
}

func (s Server) Start() {

	http.Handle("/api/register", middleware.ErrHandler(s.handler.Register))

	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatalf("error listening to address %v, err=%v", addr, err)
		}
	}()

	sig := <-signalChan
	log.Fatalf("%s signal caught", sig)

	// Doing cleanup if received signal from Operating System.
	err := db.Close()
	if err != nil {
		log.Fatalf("Error in closing DB connection. Err : %+v", err.Error())
	}
}
