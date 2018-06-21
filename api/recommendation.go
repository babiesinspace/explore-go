package api

type Recommendation struct {
	UserID        int `json:"userId"`
	RecommendedID int `json:"recommendedId"`
}
