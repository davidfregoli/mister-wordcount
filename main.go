package main

import (
	"strconv"
	"strings"
	"unicode"

	mr "fregoli.dev/mister"
)

type WordCount struct{}

func (app *WordCount) Map(filename string, contents string) []mr.KeyValue {
	// function to detect word separators.
	ff := func(r rune) bool { return !unicode.IsLetter(r) }

	// split contents into an array of words.
	words := strings.FieldsFunc(contents, ff)

	kva := []mr.KeyValue{}
	for _, w := range words {
		kv := mr.KeyValue{Key: w, Value: "1"}
		kva = append(kva, kv)
	}
	return kva
}

func (app *WordCount) Reduce(key string, values []string) string {
	return strconv.Itoa(len(values))
}

func main() {
	app := new(WordCount)
	mr.RegisterApp(app)
}
