package main

import (
	"fmt"
	"os"

	"github.com/deelawn/ajv-schema-gen/ajv"
)

func main() {

	schema, err := ajv.Generate(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := schema.String()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(out)
}
