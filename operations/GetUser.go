package operations
import(
	"fmt"
	"database/sql"
	"encoding/json"
	"net/http"
	_"github.com/lib/pq"
)
type UserId struct{
	ID int `json:"id"`
}
// var db *sql.DB
func GetUser(db ,res http.ResponseWriter,req,*http.Request){
	if req.Method!=http.MethodPost{
		http.Error(res,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	body,err:=ioutill.ReadAll(req.Body)
	if err!=nil{
		http.Error(res,"Error reading the request body",http.StatusBadRequest)
	}
	
	err=json.Unmarshal(body,&userId)
	if err!=nil{
		http.Error(res,"Error unmarshalling json",http.StatusBadRequest)
        return
	}
	var user User
	row:=db.QueryRow("Select id,name from users where id=$1",userId.ID)
	err=Scan(&user.ID,&user.Name)
	if err!=nil{
		if err==sql.ErrNoRows{
			http.Error(res,"Unable to find the data for given id",http.StatusNotFound)
			return
		}
		else{
			http.Error(res,"Internal Server Error",http.StatusInternalServerError)
			fmt.Println("Error fetching the data",err)
			return
		}
	}
	res.Header().Set("Content-Type":"application/json")
	jsonResponse,err=json.Marshal(user)
	if err!=nil{
		http.Error(res,"Internal Server Error",http.StatusInternalServerError)
		return
	}
	res.Write(jsonResponse)
}