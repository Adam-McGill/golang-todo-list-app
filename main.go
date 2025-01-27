package main

import (
	"fmt"
	"net/http"
)

var shortGolang = "Watch Go crash course"
var fullGolang = "Watch TWN Golang Full Course"
var rewardDessert = "Reward myself with a donut"
var taskItems = []string{shortGolang, fullGolang, rewardDessert}

func main() {

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)

	http.ListenAndServe(":8080", nil)

}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, task := range taskItems {
		fmt.Fprintln(writer, task)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello user. Welcome to our TodoList App!"
	fmt.Fprintln(writer, greeting)
}

//func printTasks(taskItems []string) {
//	fmt.Println("List of my Todos")
//	for index, task := range taskItems {
//		fmt.Printf("%d. %s\n", index+1, task)
//	}
//}

//func addTask(taskItems []string, newTask string) []string {
//	var updatedTaskItems = append(taskItems, newTask)
//	return updatedTaskItems
//}
