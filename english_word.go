package count

import (
	"regexp"
	"strings"
)

// CountEnglishWord get the count of english word in string, output is int
func CountEnglishWord(s string) int {
	// 定义一个正则表达式匹配英文单词
	re := regexp.MustCompile(`[a-zA-Z]+`)
	// 查找所有匹配的单词
	words := re.FindAllString(s, -1)
	// 返回单词的个数
	return len(words)
}

// GetEnglishWordCount get the count of english word in string, output is map[string]int
func GetEnglishWordCount(s string) map[string]int {
	// 定义一个正则表达式匹配英文单词
	re := regexp.MustCompile(`[a-zA-Z]+`)
	// 查找所有匹配的单词
	words := re.FindAllString(s, -1)
	wordMap := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(word)
		wordMap[word]++
	}
	return wordMap
}
