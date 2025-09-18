package file

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
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
			return nil, fmt.Errorf("failed to decode file: %w", err)
		}
		results = append(results, res)
	}
	return results, nil
}

func DecodeGz(f io.Reader) (entry.Entries, error) {
	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode gzip: %w", err)
	}
	defer gr.Close()

	return Decode(gr)
}
