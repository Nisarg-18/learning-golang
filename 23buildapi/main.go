package main

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
	return c.CourseId != "" && c.CourseName != ""
}

func main()  {
	
}