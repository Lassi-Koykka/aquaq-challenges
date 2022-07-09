package main

import (
	"os"
	"strings"
)

func check(e error){
    if(e != nil){
        panic(e)
    } 
}

//  Set the string's non-hexadecimal characters to 0.
//  Pad the string length to the next multiple of 3.
//  Split the result into 3 equal sections.
//  The first two digits of each remaining section are the hex components.

const hexChars string = "0123456789abcdef"

func main() {
    test := 0

    content := "kdb4life"
    if (test == 0) {
        data, err := os.ReadFile("input.txt")
        check(err);
        content = string(data)
    }
    content = strings.ReplaceAll(content, "\n", "")
    strChars := strings.Split(content, "")

    for i := 0; i < len(strChars); i++ {
        char := strings.ToLower(strChars[i])
        if !strings.Contains(hexChars, char) {
            strChars[i] = "0"
        }
    }
    if len(strChars) % 3 != 0 {
        difference := 3 - ( len(strChars) % 3 )
        for i := 0; i < difference; i++ {
            strChars = append(strChars, "0")
        }
    }

    println(strings.Join(strChars, ""))

    third := len(strChars) / 3
    componentIndexes := []int {0, 1, third, third + 1, third * 2, third * 2 + 1}
    result := ""
    for _, idx := range componentIndexes {
        result += strChars[idx]
    }
    println(result)

}
