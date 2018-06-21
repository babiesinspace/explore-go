package api

import "database/sql"

type DataManager struct {
	DB *sql.DB
}

// generateRecommendations is a dummy function to generate
// recommended matches for a user. It spits out random integers
// that represent the user IDs of people who will be shown to
// the requesting user. The outputted IDs are not guaranteed to be
// unique nor exclusive of the requesting user.
// You do not need to make any updates to this function.
func (dm *DataManager) generateRecommendations(userID int) ([]*Recommendation, error) {
	min := 100000
	max := 999999
	numRecs := 10
	out := make([]*Recommendation, numRecs)
	for i := 0; i < numRecs; i++ {
		var id int
		err := dm.DB.QueryRow(`SELECT ABS(RANDOM()) % (? - ?) + ?;`, max, min, min).Scan(&id)
		if err != nil {
			return nil, err
		}
		out[i] = &Recommendation{
			UserID:        userID,
			RecommendedID: id,
		}
	}
	return out, nil
}

func (dm *DataManager) likes(userID int) ([]*Like, error) {
	// TODO: Finish me!
	return nil, nil
}
