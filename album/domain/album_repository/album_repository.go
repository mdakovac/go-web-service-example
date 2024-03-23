package album_repository

import (
	"album_domain/album_models"
	"context"
	"db/db"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type AlbumRepository struct{
	conn db.DbConnectionType
}


func (ar *AlbumRepository) Add(a album_models.Album) (int64, error) {
	var insertedID int64 // Variable to store the inserted ID

    err := ar.conn.QueryRow(
		context.Background(), 
		"INSERT INTO album (title, artist, price) VALUES ($1, $2, $3) RETURNING id", 
		a.Title, 
		a.Artist, 
		a.Price,
	).Scan(&insertedID)

    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }
	
    return insertedID, nil
}

func (ar *AlbumRepository) FindByArtist(name  string) ([]album_models.Album, error){
	rows, err := ar.conn.Query(context.Background(), "SELECT * FROM album WHERE artist = $1", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist1 %q: %v", name, err)
    }
	
    defer rows.Close()

    albums, err := pgx.CollectRows(rows, pgx.RowToStructByName[album_models.Album])
	if err != nil {
        return nil, fmt.Errorf("albumsByArtist2 %q: %v", name, err)
    }

    return albums, nil
}

func NewAlbumRepository(conn db.DbConnectionType) *AlbumRepository{
	return &AlbumRepository{
		conn,
	}
}
