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

var buttons = map[string]string{
    "1": "",
    "2": "abc",
    "3": "def",
    "4": "ghi",
    "5": "jkl",
    "6": "mno",
    "7": "pqrs",
    "8": "tuv",
    "9": "wxyz",
    "0": " ",
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
        button := x[0]
        presses, err := strconv.ParseInt(x[1], 10, 0)
        check(err)
        // println("Button", button, "Presses", presses)
        char := strings.Split(buttons[button], "")[presses - 1]
        result += char
    }

    println(result)

}
