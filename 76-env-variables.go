/**


 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("FOO", "1") //set environment

	fmt.Println("FOO:", os.Getenv("FOO")) //get an env from machine

	fmt.Println("BAR:", os.Getenv("BAR")) //get non-existent env

	// get all envs in a machine
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
