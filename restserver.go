// personrest
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var VALID_PATH_REGEX *regexp.Regexp

//Nested structure
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

func init() {
	VALID_PATH_REGEX = regexp.MustCompile("/people/[0-9]+$")
}

type RestServer struct {
	cache PersonsCache
}

func NewRestServer(c *PersonsCache) *RestServer {
	return &RestServer{cache: *c}
}

func (server *RestServer) Listen(address string) error {
	//https ListenAndServe ALWAYS return error
	err := http.ListenAndServe(address, server)
	return err
}

func (server *RestServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	//TODO proper testing
	path := req.URL.Path
	if path == "/" {
		fmt.Fprintf(rw, "Welcome  to the RestSever! \n")
		fmt.Fprintf(rw, "Use /people/[0-9]+ to commence operations ! \n")
		rw.WriteHeader(http.StatusAccepted)
		return
	} else if match := VALID_PATH_REGEX.FindAllString(path, 1); len(match) == 0 {
		fmt.Printf("not found path is %s \n", path)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	match := VALID_PATH_REGEX.FindAllString(path, 1)
	fmt.Printf("found match is %s \n", match[0])
	sf := strings.SplitAfter(match[0], "/")
	lsf := sf[2]
	fmt.Printf("last field is %s \n", lsf)

	personId, err := strconv.ParseInt(lsf, 10, 64)
	if handlePossibleInternalServerError(&err, &rw) {
		return
	}

	method := req.Method
	if method == "GET" {

		fmt.Printf("Method GET \n")

		person, found := server.cache.GetPerson(personId)
		if found {
			rw.Header().Set("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(person)
			if handlePossibleInternalServerError(&err, &rw) {
				return
			}

		} else {
			http.Error(rw, "No entry with this Id", http.StatusNotFound)
		}
	} else if method == "POST" {

		body, err := ioutil.ReadAll(req.Body)
		if handlePossibleInternalServerError(&err, &rw) {
			return
		}

		person, err := createPersonStructure(body)
		if handlePossibleInternalServerError(&err, &rw) {
			return
		}

		server.cache.InsertPerson(personId, person)

	} else if method == "DELETE" {
		server.cache.DeletePerson(personId)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
}

func handlePossibleInternalServerError(err *error, rw *http.ResponseWriter) (errorHappened bool) {
	if *err != nil {
		http.Error(*rw, (*err).Error(), http.StatusInternalServerError)
		errorHappened = true
	}
	errorHappened = false
	return
}

func createPersonStructure(jsonBytes []byte) (p Person, err error) {
	err = json.Unmarshal(jsonBytes, &p)
	if err != nil {
		return p, err
	}
	return p, nil
}

func createPersonJson(person Person) (jsonBytes []byte, err error) {
	//err = json.Unmarshal(jsonBytes, &p)
	jsonBytes, err = json.Marshal(person)
	if err != nil {
		return
	}
	return
}
