package main

import (
	"os"
	"strings"
    "strconv"
)
/*
 You're in an oddly shaped room, 
 there are squares on the floor and you move one square at a time.
 The room looks like this:

  ##  
 #### 
######
######
 #### 
  ##  

This is a six-by-six area defined in indices from 0 to 5 on each axis.
You start in the first # in the top row, or position 0 2 in indices,
and recieve a series of instructions to step up (U) left (L) right (R) or down (D) on the map above.
You can't step outside the # - if you're given an instruction to do so, 
ignore it, and move on to the next instruction.

For example, with input UDRR, you eventually run out of instructions at position 1 4.

After processing all the movements in your input,
what is the sum of the indices of each position you finished on at each step (including steps where you did not move)?

For example, with input UDRR, you would start on 0 2, stay on 0 2 then move through 1 2, 1 3 and 1 4. 

The sum of these positions is 14 - the first position is not counted.
*/

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
