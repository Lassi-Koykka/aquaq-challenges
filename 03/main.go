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

var roomStr []string = []string {
    "  ##  ",
    " #### ",
    "######",
    "######",
    " #### ",
    "  ##  ",
}

func main() {
    test := 0

    room := [][]string{}
    for _, row := range roomStr {
        room = append(room, strings.Split(row, ""))
    }

    content := "UDRR"
    if (test == 0) {
        data, err := os.ReadFile("input.txt")
        check(err);
        content = string(data)
    }
    strings.ReplaceAll(content, "\n", "")
    instructions := strings.Split(content, "")
    x := 0
    y := 2

    result := 0
    for _, dir := range instructions {
        if dir == "U" && x > 0 && room[x-1][y] == "#" {
            x -= 1
        } else if dir == "D" && x < len(room[x]) - 1 && room[x+1][y] == "#" {
            x += 1
        } else if dir == "L" && y > 0 && room[x][y-1] == "#" {
            y -= 1
        } else if dir == "R" && y < len(room[y]) - 1 && room[x][y+1] == "#" {
            y += 1
        } else if(!strings.Contains("UDLR", dir)) {
            continue
        }
        println(x, y)
        result += x + y
    }

    println(result)

}
