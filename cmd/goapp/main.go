package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lexunix/goapp/pkg/http"
	"github.com/lexunix/goapp/pkg/postgres"
)

func main() {
	db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ps := &postgres.PostService{DB: db}
	us := &postgres.UserService{DB: db}

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
	router.Use(sessions.Sessions("cook-my-sess", store))

	http.PostRoutes(router, &ph)
	http.UserRoutes(router, &uh)

	router.Run()
}
