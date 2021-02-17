package main

import (
	"fmt"
	"time"
)

func main() {

	ts := time.Now().UTC().Format(time.RFC3339)

	fmt.Println(ts)

}
