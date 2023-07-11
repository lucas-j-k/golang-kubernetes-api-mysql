package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/lucas-j-k/kube-go-microservices/user-service/user"
	"github.com/redis/go-redis/v9"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

func main() {

	// initialize env vars and SQL connection
	viper.AutomaticEnv()

	sqlPass := viper.GetString("MYSQL_PASSWORD")
	sqlDB := viper.GetString("MYSQL_DATABASE")
	sqlHost := viper.GetString("MYSQL_HOST")
	sqlPort := viper.GetString("MYSQL_PORT")
	sqlUser := viper.GetString("MYSQL_USER")
	redisPass := viper.GetString("REDIS_PASSWORD")
	redisHost := viper.GetString("REDIS_HOST")
	redisPort := viper.GetString("REDIS_PORT")

	fmt.Printf("VARS: %v|%v|%v|%v|%v|%v \n\n", sqlPass, sqlHost, sqlUser, redisPass, redisHost, redisPort)

	// set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true, // accept cookies
		Debug:            true, // should be development env only
	})

	// initialize Database connection and database services
	fmt.Println("Trying to connect to SQL")
	dataSourceName := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True", sqlUser, sqlPass, sqlHost, sqlPort, sqlDB)
	connection := sqlx.MustConnect("mysql", dataSourceName)
	fmt.Println("After sql connection")

	userService := user.UserService{
		Db: connection,
	}

	// Initialize Redis client and services
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", redisHost, redisPort), // TODO - env vars and concatenate
		Password: fmt.Sprintf("%v", redisPass),               // no password set
		DB:       0,                                          // use default DB
	})

	redisSessionManager := user.RedisSessionManager{
		Client: redisClient,
	}

	// setup Chi router and global middlewares
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(c.Handler)

	// initialise custom middleware. This is in an interface so we can inject our redis session manager
	cacheMiddleware := user.CacheMiddleware{SessionManager: &redisSessionManager}

	// protected test - delete this, replace with a user profile GET
	r.Route("/admin", func(r chi.Router) {
		r.Use(cacheMiddleware.SessionGuard)
		r.Get("/profile", user.GetUserProfile(&userService))
	})

	r.Post("/signup", user.Signup(&userService))
	r.Post("/login", user.Login(&userService, redisSessionManager))
	r.Post("/logout", user.Logout(&userService, redisSessionManager))

	// Healthcheck endpoint
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	fmt.Printf("Server running on port :8080\n\n")
	http.ListenAndServe(fmt.Sprintf(":%v", 8080), r)
}
