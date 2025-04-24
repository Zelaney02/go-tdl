package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash Course"
var fullGolang = "Watch Nana's Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

// note := shorthand only works for local vars not global

func main() {
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil) // ignore if something happens
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to out Todolist App!"
	fmt.Fprintln(writer, greeting)
}
