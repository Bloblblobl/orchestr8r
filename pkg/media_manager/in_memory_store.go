package media_manager

import om "github.com/Bloblblobl/orchestr8r/pkg/object_model"

type inMemoryStore struct {
	tracks  map[int]om.Track
	albums  map[int]om.Album
	artists map[int]om.Artist
}

func newInMemoryStore() (*inMemoryStore, error) {
	return &inMemoryStore{
		tracks: map[int]om.Track{},
		albums: map[int]om.Album{},
		artists: map[int]om.Artist{},
	}, nil
}

func (ms *inMemoryStore) AddTrack(track om.Track) (err error) {
	return
}

func (ms *inMemoryStore) LookupTrack(id int) (track om.Track, err error) {
	return
}

func (ms *inMemoryStore) GetTracks(startToken int) (tracks []om.Track, err error) {
	return
}
