package main

import (
	"log"
	"os"

	"github.com/hieuphq/califit-be/goa/app"

	"github.com/hieuphq/califit-be/ext/controller"

	"github.com/goadesign/goa"
	"github.com/joho/godotenv"
)

func main() {
	// Create service
	service := goa.New("LLRA Backend API")
	if os.Getenv("ENV") == "DEV" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// // Setup external service
	// pgRoutingService := pgrouting.NewService()
	// mapboxService := mapbox.NewService(os.Getenv("MAPBOX_API_KEY"))

	// // Setup middlewares
	// jwtMiddleware, err := jwt.NewJWTMiddleware(app.NewJWTSecurity, os.Getenv("JWT_PUBLIC_KEY"))
	// if err != nil {
	// 	log.Fatal("failed to make jwtMiddleware by error: ", err)
	// }

	// // create sendgrid client
	// sendgrider := sendgrid.NewClient(os.Getenv("SENDGRID_API_KEY"))
	// // Mount middleware
	// app.UseJWTMiddleware(service, jwtMiddleware)
	// service.Use(middleware.RequestID())
	// service.Use(middleware.LogRequest(true))
	// service.Use(middleware.ErrorHandler(service, true))
	// service.Use(middleware.Recover())

	// // Setup DB
	// db, closeDB := config.NewPGConnection(os.Getenv("DB_DATASOURCE"))
	// defer closeDB()

	// // Setup firebase service
	// fbclient := firebase.NewService(
	// 	os.Getenv("FIREBASE_API_KEY"),
	// 	os.Getenv("FIREBASE_DOMAIN_PREFIX"),
	// 	os.Getenv("ANDROID_PACKAGE_NAME"),
	// 	os.Getenv("IOS_BUNDLE_ID"),
	// 	os.Getenv("WEB_ENDPOINT"),
	// 	os.Getenv("IOS_APPSTORE_ID"),
	// )

	// // Setup store
	// store := &repository.Repository{
	// 	User:    user.NewPGStore(db),
	// 	Vehicle: vehicle.NewPGStore(db),
	// 	Route:   route.NewPGStore(db),
	// 	Driver:  driver.NewPGStore(db),
	// }

	// // Mount "authentication" controller
	c := controller.NewAuthenticationController(service)
	app.MountAuthenticationController(service, c)

	// // Mount "swagger" controller
	// c2 := controller.NewSwaggerController(service)
	// app.MountSwaggerController(service, c2)

	// c3 := controller.NewRoutesController(service, pgRoutingService, mapboxService, store, fbclient)
	// app.MountRoutesController(service, c3)

	// c4 := controller.NewVehiclesController(service, store)
	// app.MountVehiclesController(service, c4)

	// c5 := controller.NewDriversController(service, store)
	// app.MountDriversController(service, c5)

	// c6 := controller.NewUserController(service, store, sendgrider)
	// app.MountUsersController(service, c6)

	// c7 := controller.NewFilesController(service)
	// controller.MountFilesController(service, c7)

	// Start service
	if err := service.ListenAndServe(":9000"); err != nil {
		service.LogError("startup", "err", err)
	}

}
