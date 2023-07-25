package main

import (
	"base-go/adapter/repositories"
	blogs_repo "base-go/adapter/repositories/blogs-repo"
	cats_repo "base-go/adapter/repositories/cats-repo"
	events_repo "base-go/adapter/repositories/events-repo"
	users_repo "base-go/adapter/repositories/users-repo"
	vouchers_repo "base-go/adapter/repositories/vouchers-repo"
	"base-go/application"
	"base-go/application/auth"
	"base-go/application/blogs"
	"base-go/application/cats"
	"base-go/application/events"
	"base-go/application/users"
	"base-go/application/vouchers"
	"base-go/common/config"
	"base-go/common/logger"
	gw_http "base-go/gateway/http"
	"base-go/migrations"
	blogs_service "base-go/services/blogs"
	cats_service "base-go/services/cats"
	events_service "base-go/services/events"
	users_service "base-go/services/users"
	vouchers_service "base-go/services/vouchers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "base-go/docs"
)

// @title Grand-Demo-API
// @description This api for demo app
// @version 1.0
// @host localhost:8001
// @BasePath
func main() {
	mainEcho()
}

func mainEcho() {
	logger.Init()
	cnf := config.Get()

	logger.Info("Initializing...")
	gormdb := repositories.NewGormdb(cnf)
	migrations.Migrate(gormdb)

	logger.Info("Constructing dependencies...")
	catRepo := cats_repo.NewCatsRepo(gormdb)
	catsService := cats_service.NewCatsService(catRepo)
	catsInteractor := cats.NewCatsInteractor(catsService)
	userRepo := users_repo.NewUsersRepo(gormdb)
	usersService := users_service.NewUserService(userRepo)
	authInteractor := auth.NewAuthInteractor(usersService)
	usersInteractor := users.NewUsersInteractor(usersService)
	blogRepo := blogs_repo.NewBlogsRepo(gormdb)
	blogsService := blogs_service.NewBlogService(blogRepo)
	blogsInteractor := blogs.NewBlogsInteractor(blogsService)
	eventRepo := events_repo.NewEventsRepo(gormdb)
	eventsService := events_service.NewEventService(eventRepo)
	eventsInteractor := events.NewEventsInteractor(eventsService)
	voucherRepo := vouchers_repo.NewVouchersRepo(gormdb)
	vouchersService := vouchers_service.NewVoucherService(voucherRepo)
	vouchersInteractor := vouchers.NewVouchersInteractor(vouchersService)
	app := application.NewApp(authInteractor, catsInteractor, usersInteractor, blogsInteractor, eventsInteractor, vouchersInteractor)

	logger.Info("Constructing http server...")
	router := gw_http.EchoRouter(cnf, app)
	server := gw_http.NewHttpServer(cnf, router)
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	logger.Info("Setup OS signal handler...")
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Error("graceful shutdown timed out.. forcing exit, error: %s", shutdownCtx.Err().Error())
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	logger.Info("Starting http server at %s:%d", cnf.HttpConfig.Host, cnf.HttpConfig.Port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	logger.Info("Waiting for http server to exit...")
	<-serverCtx.Done()
	logger.Info("Goodbye.")
}
