package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
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
	test := 1

	content := "1 4 3 2 4 7 2 6 3 6"
	if test == 0 {
		data, err := os.ReadFile("input.txt")
		check(err)
		content = string(data)
	}

	content = strings.ReplaceAll(content, "\n", "")
	strs := strings.Split(content, " ")
	nums := strsToInts(strs)
	fmt.Println(nums)

	for i := 0; i < len(nums); {
		duplicate := false
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] == nums[j] {
                // println("DUPLICATE:", nums[i], "i:", i, "j:", j )
				duplicate = true
                newNums := []int{}
				if i > 0 && j < len(nums) - 1 {
                    newNums = append(nums[:i+1], nums[j+1:]...)
				} else if j < len(nums) - 1 {
                    newNums = nums[j:]
				} else if i > 0  {
                    newNums = nums[:i+1]
				}
                nums = newNums
				// fmt.Println(nums, len(nums))
				break
			}
		}
		if !duplicate {
			i++
		}

	}


	fmt.Println(nums)
    result := 0

    for _, num := range nums {
        result += num
    }

	println(result)

}
