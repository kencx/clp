package main

import (
	"bufio"
	"encoding/json"
	"io"
)

func parseFile(f io.Reader) (Entries, error) {
	var results Entries
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var res Entry

		err := json.Unmarshal(scanner.Bytes(), &res)
		if err != nil {
			return nil, err
		}

		if err := scanner.Err(); err != nil {
			return nil, err
		}

		results = append(results, res)
	}
	return results, nil
}
