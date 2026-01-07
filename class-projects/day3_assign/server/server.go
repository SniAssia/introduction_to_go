package main

// the protocol (http) is a set of standard instructions to ensure
import (
	"encoding/json"
	"log"
	"net/http"
)

type student struct {
	Name string      `json:"name"`
	Age int      `json:"age"`
	Address string      `json:"address"`
	Courses []string `json:"courses"`
}

var students []student

func getUsers(w http.ResponseWriter,_*http.Request){
	
	
	jsonresponse,err := json.Marshal(students)
	if err != nil {
		log.Println(err)
	}
	w.Write(jsonresponse)


}
func postusers(w http.ResponseWriter,e *http.Request){
	//e.Body
	if e.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newuser student
	err := json.NewDecoder(e.Body).Decode(&newuser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	students = append(students, newuser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(newuser) 
}
func handler(w http.ResponseWriter,e *http.Request){
	switch e.Method {
	case http.MethodGet :
		getUsers(w,e)
	case http.MethodPost : 
		postusers(w,e)
	default : 
		w.WriteHeader(http.StatusNotImplemented)
}
}
func main() {
	courses := []string {"go","nahii"}
	s1 := student{"jane",20,"marjane",courses}
	courses1 := []string {"web","dev"}
	s2 := student{"jane2",21,"marjane",courses1}
	courses2 := []string {"course3","course4"}
	s3 := student{"jane3",22,"marjane",courses2}
	students = append(students,s1)
	students = append(students,s2)
	students = append(students,s3)
	http.HandleFunc("/students",handler)
	http.ListenAndServe(":8080",nil)
}
