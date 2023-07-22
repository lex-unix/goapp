package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lexunix/goapp/pkg/database"
	"github.com/lexunix/goapp/pkg/http"
	_ "github.com/lib/pq"
)

func main() {
	conn := os.Getenv("DATABASE_URL")
	fmt.Println(conn)
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ps := &database.PostService{DB: db}
	us := &database.UserService{DB: db}

	var ph http.PostHandler
	var uh http.UserHandler
	ph.PostService = ps
	uh.UserService = us

	router := gin.Default()

	store, err := redis.NewStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte("secret"))
	if err != nil {
		log.Fatal(err)
	}
	redis.SetKeyPrefix(store, "mysession:")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true

	router.Use(sessions.Sessions("cook-my-sess", store))
	router.Use(cors.New(config))

	http.PostRoutes(router, &ph)
	http.UserRoutes(router, &uh)

	router.Run()
}
