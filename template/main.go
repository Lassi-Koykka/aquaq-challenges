package main

import (
	"os"
	"strings"
    "strconv"
)

func check(e error){
    if(e != nil){
        panic(e)
    } 
}

func strsToInts(strs []string) []int {
	result := []int{}
	for _, str := range strs {
		num, err := strconv.ParseInt(str, 10, 0)
		check(err)
		result = append(result, int(num))
	}
	return result
}

func main() {
    test := 0

    content := "testContent"
    if (test == 0) {
        data, err := os.ReadFile("input.txt")
        check(err);
        content = string(data)
    }
    // strings.ReplaceAll(content, "\n", "")
    lines := strings.Split(content, "\n")
    result := ""
    for _, line := range lines {
        x := strings.Split(line, " ")
        if(len(x) < 2) {
            continue
        }
    }

    println(result)

}
