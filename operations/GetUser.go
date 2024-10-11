package operations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// type User struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

type UserRequest struct {
	ID int `json:"id"`
}

func GetUser(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse the JSON
	var userReq UserRequest
	err = json.Unmarshal(body, &userReq)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}

	// Query the database for the specific user
	row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", userReq.ID)

	var user User
	err = row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			fmt.Println("Error querying database:", err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}