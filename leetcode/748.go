package leetcode

import (
	"fmt"
	"strings"
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	wordsLen := len(words)
	minword := ""
	for i:=0;i<wordsLen;i++{
		hashtable := make([]int, 26)
		for j:=0;j<len(words[i]);j++ {
			hashtable[words[i][j]-'a'] ++
		}
		if match(licensePlate, hashtable) {
			if minword == "" {
				minword = words[i]
			}else if len(minword) > len(words[i]) {
				minword = words[i]
			}
		}
	}
	return minword
}


func match(licensePlate string, hashtable []int) bool {
	lowerLicensePlate := strings.ToLower(licensePlate)
	for i:=0;i<len(lowerLicensePlate);i++ {
		if letter := lowerLicensePlate[i]; unicode.IsLetter(rune(letter)) {
			fmt.Println(string(letter), rune(letter)-'a')
			hashtable[rune(letter) - 'a'] --
			if hashtable[rune(letter)- 'a'] < 0 {
				return false
			}
		}
	}
	return true
}