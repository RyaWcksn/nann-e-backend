package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/nann-e-backend/api/handler"
	"github.com/nann-e-backend/api/usecase"
	"github.com/nann-e-backend/config"
	database "github.com/nann-e-backend/pkgs/db"
	"github.com/nann-e-backend/server/middleware"
	ai "github.com/nann-e-backend/store/AI"
	gpt "github.com/nann-e-backend/store/GPT"
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
	fmt.Println(dbConn)
	if dbConn == nil {
		log.Fatal("Expecting DB connection but received nil")
	}

	db = dbConn.DBConnect()
	fmt.Println(db)
	if db == nil {
		log.Fatal("Expecting DB connection but received nil")
	}

	ai := ai.NewAi(db)
	gpt := gpt.NewGpt(s.cfg.App.API)

	// Register service
	s.usecase = usecase.NewUsecase(ai, gpt)

	// Register handler
	s.handler = handler.NewHandler(s.usecase, *s.cfg)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	http.Handle("/api/register", middleware.ErrHandler(s.handler.Register))
	http.Handle("/api/chat", middleware.ErrHandler(s.handler.Chat))
	http.Handle("/api/dashboard", middleware.ErrHandler(s.handler.GetData))
	http.Handle("/api/generate", middleware.ErrHandler(s.handler.GenerateUrl))

	srv := http.Server{Addr: addr}
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatalf("error listening to address %v, err=%v", addr, err)
		}
	}()

	select {
	case <-signalChan:
		log.Fatal("Signal received. Shutting down...")
	}
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	// Doing cleanup if received signal from Operating System.
	err := db.Close()
	if err != nil {
		log.Fatalf("Error in closing DB connection. Err : %+v", err.Error())
	}
}
