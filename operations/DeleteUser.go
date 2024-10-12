package operations
import(
	"fmt"
	"encoding/json"
	"io/ioutil"
	_"github.com/lib/pq"
	"database/sql"
	"net/http"
)
var db *sql.DB
func DeleteUser(res http.ResponseWriter,req *http.Request){
	if req.Method!=http.MethodDelete{
		http.Error(res,"Method Not Allowed",http.StatusMethodNotAllowed)
		return
	}
	body,err:=ioutil.ReadAll(req.Body)
	if err!=nil{
		http.Error(res,"Error parsing the request body",http.StatusBadRequest)
		return
	}
	var user User
	var u_id UserId
	err=json.Unmarshal(body,&u_id)
	query:="DELETE FROM users where id=$1"
	result,err:=db.Exec(query,&u_id.ID).Scan(&user.ID,&user.Name)
	if err!=nil{
		if err==sql.ErrNoRows{
			http.Error(res,"User not found",http.StatusNotFound)
			return
		}
		fmt.Println(err)
		http.Error(res,"Internal Server Error",http.StatusInternalServerError)
		return 
	}
	response,err:=json.Marshal(&user)
	if err!=nil{
		http.Error(res,"Internal Server Error",http.StatusInternalServerError)
		return
	}
	res.Header().Set("Content-Type","application/json")
	res.Write(response)
}