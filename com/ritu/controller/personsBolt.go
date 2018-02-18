package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rituK/com/ritu/modal"

	"github.com/rituK/com/ritu/dao"
	"github.com/rituK/com/ritu/utils"
)

var DBClient dao.IBoltDB

func GetBoltUsers(w http.ResponseWriter, req *http.Request) {
	utils.LogTracingPrint("GetUsers --- starts")
	var tmp []modal.Person
	var _err error
	tmp, _err = DBClient.QueryUsers(req.Context())
	// Serialize the struct to JSON
	jsonBytes, err := json.Marshal(tmp)
	if err != nil || _err != nil {
		fmt.Println("error:", err)
	}

	SetHttpResponseHeader(w)
	w.Write(jsonBytes)

	utils.LogTracingPrint("GetUsers --- ends")
}

func GetBoltUser(w http.ResponseWriter, req *http.Request) {

	utils.LogTracingPrint("GetBoltUser -- starts")
	params := mux.Vars(req)
	id := params["id"]

	person, _err := DBClient.QueryUser(req.Context(), id)

	jsonBytes, err := json.Marshal(person)

	if err != nil || _err != nil {
		panic(err)
	}
	SetHttpResponseHeader(w)
	w.Write(jsonBytes)
	utils.LogTracingPrint("GetBoltUser -- ends")
}

func SaveOrUpdateBoltUser(w http.ResponseWriter, req *http.Request) {

	utils.LogTracingPrint("save user --- starts")
	var person modal.Person
	json.NewDecoder(req.Body).Decode(&person)
	fmt.Println(person)

	jsonBytes, err := json.Marshal(person)

	_err := DBClient.SaveOrUpdateUser(req.Context(), jsonBytes, strconv.Itoa(person.ID))
	if _err != nil || err != nil {
		panic(_err)
	}
	utils.LogTracingPrint("save user --- ends")
}
