package album_models

type Album struct {
    ID     int64    `json:"id"`
    Title  string   `json:"title"`
    Artist string   `json:"artist"`
    Price  float32  `json:"price"`
}


type AlbumRepositoryType interface {
	Add(a Album) (int64, error) 
	FindByArtist(n string) ([]Album, error)
}


type AlbumServiceType interface {
	Add(a Album) (int64, error) 
	FindByArtist(n string) ([]Album, error)
}