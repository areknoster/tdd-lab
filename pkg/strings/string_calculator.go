package strings

import (
    "log"
    "regexp"
    "strconv"
)

type StringCalculator struct {}

func (sc *StringCalculator) splitTokens(text string) []string{
    tokens := regexp.MustCompile("[\n,]+").Split(text, -1)
    return tokens
}

func (sc *StringCalculator) Add(numbers string) int {
    if numbers == "" {
        return 0
    }
    sum := 0
    tokens := sc.splitTokens(numbers)

    for i, token := range tokens {
        number, err := strconv.Atoi(token)
        if err != nil{
            log.Panicf("could not convert %dth token to number: %s", i, err.Error())
        }
        sum += number
    }
    return sum
}
