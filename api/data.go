package api

import "database/sql"
import "fmt"
import "log"

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
	rows, err := dm.DB.Query("SELECT id, sender_id, receiver_id, comment, created FROM like WHERE receiver_id=?", userID)
    if err != nil {
 		return nil, err
    }
	defer rows.Close()

    out := make([]*Like, 0)

    for rows.Next() {
    	like := new(Like)
        err := rows.Scan(&like.ID, &like.SenderID, &like.ReceiverID, &like.Comment, &like.Created)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%s, %s, %s, £%.2f\n", like.ID, like.SenderID, like.ReceiverID, like.Comment, like.Created)
        out = append(out, like)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
    for _, like := range out {
    	fmt.Printf("%s, %s, %s, %s", like.ID, like.SenderID, like.ReceiverID, like.Comment, like.Created)
  	}
  	
	return out, nil
}

