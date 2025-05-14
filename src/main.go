// package main

// import (
// 	"fmt"
// 	"src/test"
// )

// func main() {
// 	subFunc()
// 	test.MyFunction()

// 	var name string = "John Doe"
// 	fmt.Println(name)
// }
// func subFunc() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Hello, World2!")
// 	fmt.Println("Hello, World3!")
// }

// main tells go that is the main function and an executable file not a library
package main

import (
	"fmt"
	"net/http"
)
//net/http --> package that helps build HTTP servers and clients

func main() {

	// the line below tells the Go HTTP server to listen for HTTP requests on the /create-payment path, and when the request arrives, call the handleCreatePayment function
	http.HandleFunc("/create-payment", handleCreatePayment)
	http.ListenAndServe("localhost:8080", nil)
}

// function that runs in case a request comes through the create-payment route
// w is used to send a response back to the client while r contains info about the incoming HTTP req
func handleCreatePayment(w http.ResponseWriter, r *http.Request) {

}
