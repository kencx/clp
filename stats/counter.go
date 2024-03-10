package stats

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/kencx/clp/entry"
)

type Counter map[string]int

// count unique occurences of given field in Entries
func NewCounter(entries entry.Entries, field string) (Counter, error) {
	counter := make(Counter)

	for _, entry := range entries {
		key, err := getStructField(entry, field)
		if err != nil {
			return nil, err
		}

		if key == "" {
			continue
		}

		_, ok := counter[key]
		if !ok {
			counter[key] = 1
		} else {
			counter[key]++
		}
	}

	return counter, nil
}

func NewCounterFilteredTime(entries entry.Entries, field string, period string) (Counter, error) {
	filtered, err := FilterByPeriod(entries, period)
	if err != nil {
		return nil, err
	}
	return NewCounter(filtered, field)
}

// dynamically get struct field
func getStructField(v interface{}, field string) (string, error) {
	rv := reflect.ValueOf(v)

	if rv.Kind() == reflect.Ptr {
		// deference pointer
		rv = rv.Elem()
	}

	fv := rv.FieldByName(field)
	if !fv.IsValid() {
		// handle nested struct field
		for i := 0; i < rv.NumField(); i++ {
			val := rv.Field(i)
			if val.Kind() == reflect.Struct {
				return getStructField(val.Interface(), field)
			}
		}
		return "", fmt.Errorf("invalid field name: %s", field)
	}

	switch fv.Kind() {
	case reflect.Slice:
		if fv.Len() > 0 {
			elem := fv.Index(0)
			return elem.String(), nil
		}
		return "", nil

	case reflect.Int:
		return strconv.FormatInt(fv.Int(), 10), nil
	case reflect.Float64:
		return strconv.FormatFloat(fv.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(fv.Bool()), nil
	default:
		return fv.String(), nil
	}
}

func (c Counter) PrintTopN(n int) {
	for _, index := range c.TopN(n) {
		fmt.Printf("%s -> %d\n", index, c[index])
	}
}

func (c Counter) TopN(n int) []string {
	if n >= len(c) {
		return c.rank()
	}

	return c.rank()[:n]
}

// sort Counter by count
func (c Counter) rank() []string {
	type kv struct {
		Key   string
		Value int
	}
	var ss []kv

	for k, v := range c {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	ranked := make([]string, len(c))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked
}
