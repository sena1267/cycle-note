package main

import (
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"github.com/sena1267/cycle-note/config"
	"github.com/sena1267/cycle-note/gen/protobuf/auth/v1/authv1connect"
	"github.com/sena1267/cycle-note/gen/protobuf/medicine/v1/medicinev1connect"
	"github.com/sena1267/cycle-note/gen/protobuf/ping/v1/pingv1connect"
	"github.com/sena1267/cycle-note/gen/protobuf/user/v1/userv1connect"
	"github.com/sena1267/cycle-note/handler"
	"github.com/sena1267/cycle-note/infrastructure"
	"github.com/sena1267/cycle-note/infrastructure/repository"
	"github.com/sena1267/cycle-note/middleware"
	"github.com/sena1267/cycle-note/pkg/auth"
	"github.com/sena1267/cycle-note/pkg/util/xidgen"
	"github.com/sena1267/cycle-note/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}
	dbClient, err := infrastructure.NewDBClient(cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}

	authenticator, err := auth.NewAuthenticator(cfg.Auth)
	if err != nil {
		log.Fatalln(err)
	}

	userRepo := repository.NewUserRepository(dbClient)
	authUsecase := usecase.NewAuthUsecase(userRepo, authenticator, xidgen.XIDGenerator{})
	authServiceHandler := handler.NewAuthHandler(authUsecase)

	userUsecase := usecase.NewUserUsecase(userRepo)
	userServiceHandler := handler.NewUserHandler(userUsecase)

	// medicine
	medicineRepo := repository.NewMedicineRepository(dbClient)
	medicineUsecase := usecase.NewMedicineUsecase(medicineRepo)
	medicineServiceHandler := handler.NewMedicineHandler(medicineUsecase)

	interceptors := connect.WithInterceptors(
		middleware.NewAuthInterceptor(authenticator),
		middleware.NewCurrentTimeInterceptor(),
	)

	var _ authv1connect.AuthServiceHandler = (*handler.AuthHandler)(nil)

	mux := http.NewServeMux()
	//authPath, authHandler := authv1connect.NewAuthServiceHandler(authServiceHandler)
	authPath, authHandler := authv1connect.NewAuthServiceHandler(authServiceHandler)
	mux.Handle(authPath, authHandler)
	fmt.Println(authPath)
	mux.Handle(userv1connect.NewUserServiceHandler(userServiceHandler, interceptors))
	mux.Handle(medicinev1connect.NewMedicineServiceHandler(medicineServiceHandler, interceptors))
	mux.Handle(pingv1connect.NewPingServiceHandler(&handler.PingHandler{}))
	err = http.ListenAndServe(
		"0.0.0.0:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatalln(err)
	}
}
