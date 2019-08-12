package object_model

type TrackManager interface {
	AddTrack(track Track) error
	LookupTrack(id int) (Track, error)
	GetTracks(startToken int) ([]Track, int, error)
}

type AlbumManager interface {
	AddAlbum(album Album) error
	LookupAlbum(id int) (Album, error)
	GetAlbums(startToken int) ([]Album, int, error)
}

type ArtistManager interface {
	AddArtist(artist Artist) error
	LookupArtist(id int) (Artist, error)
	GetArtists(startToken int) ([]Artist, int, error)
}