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

var roomStr []string = []string {
    "  ##  ",
    " #### ",
    "######",
    "######",
    " #### ",
    "  ##  ",
}

func main() {
    test := 1

    room := [][]string{}
    for _, row := range roomStr {
        room = append(room, strings.Split(row, ""))
    }

    content := "LRDLU"
    if (test == 0) {
        data, err := os.ReadFile("input.txt")
        check(err);
        content = string(data)
    }
    strings.ReplaceAll(content, "\n", "")
    instructions := strings.Split(content, "")

    result := 0
    for _, dir := range instructions {
        println(dir)
        if dir == "U" {
        } else if dir == "D" {
        } else if dir == "L" {
        } else if dir == "R" {
        } else if(!strings.Contains("UDLR", dir)) {
            continue
        }
    }

    println(result)

}

