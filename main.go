package main

import (
	"fmt"
	"os"
	"strings"

	uuid "go.vxfiber.dev/vx-uuid"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "v4" {

		fmt.Println(strings.ToUpper(uuid.NewV4()))
	} else {
		fmt.Println(uuid.New())
	}
}
