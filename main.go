package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/masterraf21/trofonomie-backend/apis"
	repoMongo "github.com/masterraf21/trofonomie-backend/repositories/mongodb"
	"github.com/masterraf21/trofonomie-backend/usecases"

	"github.com/gorilla/mux"
	"github.com/masterraf21/trofonomie-backend/configs"
	"github.com/masterraf21/trofonomie-backend/utils/mongodb"
	"github.com/rs/cors"
)

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func main() {
	instance := mongodb.ConfigureMongo()
	serverReady := make(chan bool)
	server := Server{
		Instance:    instance,
		Port:        configs.Server.Port,
		ServerReady: serverReady,
	}
	server.Start()
}

// Start will start the server
func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8000"
	}

	r := new(mux.Router)

	counterRepo := repoMongo.NewCounterRepo(s.Instance)
	customerRepo := repoMongo.NewCustomerRepo(s.Instance, counterRepo)
	providerRepo := repoMongo.NewProviderRepo(s.Instance, counterRepo)
	menuRepo := repoMongo.NewMenuRepo(s.Instance, counterRepo)
	orderRepo := repoMongo.NewOrderRepo(s.Instance, counterRepo)

	customerUsecase := usecases.NewCustomerUsecase(customerRepo)
	providerUsecase := usecases.NewProviderUsecase(providerRepo)
	orderUsecase := usecases.NewOrderUsecase(orderRepo, customerRepo, providerRepo, menuRepo)
	menuUsecase := usecases.NewMenuUsecase(menuRepo, providerRepo)

	apis.NewCustomerAPI(r, customerUsecase)
	apis.NewMenuAPI(r, menuUsecase)
	apis.NewProviderAPI(r, providerUsecase)
	apis.NewOrderAPI(r, orderUsecase)

	handler := cors.Default().Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Printf("Starting server on port %s!", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Shutting Down Server...")
			log.Fatal(err.Error())
		}
	}()

	if s.ServerReady != nil {
		s.ServerReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
