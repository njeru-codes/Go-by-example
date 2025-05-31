/**


 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("=== GET Request ===")
	resp, err := http.Get("https://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))

	fmt.Println("=== POST Request ===")
	data := `name=Go&message=Hello`

	// Send POST request
	response, error := http.Post("https://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(data),
	)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	// Read and print response body
	newbody, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(newbody))

	fmt.Println("=== DELETE Request ===")

}
