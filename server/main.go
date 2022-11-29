package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", serverHome).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":4000", r))
// }

// func serverHome(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("<h1>Hello from the other world</h1>"))
// }

// model for course
type Course struct {
	CourseId    string `json:"courseId"`
	CourseName  string
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake Db
var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
	// return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome")
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "React Js", CoursePrice: 299, Author: &Author{Fullname: "Rajat", Website: "test.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Angular Js", CoursePrice: 399, Author: &Author{Fullname: "Rajat Kumar", Website: "test2.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course", getAllCourse).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to code
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers
//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to my golang api</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all course")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(courses)

	//get id from requets
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given Id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Create course"))
	w.Header().Set("Content-Type", "application/json")

	//if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please sends some data")
	}

	//empty object
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please enter some data in json")
		return
	}
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Create course"))
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop value and find acc to id . delete it and yupdate it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Create course"))
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//loop value and find acc to id . delete it and yupdate it
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			// json.NewEncoder(w).Encode(course)
			break
		}
	}
}
