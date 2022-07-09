package main

import (
	"os"
	"strings"
    // "strconv"
)

func check(e error){
    if(e != nil){
        panic(e)
    } 
}

func main() {
    data, err := os.ReadFile("input.txt")
    check(err);
    content := string(data)
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
