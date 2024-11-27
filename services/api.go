package services

import (
	"encoding/json"
	"fmt"
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

// FetchRelations fetches relations data from the external API
func FetchRelations() ([]models.RelationData, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var relations struct {
		Index []models.RelationData `json:"index"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&relations); err != nil {
		return nil, err
	}

	return relations.Index, nil
}

// MergeData fetches and merges artist and relation data
func MergeData() ([]models.ArtistFull, error) {
	// Use concurrency to fetch artists and relations
	var artists []models.Artist
	var relations []models.RelationData
	var errArtists, errRelations error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		artists, errArtists = FetchArtists()
	}()

	go func() {
		defer wg.Done()
		relations, errRelations = FetchRelations()
	}()

	wg.Wait()

	if errArtists != nil || errRelations != nil {
		return nil, fmt.Errorf("error fetching data: %v, %v", errArtists, errRelations)
	}

	// Map relations by artist ID for quick lookup
	relationsMap := make(map[int]models.RelationData)
	for _, relation := range relations {
		relationsMap[relation.ID] = relation
	}

	// Merge artist and relation data
	var artistsFull []models.ArtistFull
	for _, artist := range artists {
		artistFull := models.ArtistFull{
			Artist:   artist,
			Relation: relationsMap[artist.ID],
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
