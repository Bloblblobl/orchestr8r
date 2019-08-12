package media_manager

import (
	"errors"
	om "github.com/Bloblblobl/orchestr8r/pkg/object_model"
)

type inMemoryStore struct {
	tracks     map[int]om.Track
	albums     map[int]om.Album
	artists    map[int]om.Artist
}

func newInMemoryStore() (*inMemoryStore) {
	return &inMemoryStore{
		tracks: map[int]om.Track{},
		albums: map[int]om.Album{},
		artists: map[int]om.Artist{},
	}
}

func (ms *inMemoryStore) AddTrack(track om.Track) (err error) {
	ms.tracks[track.Id] = track
	return
}

func (ms *inMemoryStore) LookupTrack(id int) (track om.Track, err error) {
	track, ok := ms.tracks[id]
	if !ok {
		err = errors.New("not found")
	}
	return
}

func (ms *inMemoryStore) GetTracks(startToken int) (tracks []om.Track, nextToken int, err error) {
	for _, track := range ms.tracks {
		tracks = append(tracks, track)
	}
	nextToken = -1
	return
}
