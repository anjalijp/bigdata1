package models

type Review struct {
	ListingID    string `json:"listing_id"`
	ID           string `json:"id"`
	Date         string `json:"date"`
	ReviewerID   string `json:"reviewer_id"`
	ReviewerName string `json:"reviewer_name"`
	Comments     string `json:"comments"`
}
