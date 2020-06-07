package main

import (
	"fmt"
	"os"

	"github.com/mmd-afegbua/web-crawler/client"
)

func main() {
	err := client.New().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
