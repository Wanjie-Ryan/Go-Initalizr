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
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v74"
	"log"
	"net/http"
)

//net/http --> package that helps build HTTP servers and clients

func main() {

	// the line below tells the Go HTTP server to listen for HTTP requests on the /create-payment path, and when the request arrives, call the handleCreatePayment function
	http.HandleFunc("/create-payment", handleCreatePayment)
	http.HandleFunc("/health", healtCheck)
	log.Println("Listening on port 8080")
	var err error = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	// var names [] string = []string {"John", "Doe"}
	// fmt.Println(names)

	// type conversion in GO, converting string to byte

	// var responseString string = "Server is up and running"
	// var response [] byte = []byte(responseString)
	// fmt.Println(response)
}

// function that runs in case a request comes through the create-payment route
// w is used to send a response back to the client while r contains info about the incoming HTTP req
func healtCheck(w http.ResponseWriter, r *http.Request) {

	// instead of writing, you see here line 56 how the response variable is assigned a byte array
	// var response [] byte = []byte ("Server is up and running")

	// one can instead write
	response := []byte("Server is up and running")
	// the line below means that, w.write expects an int and error to be returned, BUT we don't really care about the int, only the error returned, therefore the _ is used to omit the int being returned, then the err is assigned to the error variable
	_, err := w.Write(response)
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	//fprintf writes to the http response w
	// fmt.Fprintf(w, "Hello, World!")
}

func handleCreatePayment(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// represents the data coming in from the request
	var req struct{
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Address1 string `json:"address_1"`
		Address2 string `json:"address_2"`
		City string `json:"city"`
		State string `json:"state"`
		Zip string `json:"zip"`
		Country string `json:"country"`

	}
	// once the request comes in, decode the request body, then assign it to the struct object created there on top
	// the & sign is a pointer to the struct
	err :=json.NewDecoder(r.Body).Decode(&req)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("request body error",err)
		log.Println(err)
		return
	}
	fmt.Println("incoming request",req)
	params := &stripe.PaymentIntentParams{
		
	}
}
