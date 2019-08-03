package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"errors"
	//"other/otherpackage"
)

func main1() {
	http.HandleFunc("/dipak", sampleFunc)
	http.ListenAndServe(":2000", nil)

	//http.Handle("/handle", &myHandler{greeting: "Hello"})

}

func sampleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!"))
}

type myHandler struct {
	greeting string
}

//type Inc Incident

type Incident struct {
	IncientID   string `json:"incidentId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {

	//otherpackage.HelloFromOtherPackage()
	x := -1
	y := 15

	result, err := add(x, y)
	if err != nil {
		fmt.Println("do something...!!!", err)
	}
	fmt.Println(result)

	router := mux.NewRouter()

	router.HandleFunc("/incident", CreateIncident).Methods("POST")

	router.HandleFunc("/handle", GetGreeting).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func GetGreeting(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get greeting is executed")
	response := []byte("ibvckjsdbvgsduygcbkjsd kljvsdjlvblkjbsdlkjv")
	fmt.Println(response)
	w.Write(response)
}

func CreateIncident(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create incident handler")
	i := Incident{}
	fmt.Println("r.body : ", r.Body)
	fmt.Println("i before decoding : ", i)
	err := json.NewDecoder(r.Body).Decode(&i)
	fmt.Println("i after decoding : ", i)
	i.IncientID = "1234"

	err = errors.New("Some error occured")

	fmt.Println(i.IncientID, err)
	res, _ := json.Marshal(i)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(201)
		w.Write(res)
	}
}

func add(a int, b int) (int, error) {

	if a < 0 {
		return 0, errors.New("a is negative")
	} else if b < 0 {
		return 0, errors.New("b is negative")
	} else {
		return a + b, nil
	}
}
