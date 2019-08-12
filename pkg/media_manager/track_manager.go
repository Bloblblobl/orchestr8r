package media_manager

import om "github.com/Bloblblobl/orchestr8r/pkg/object_model"

type trackManager struct {
	store om.TrackManager
}

func NewTrackManager(store om.TrackManager) (om.TrackManager, error) {
	return &trackManager{store: store}, nil
}

func (tm *trackManager) AddTrack(track om.Track) (err error) {
	return tm.store.AddTrack(track)
}

func (tm *trackManager) LookupTrack(id int) (track om.Track, err error) {
	return tm.store.LookupTrack(id)
}

func (tm *trackManager) GetTracks(startToken int) (tracks []om.Track, nextToken int, err error) {
	return tm.store.GetTracks(startToken)
}