////////////////////////////////////////////////////////////////////////////////
// Appologies in advance for the overkill, I'm learning Go right now and wanted
// to noodle a little.
// Program takes an argument of a JSON file with a list of rules. 
// If no argument is provided, it uses the default rules. 
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// The json file has rules formatted like this, meaning that it's not just 
// divisiblity that it's testing, but what the remainder is in a list.
// [
//   {
//     "divisor": 13,
//     "word": "Chugga",
//     "remainders": [0]
//   },
//   {
//     "divisor": 6,
//     "word": "Choo",
//     "remainders": [ 1, 5 ]
//   }
// ]
////////////// output: 
// 1 Choo
// 2
// 3
// 4
// 5 Choo
// 6
// 7 Choo
// 8
// 9
// 10
// 11 Choo
// 12
// 13 ChuggaChoo
// 14
// 15
// 16
// 17 Choo
// 18
// 19 Choo
// 20
// 21
// 22
// 23 Choo
// 24
// 25 Choo
// 26 Chugga
// 27
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strconv"
	"os"
	"io/ioutil"
	"encoding/json"
	"slices"
)

type WordRule struct {
	Divisor    int    `json:"divisor"`
	Word       string `json:"word"`
	Remainders []int  `json:"remainders"`
}

func main() {
	args := os.Args[1:]
	wordRules := wordRules(args)
	fmt.Println(wordRules)

	count := 100
	for i := 1; i < count+1; i++ {
		result := ""
		for _, rule := range wordRules {
			if slices.Contains(rule.Remainders, i%rule.Divisor) {
				result += rule.Word
			}
		}
		if result == "" {
		   result = strconv.Itoa(i)
		}
		fmt.Println(result)
	}
}

func wordRules(args []string) []WordRule {
	if (len(args) == 0) {
		return defaultRules()
	}
	jsonFile, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rules []WordRule
	json.Unmarshal(byteValue, &rules)
	return rules
}

func defaultRules() []WordRule {
	return []WordRule{
		WordRule{3, "Crackle", []int{0}},
		WordRule{5, "Pop", []int{0}},
	}
}