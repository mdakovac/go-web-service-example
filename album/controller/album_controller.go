package album_controller

import (
	"album_domain/album_models"
	"fmt"
	"net/http"
	"router"
)

func SetupAlbumController(r router.Router, service album_models.AlbumServiceType) {
	albumRoutes := r.Group("/album")
	{
		albumRoutes.GET("/by-artist", func(c router.Context) {
			name := c.Param("name")
			albums, _ := service.FindByArtist(name)

			c.JSON(200, albums)
		})

		albumRoutes.POST("/", func(c router.Context) {
			var newAlbum album_models.Album

			fmt.Println(newAlbum)

			// Call BindJSON to bind the received JSON to
			// newAlbum.
			if err := c.BindJSON(&newAlbum); err != nil {
				c.JSON(http.StatusBadRequest, err.Error())
				return
			}

			id, _ := service.Add(newAlbum)
			c.JSON(200, id)
		})
	}
}
