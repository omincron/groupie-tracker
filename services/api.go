package services

import (
	"encoding/json"
	"groopie_local/models"
	"net/http"
	"sync"
	"time"
)

var (
	cache         []models.ArtistFull
	cacheLock     sync.Mutex
	cacheUpdated  bool
	lastCacheTime time.Time
)

func fetchFromAPI(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func FetchArtists() ([]models.Artist, error) {
	var artists []models.Artist
	err := fetchFromAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists, err
}

func FetchLocations() ([]models.Location, error) {
	var response struct {
		Index []models.Location `json:"index"`
	}
	err := fetchFromAPI("https://groupietrackers.herokuapp.com/api/locations", &response)
	return response.Index, err
}

func FetchRelations() ([]models.Relations, error) {
	var response struct {
		Index []models.Relations `json:"index"`
	}
	err := fetchFromAPI("https://groupietrackers.herokuapp.com/api/relation", &response)
	return response.Index, err
}

func MergeData() ([]models.ArtistFull, error) {
	var artists []models.Artist
	var locations []models.Location
	var relations []models.Relations
	var err error

	var wg sync.WaitGroup
	errChan := make(chan error, 3)

	wg.Add(3)
	go func() {
		defer wg.Done()
		artists, err = FetchArtists()
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		locations, err = FetchLocations()
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		defer wg.Done()
		relations, err = FetchRelations()
		if err != nil {
			errChan <- err
		}
	}()
	wg.Wait()
	close(errChan)

	for e := range errChan {
		return nil, e
	}

	locationMap := make(map[int]models.Location)
	for _, loc := range locations {
		locationMap[loc.ID] = loc
	}

	relationsMap := make(map[int]models.Relations)
	for _, rel := range relations {
		relationsMap[rel.ID] = rel
	}

	var artistsFull []models.ArtistFull
	for _, artist := range artists {
		artistsFull = append(artistsFull, models.ArtistFull{
			Artist:    artist,
			Location:  locationMap[artist.ID],
			Relations: relationsMap[artist.ID],
		})
	}

	return artistsFull, nil
}

func GetCachedData() ([]models.ArtistFull, error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if !cacheUpdated || time.Since(lastCacheTime) > 5*time.Minute {
		var err error
		cache, err = MergeData()
		if err != nil {
			return nil, err
		}
		cacheUpdated = true
		lastCacheTime = time.Now()
	}
	return cache, nil
}
