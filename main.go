package main

import (
	"album_controller"
	"album_domain/album_repository"
	"album_domain/album_service"
	"db/db"
	"router"
	"util/env_vars"
)

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {
	conn := db.Connect(env_vars.EnvVariables.DATABASE_URL)
	defer db.Disconnect(conn)

	router := router.NewRouter()

	albumRepository := album_repository.NewAlbumRepository(conn)
	albumService := album_service.NewAlbumService(albumRepository)
	album_controller.SetupAlbumController(router, albumService)

	
    router.Run("localhost:8080")
}
