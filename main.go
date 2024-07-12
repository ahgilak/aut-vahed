package main

import (
	"flag"
	"strings"
	// "io"
	// "os"
)

func main() {
	var username, password, courses_str string
	var courses []string

	flag.StringVar(&username, "username", "", "username")
	flag.StringVar(&password, "password", "", "password")
	flag.StringVar(&courses_str, "courses", "", "a space seperated list of courses, course id is like 1021303_100__")
	flag.Parse()

  if username == "" || password == "" || courses_str == "" {
    flag.PrintDefaults()
    return
  }

	courses = strings.Split(courses_str, " ")

	ctx := NewContext()
	ctx.InitLogin()
  
	for {
		err := ctx.Login(username, password)
		if err == nil {
			break
		}
	}

	for _, c := range courses {
		for {
			err := ctx.AddCourse(c)
			if err == nil {
				break
			}
		}
	}
}
