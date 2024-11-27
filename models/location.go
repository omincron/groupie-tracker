package models

type Location struct {
	ID        int
	Locations []string
	Dates     map[string][]string
}
