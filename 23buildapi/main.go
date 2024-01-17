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

// models - diff file
type Courses struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake DB
var courses []Courses

// middleware, helpers - diff file
func (c *Courses) isEmpty() bool {
	// return c.CourseId != "" && c.CourseName != ""
	return c.CourseName == ""
}

func main()  {
	fmt.Println("Welcome to APIs")

	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Courses{CourseId: "2",CourseName: "Python", CoursePrice: 100, Author: &Author{Fullname: "ABC",Website: "ABC.com"}})
	courses = append(courses, Courses{CourseId: "3",CourseName: "GO", CoursePrice: 300, Author: &Author{Fullname: "ABC",Website: "ABC.com"}})
	courses = append(courses, Courses{CourseId: "4",CourseName: "Java", CoursePrice: 200, Author: &Author{Fullname: "ABC",Website: "ABC.com"}})
	courses = append(courses, Courses{CourseId: "5",CourseName: "C++", CoursePrice: 1000, Author: &Author{Fullname: "ABC",Website: "ABC.com"}})
	courses = append(courses, Courses{CourseId: "6",CourseName: "Flutter", CoursePrice: 1100})

	// routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	r.HandleFunc("/course",createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")
	r.HandleFunc("/deleteallcourses", deleteAllCourses).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))
}


// controllers - diff file

func serveHome (w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("getting all courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourseById (w http.ResponseWriter, r *http.Request){
	fmt.Println("get course by ID")
	w.Header().Set("Content-type","application/json")

	// get the id from request
	params := mux.Vars(r)

	for _, course := range courses{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the id:"+ params["id"])
}

func createCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("create a course")
	w.Header().Set("Content-type","application/json")

	body := r.Body
	if(body==nil){
		json.NewEncoder(w).Encode("Body is empty")
		return
	}

	var course Courses

	_ = json.NewDecoder(r.Body).Decode(&course)

	for _, c := range courses{
		if(c.CourseName == course.CourseName){
			json.NewEncoder(w).Encode("Name already exists")
			return
		}
	}

	if(course.isEmpty()){
		json.NewEncoder(w).Encode("JSON is empty")
		return
	}

	// generate id 
	rand.NewSource(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100))

	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateCourseById (w http.ResponseWriter, r *http.Request){
	fmt.Println("Update course by id")
	w.Header().Set("Content-type","application/json")

	params := mux.Vars(r)
	body := r.Body

	if body == nil {
		json.NewEncoder(w).Encode("Body is null")
		return
	}

	var course Courses

	_ = json.NewDecoder(r.Body).Decode(&course)

	if(course.isEmpty()){
		json.NewEncoder(w).Encode("JSON is empty")
		return
	}

	for index, c := range courses{
		if(c.CourseId == params["id"]){
			courses = append(courses[:index], courses[index+1:]...)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found") 
}

func deleteCourseById (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete course by id")

	w.Header().Set("Content-type","application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if(course.CourseId==params["id"]){
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted Successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with the id")
}

func deleteAllCourses (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")

	// for index, _ := range courses {
	// 	courses = append(courses[:index],courses[index+1:]...)
	// }

	courses = nil
	
	json.NewEncoder(w).Encode("All courses deleted")
}