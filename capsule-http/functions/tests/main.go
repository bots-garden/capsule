// Package main to do some tests
package main

// Import the net/http, encoding/json and io/ioutil packages
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jaswdr/faker"
)

// Data '{"name":"Bob Morane","age":42}'
// Data is a struct that represents the JSON data to send and receive
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// PostJSON is a function that makes a POST request with JSON data to a given URL and returns the JSON response
func PostJSON(url string, data Data) (string, error) {
	// Encode the data struct as JSON and store it in a buffer
	buf, err := json.Marshal(data)
	// If there is an error, return the empty response and the error
	if err != nil {
		return "", err
	}
	// Make a POST request with the buffer as the body and set the content type to JSON
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(buf))
	// If there is an error, return the empty response and the error
	if err != nil {
		return "", err
	}
	// Close the response body when the function returns
	defer resp.Body.Close()
	// Read the response body as bytes
	body, err := ioutil.ReadAll(resp.Body)
	// If there is an error, return the empty response and the error
	if err != nil {
		return "", err
	}
	// Return the res struct and nil error
	return string(body), nil
}

// Main function
func main() {
	fake := faker.New()
	success := 0
	failure := 0

	failures := []string{}

	//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()
	//go func() {
	for i := 0; i <= 1000; i++ {
		// Create a Data struct with some values
		data := Data{
			Name: fake.Person().Name(),
			Age:  fake.Person().Faker.IntBetween(20, 80),
		}
		// Call the PostJSON function with a URL and the data struct and print the result or error
		res, err := PostJSON("http://localhost:8090", data)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// {"message":"Adrianna Kautzer 31","things":{"emoji":"🐯"}}
			if res == `{"message":"`+data.Name+" "+strconv.Itoa(data.Age)+`","things":{"emoji":"🐯"}}` {
				fmt.Println("🟢", res)
				success++
			} else {
				fmt.Println("🔴", res, "Expected:", data.Name+" "+strconv.Itoa(data.Age))
				failure++
				failures = append(failures, res+" / "+data.Name+" "+strconv.Itoa(data.Age))
			}
			//fmt.Println("Response:", res)
		}

	}
	fmt.Println("Nb Success:", success)
	fmt.Println("Nb Failure:", failure)
	//fmt.Println(failures)

	// display all failures
	for _, f := range failures {
		fmt.Println("-", f)
	}

	//}()
	//<-ctx.Done()
	//stop()

}
