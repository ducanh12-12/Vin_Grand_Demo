package http

import (
	"base-go/application"
	"base-go/gateway/http/controllers"
	"fmt"
	"net/http"

	config "base-go/common/config"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewHttpServer(cnf *config.Config, engine http.Handler) *http.Server {
	//tlsConfig := &tls.Config{
	//	MinVersion:               tls.VersionTLS12,
	//	PreferServerCipherSuites: true,
	//}

	//tlsConfig.Certificates = make([]tls.Certificate, 1)
	//var err error
	//if tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(cnf.X509CertFile, cnf.X509KeyFile); err != nil {
	//	logger.Get().Err(err).Msg("Unable to load TLS certificates, aborting...")
	//	panic(err)
	//}

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cnf.HttpConfig.Host, cnf.HttpConfig.Port),
		Handler: engine,
		//TLSConfig: tlsConfig,
		//ReadTimeout:    readTimeout,
		//WriteTimeout:   writeTimeout,
		//MaxHeaderBytes: maxHeaderBytes,
	}
}

// @title Grand-Demo-API
// @description This api for demo app
// @version 1.0
// @host localhost:8001
// @BasePath
func EchoRouter(cnf *config.Config, app *application.App) http.Handler {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				// Log lỗi
				fmt.Println("Error:", err)
			}
			return err
		}
	})
	// middlewares config goes here...
	// subrouters/controllers mounting
	catsController := controllers.NewCatsController(app.Cats)
	catsController.Mount(e)
	usersController := controllers.NewUsersController(app.Users)
	usersController.Mount(e)
	authController := controllers.NewAuthController(app.Auth)
	authController.Mount(e)
	blogsController := controllers.NewBlogsController(app.Blogs)
	blogsController.Mount(e)
	eventsController := controllers.NewEventsController(app.Events)
	eventsController.Mount(e)
	vouchersController := controllers.NewVouchersController(app.Vouchers)
	vouchersController.Mount(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
