package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	path := "strings.swift"

	if len(os.Args) >= 2 {
		path = os.Args[1]
	}

	source, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// NSLocalizedString("UpdatePin.updateTitle", value:"Update PIN", comment: "Update Pin title")

	NSLocalized := regexp.MustCompile(`(?m)NSLocalizedString\((.*)\)`)
	hits := NSLocalized.FindAllStringSubmatch(string(source), -1)

	if len(hits) == 0 {
		fmt.Println("Found no instances of NSLocalizedString")
		return
	}

	type Entry struct {
		comment string
		key     string
		value   string
	}

	entries := make(map[string]Entry)
	var keys []string

	for _, entry := range hits {
		var e Entry

		details := strings.TrimSpace(entry[1])
		key := regexp.MustCompile(`^"([^"]*)"`).FindStringSubmatch(details)
		if len(key) < 2 {
			fmt.Println("No key found in ", entry[0])
			return
		}
		e.key = key[1]

		others := regexp.MustCompile(`,([ a-z:]*)"([^"]*)"`).FindAllStringSubmatch(details, -1)

		for i, v := range others {
			switch strings.Replace(v[1], ` `, ``, -1) {
			case `value:`:
				e.value = v[2]
				continue
			case `comment:`:
				e.comment = v[2]
				continue
			}
			switch i {
			case 3:
				e.value = v[2]
			case 4:
				e.comment = v[2]
			}
		}

		if e.value == `` {
			fmt.Println("No value found in ", entry[0])
			return
		}
		if e.comment == `` {
			fmt.Println("Warning: No comment found in ", entry[0])
			e.comment = `No comment provided.`
		}
		entries[e.key] = e
		keys = append(keys, e.key)
	}

	// Sort keys
	sort.Strings(keys)

	var output string

	for i := range keys {
		output += `/* ` + entries[keys[i]].comment + ` */
"` + entries[keys[i]].key + `" = "` + entries[keys[i]].value + `";

`
	}

	err = ioutil.WriteFile("Localizable.strings", []byte(output), 0644)
	if err == nil {
		fmt.Println("All done!")
	} else {
		fmt.Println("Error writing file:", err)
	}

}
