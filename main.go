package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + " "
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return content, nil
}

func extractKeywords(jobDescription string) []string {
	words := strings.Fields(jobDescription)
	keywords := make(map[string]struct{})
	for _, word := range words {
		word = strings.ToLower(strings.TrimSpace(word))
		if word != "" {
			keywords[word] = struct{}{}
		}
	}

	var keywordList []string
	for keyword := range keywords {
		keywordList = append(keywordList, keyword)
	}
	return keywordList
}

func matchKeywords(resume string, keywords []string) int {
	resumeWords := strings.Fields(resume)
	keywordCount := 0

	for _, word := range keywords {
		for _, resumeWord := range resumeWords {
			if strings.ToLower(resumeWord) == word {
				keywordCount++
				break
			}
		}
	}

	return keywordCount
}

func calculateMatchPercentage(keywordCount, totalKeywords int) float64 {
	if totalKeywords == 0 {
		return 0.0
	}
	return (float64(keywordCount) / float64(totalKeywords)) * 100
}

func main() {
	
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current working directory:", dir)

	
	jobDescriptionFile := "jobdescription.txt"
	resumeFile := "resume.txt"

	
	jobDescription, err := readFile(jobDescriptionFile)
	if err != nil {
		fmt.Println("Error reading job description:", err)
		return
	}

	
	resume, err := readFile(resumeFile)
	if err != nil {
		fmt.Println("Error reading resume:", err)
		return
	}

	
	keywords := extractKeywords(jobDescription)

	// Match keywords in the resume
	keywordCount := matchKeywords(resume, keywords)

	// Calculate match percentage
	matchPercentage := calculateMatchPercentage(keywordCount, len(keywords))

	// Display results
	fmt.Printf("Job Description Keywords: %d\n", len(keywords))
	fmt.Printf("Matched Keywords: %d\n", keywordCount)
	fmt.Printf("Match Percentage: %.2f%%\n", matchPercentage)
}
