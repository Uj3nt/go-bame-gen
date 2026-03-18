package bame

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

func getNamesFromFile(filename string) []string {
	file, _ := os.Open(filename)
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ToLower(line)
		words = append(words, strings.Fields(line)...)
	}

	return words
}


func getTransitionMatrix(names []string) map[rune][]rune {
	matrix := make(map[rune][]rune)

	for _, name := range names {
		runes := []rune(name)
		for i := 0; i < len(runes)-1; i++ {
			if matrix[runes[i]] == nil {
				matrix[runes[i]] = make([]rune, 0)
			}

			matrix[runes[i]] = append(matrix[runes[i]], runes[i+1])
		}
	}

	return matrix
}


func сapitan(s string) string {
	if len(s) == 0 {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}


func genBame(size int, matrix map[rune][]rune) (string) {
	var s string
	start_pull := make([]rune, 0, len(matrix))
	for key := range matrix {
		start_pull = append(start_pull, key)
	}

	random_key := start_pull[rand.Intn(len(matrix))]
	s = s + string(random_key)

	for i := 0; i < size-1; i++ {
		size_cur := len(matrix[random_key])
		if size_cur < 1 {
			random_key := start_pull[rand.Intn(len(matrix))]
			s = s + string(random_key)
		} else {
			next := rand.Intn(size_cur)
			random_key = matrix[random_key][next]
			s = s + string(random_key)
		}
	}

	return сapitan(s)
}


func GenBameFromFile(len int, filename string) (string) {
	matrix := getTransitionMatrix(getNamesFromFile(filename))
	return  genBame(len, matrix)
}

