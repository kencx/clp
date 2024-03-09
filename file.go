package main

import (
	"encoding/json"
	"io"
)

func decodeFile(f io.Reader) (Entries, error) {
	var results Entries
	dec := json.NewDecoder(f)

	for {
		var res Entry
		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
