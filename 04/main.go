package main

import ("fmt")

func main() {
    test := 0

    value := 15
    if(test == 0) {
        value = 987820
    }

    getFactors := func(num int) []int {
        factors := []int {}
        for i := num - 1; i > 1; i-- {
            if num % i == 0 {
                factors = append(factors, i)
            }
        }
        return factors
    }

    valueFactors := getFactors(value)

    coPrimes := []int {}
    result := 0

    for i := 0; i < value; i++ {
        isCoPrime := true
        for _, factor := range valueFactors {
            if i % factor == 0 {
                isCoPrime = false
                break;
            }
        }
        if isCoPrime {
            result += i
            coPrimes = append(coPrimes, i)
        }
    } 

    fmt.Println(coPrimes)
    println(result)

}
