package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Data types

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

// main
func main() {

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET") /*TBD*/
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE") /*TBD*/
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	// return the whole list of users
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)          // the ID of the person requested
	item := retrieve(params["id"]) // get it from Redis
	if item == "" {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(fmt.Sprintf("{\"error\":\"ID %s not found\"}", params["id"]))
	} else {
		json.NewEncoder(w).Encode(item) // write it back into the ResponseWriter
	}
	return
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	// create a new Person record
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	id := uuid.New().String()       // create our own UUID like a POST should
	mp, err := json.Marshal(person) // turn it into a slice of byte

	if err != nil {
		panic(err)
	}
	store(id, mp) // send it off to redis

	json.NewEncoder(w).Encode(fmt.Sprintf("{\"id\":\"%s\"}", id)) // write the ID back out
	return
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	numRemoved := remove(params["id"])
	if numRemoved < 1 {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(fmt.Sprintf("{\"error\":\"ID %s not found for delete\"}", params["id"]))
	} else {
		r.Write(w)
	}
}

func newRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.168:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func store(id string, p []byte) {
	client := newRedisClient()
	defer client.Close()

	err := client.Set(id, p, 0).Err() // {id, Person}
	if err != nil {
		panic(err)
	}
}

func retrieve(id string) string {
	client := newRedisClient()
	defer client.Close()

	val, err := client.Get(id).Result()

	if err != nil {
		return ""
	}
	return val
}

func remove(id string) int64 {
	client := newRedisClient()
	defer client.Close()

	success, err := client.Del(id).Result()
	if err != nil {
		panic(err)
	}
	return success
}
