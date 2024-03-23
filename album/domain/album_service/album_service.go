package album_service

import (
	"album_domain/album_models"
)

type AlbumService struct{
	albumRepository album_models.AlbumRepositoryType
}

func (as *AlbumService) Add(a album_models.Album) (int64, error) {
	return as.albumRepository.Add(a)
}

func (as *AlbumService) FindByArtist(name  string) ([]album_models.Album, error){
	return as.albumRepository.FindByArtist(name)
}

func NewAlbumService(albumRepository album_models.AlbumRepositoryType) *AlbumService {
	return &AlbumService{
		albumRepository,
	}
}