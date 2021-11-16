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

		columns := []string{}
		// generate columns
		if len(os.Args) > 2 {
			for _, k := range os.Args[2:] {
				fmt.Printf(" --column=%s", k)
				columns = append(columns, k)
			}
		} else {
			for k := range data[0] {
				fmt.Printf(" --column=%s", k)
				columns = append(columns, k)
			}
		}

		// generate data
		for _, d := range data {
			for _, c := range columns {
				s := string(d[c])
				s = strings.TrimSuffix(s, "\"")
				s = strings.TrimPrefix(s, "\"")

				s = strings.ReplaceAll(s, " ", "\\ ")
				fmt.Printf(" %s", s)
			}
		}
	}
}
