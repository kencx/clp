package file

import (
	"encoding/json"
	"io"

	"github.com/kencx/clp/entry"
)

func Decode(f io.Reader) (entry.Entries, error) {
	var results entry.Entries
	dec := json.NewDecoder(f)

	for {
		var res entry.Entry
		if err := dec.Decode(&res); err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	return results, nil
}
