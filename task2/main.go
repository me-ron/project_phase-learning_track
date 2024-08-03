package main

import (
	"fmt"
	"strings"
)

func word_count(word string) map[string]int{
	count := make(map[string]int)
	word = strings.ToLower(word)
	for i := range len(word) {
		ord := int(word[i]) - int('a')
		if 0 <= ord  && ord <= 26{
			count[string(word[i])] ++
		}

	}

	return count
}

func palindrom(word string)bool{
	i := 0
	j := len(word) - 1
	for i < j{
		for i < j && int(word[i]) - int('a') < 0{
			i += 1
		}
		for i < j && int(word[j]) - int('a') < 0{
			j -= 1
		} 
		if word[i] != word[j] {
			return false
		}
		i ++
		j --
	}

	return true
}

func main(){
	fmt.Println(word_count("asdfghjk13n-"))
	fmt.Println(palindrom("abc-cba"))
	fmt.Println(palindrom("oioo9"))
}