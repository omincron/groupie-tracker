package services

import (
	"encoding/json"
	"groopie_local/models"
	"net/http"
	"sync"
)

var (
	cache        []models.ArtistFull
	cacheLock    sync.Mutex
	cacheUpdated bool
)

// FetchArtists fetches artist data from the external API
func FetchArtists() ([]models.Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []models.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// FetchEvents fetches events data from the external API
func FetchLocations() ([]models.Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []models.Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func FetchDate() ([]models.Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dates []models.Date
	if err := json.NewDecoder(resp.Body).Decode(&dates); err != nil {
		return nil, err
	}

	return dates, nil
}

// FetchRelations fetches relations data from the external API
func FetchRelations() ([]models.Relations, error) { // Changed from RelationData to Relations
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relations struct {
		Index []models.Relations `json:"index"` // Changed from RelationData to Relations
	}

	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations.Index, nil
}

// MergeData fetches and merges artist and relation data
func MergeData() ([]models.ArtistFull, error) {
	var artists []models.Artist
	var locations struct {
		Index []models.Location `json:"index"`
	}
	var relations struct {
		Index []models.Relations `json:"index"` // Changed from RelationData to Relations
	}

	// Fetch artists
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, err
	}

	// Fetch locations
	locResp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer locResp.Body.Close()
	if err := json.NewDecoder(locResp.Body).Decode(&locations); err != nil {
		return nil, err
	}

	// Fetch relations
	relResp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer relResp.Body.Close()
	if err := json.NewDecoder(relResp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	// Create maps for quick lookup
	locationsMap := make(map[int]models.Location)
	for _, loc := range locations.Index {
		locationsMap[loc.ID] = loc
	}

	relationsMap := make(map[int]models.Relations) // Changed from RelationData to Relations
	for _, rel := range relations.Index {
		relationsMap[rel.ID] = rel
	}

	// Create full artist data
	var artistsFull []models.ArtistFull
	for _, artist := range artists {
		artistFull := models.ArtistFull{
			Artist:    artist,
			Location:  locationsMap[artist.ID],
			Relations: relationsMap[artist.ID],
		}
		artistsFull = append(artistsFull, artistFull)
	}

	return artistsFull, nil
}

// GetCachedData returns cached data or refreshes it
func GetCachedData() ([]models.ArtistFull, error) {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	if !cacheUpdated {
		var err error
		cache, err = MergeData()
		if err != nil {
			return nil, err
		}
		cacheUpdated = true
	}

	return cache, nil
}
