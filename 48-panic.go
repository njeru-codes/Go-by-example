/**
panic means “Something went really wrong — stop everything!”

    It immediately crashes your program.

    It’s used when you can’t or shouldn’t recover from an error.

*/

package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
