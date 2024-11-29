package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`    // Changed from []string to string
	ConcertDates string   `json:"concertDates"` // Changed from []string to string
	Relations    string   `json:"relations"`
}

func (a *Artist) IsValid() bool {
	return a.ID > 0 && a.Name != ""
}
