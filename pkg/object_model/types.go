package object_model

type MediaObject struct {
	Id   int
	Name string
}

type Track struct {
	MediaObject
	AlbumId   int
	ArtistIds []int
}

type Album struct {
	MediaObject
	TrackIds  []int
	ArtistIds []int
}

type Artist struct {
	MediaObject
	AlbumIds []int
}