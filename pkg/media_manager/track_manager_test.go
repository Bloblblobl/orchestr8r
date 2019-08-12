package media_manager


import (
	om "github.com/Bloblblobl/orchestr8r/pkg/object_model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("track manager tests", func() {
	var store *inMemoryStore
	var err error
	var trackManager om.TrackManager

	BeforeEach(func() {
		store = newInMemoryStore()
		trackManager, err = NewTrackManager(store)
		Ω(err).Should(BeNil())
		Ω(trackManager).ShouldNot(BeNil())
	})

	It("should add a track", func() {
		testId := 2
		testName := "test_track"
		testAlbumId := 3
		testArtistIds := []int{5, 8}

		t := om.Track{
			MediaObject: om.MediaObject{
				Id:   testId,
				Name: testName,
			},
			AlbumId:     testAlbumId,
			ArtistIds:   testArtistIds,
		}

		err = trackManager.AddTrack(t)
		Ω(err).Should(BeNil())
		Ω(store.tracks).Should(HaveLen(1))

		ts, nt, err := store.GetTracks(0)
		Ω(ts[0].MediaObject.Id).Should(Equal(testId))
		Ω(ts[0].MediaObject.Name).Should(Equal(testName))
		Ω(ts[0].AlbumId).Should(Equal(testAlbumId))
		Ω(ts[0].ArtistIds).Should(Equal(testArtistIds))
		Ω(nt).Should(Equal(-1))
		Ω(err).Should(BeNil())
	})
	
	It("should add multiple tracks", func() {

	})

	It("should lookup a track", func() {

	})

	It("should get all tracks", func() {

	})
})