package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		panic("Missing command argument")
	}
	switch args[1] {
	case "list":
		r := bufio.NewReader(os.Stdin)
		b, err := r.ReadBytes('\n')
		if err != nil {
			panic(err)
		}
		var data = []map[string]json.RawMessage{}
		err = json.Unmarshal(b, &data)
		if err != nil {
			panic(err)
		}
		if len(data) == 0 {
			return
		}

		// generate columns
		for k := range data[0] {
			fmt.Printf(" --column=\"%s\"", k)
		}

		// generate data
		for _, d := range data {
			for _, v := range d {
				fmt.Printf(" \"%s\"", strings.TrimPrefix(strings.TrimSuffix(string(v), "\""), "\""))
			}
		}
	}
}
