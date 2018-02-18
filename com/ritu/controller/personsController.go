package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/rituK/com/ritu/modal"

	"github.com/rituK/com/ritu/utils"
)

var tmp []modal.Person

func GetUsers(w http.ResponseWriter, req *http.Request) {
	utils.LogTracingPrint("GetUsers --- starts")

	initialize()

	for p := range tmp {
		fmt.Println(tmp[p])
	}
	// Serialize the struct to JSON
	jsonBytes, err := json.Marshal(tmp)
	if err != nil {
		fmt.Println("error:", err)
	}

	SetHttpResponseHeader(w)
	w.Write(jsonBytes)

	utils.LogTracingPrint("GetUsers --- ends")
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {

	// Decode the json request
	var test modal.Person
	json.NewDecoder(req.Body).Decode(&test)
	utils.LogTracingPrint("updateUser --- starts")

	// find the person to update
	temp1, _err := findPerson(strconv.Itoa(test.ID))
	// clone the data set
	temp1.Address = test.Address
	temp1.Fname = test.Fname
	temp1.Lname = test.Lname
	temp1.ID = test.ID

	// index find from slice
	index, _errorr := findPersonIndexfromSlice(strconv.Itoa(test.ID))
	tmp[index] = *temp1

	if _errorr != nil || _err != nil {
		panic(_errorr)
	}
	GetUsers(w, req)

	utils.LogTracingPrint("updateUser --- ends")
}

func InsertUser(w http.ResponseWriter, req *http.Request) {

	utils.LogTracingPrint("InsertUser --- starts")

	var newUser modal.Person
	json.NewDecoder(req.Body).Decode(&newUser)
	newUser.ID = rand.Intn(10000)

	tmp = append(tmp, newUser)
	GetUsers(w, req)
	utils.LogTracingPrint("Insert User --- ends")
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {

	utils.LogTracingPrint("DeleteUser --- starts")
	params := mux.Vars(req)
	id := params["id"]
	index, _error := findPersonIndexfromSlice(id)

	if _error != nil {
		panic(_error)
	}
	tmp = append(tmp[:index], tmp[index+1:]...)
	GetUsers(w, req)
	utils.LogTracingPrint("DeleteUser --- ends")
}

func GetUser(w http.ResponseWriter, req *http.Request) {

	utils.LogTracingPrint("GetUser---- starts")
	params := mux.Vars(req)
	id := params["id"]
	temp1, _err := findPerson(id)

	// Serialize the struct to JSON
	jsonBytes, err := json.Marshal(temp1)
	if err != nil || _err != nil {
		fmt.Println("error:", err)
	}

	SetHttpResponseHeader(w)
	w.Write(jsonBytes)
	utils.LogTracingPrint("getUser......ends")
}

func findPersonIndexfromSlice(id string) (int, error) {

	idTemp, _ := strconv.Atoi(id)
	for p := range tmp {
		var temp1 = tmp[p]
		if temp1.ID == idTemp {
			return p, nil
		}
	}
	return len(tmp), nil
}

func findPerson(id string) (*modal.Person, error) {

	idTemp, _ := strconv.Atoi(id)
	for p := range tmp {
		var temp1 = tmp[p]
		if temp1.ID == idTemp {
			return &temp1, nil
		}
	}
	return nil, errors.New("Not Found")
}

func SetHttpResponseHeader(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

func initialize() {
	sizeOfArray := len(tmp)
	fmt.Println(sizeOfArray)

	if sizeOfArray == 0 {

		rangcount := 15
		tmp = make([]modal.Person, rangcount)

		for i := 0; i < rangcount; i++ {

			// Create an instance of our Person & Address struct

			temp2 := modal.Address{

				AddressLine1: "500 West Madison",
				City:         "Chicago",
				State:        "IL",
				ZipCode:      "60661",
			}
			temp1 := modal.Person{
				Fname:   "Adam" + strconv.Itoa(i),
				Lname:   "Smith" + strconv.Itoa(i),
				ID:      i,
				Address: &temp2,
			}

			// Assign this to array
			tmp[i] = temp1

		}
	}
}
