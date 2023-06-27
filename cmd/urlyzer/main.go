package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listFlag := flag.Bool("l", false, "read URLs from a file")
	exceptionFlag := flag.String("e", "", "comma-separated list of file extensions to exclude from wordlist")
	flag.Parse()
	printed := make(map[string]bool)

	var urls []string
	if *listFlag {
		filename := flag.Arg(0)
		if filename == "" {
			fmt.Fprintln(os.Stderr, "Error: no filename provided")
			os.Exit(1)
		}
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			urls = append(urls, scanner.Text())
		}
	}

	exceptions := make(map[string]bool)
	if *exceptionFlag != "" {
		extensions := strings.Split(*exceptionFlag, ",")
		for _, ext := range extensions {
			exceptions[strings.TrimSpace(ext)] = true
		}
	}

	wordList := make(map[string]bool)
	for _, u := range urls {
		parsed, err := url.Parse(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing URL %s: %s\n", u, err)
			continue
		}
		path := parsed.Path
		if path == "" {
			path = "/"
		}
		dir, file := filepath.Split(path)
		dir = strings.TrimPrefix(dir, "/")
		dir = strings.TrimSuffix(dir, "/")
		dirs := strings.Split(dir, "/")
		ext := filepath.Ext(file)
		base := strings.TrimSuffix(file, ext)

		endpointWords := strings.Split(strings.TrimSuffix(strings.TrimPrefix(path, "/"), ext), "/")
		for _, word := range endpointWords {
			words := strings.Split(word, "-")
			for _, w := range words {
				if w != "" && !exceptions[ext] {
					wordList[w] = true
					if ext != "" {
						wordList[w+ext] = true
					}
				}
			}
			if len(words) > 1 {
				for i := 0; i < len(words)-1; i++ {
					for j := i + 1; j < len(words); j++ {
						if words[i] != "" && words[j] != "" && !exceptions[ext] {
							wordList[words[i]+"-"+words[j]] = true
							if ext != "" {
								wordList[words[i]+"-"+words[j]+ext] = true
							}
							wordList[words[j]+"-"+words[i]] = true
							if ext != "" {
								wordList[words[j]+"-"+words[i]+ext] = true
							}
						}
					}
				}
			}
		}

		for _, dir := range dirs {
			words := strings.Split(dir, "-")
			for _, w := range words {
				if w != "" && !exceptions[ext] {
					wordList[w] = true
				}
			}
			if len(words) > 1 {
				for i := 0; i < len(words)-1; i++ {
					for j := i + 1; j < len(words); j++ {
						if words[i] != "" && words[j] != "" && !exceptions[ext] {
							wordList[words[i]+"-"+words[j]] = true
							wordList[words[j]+"-"+words[i]] = true
						}
					}
				}
			}
		}

		if base != "" && !exceptions[ext] {
			words := strings.Split(base, "-")
			for _, word := range words {
				if word != "" {
					wordList[word] = true
					if ext != "" {
						wordList[word+ext] = true
					}
				}
			}
			if len(words) > 1 {
				for i := 0; i < len(words)-1; i++ {
					for j := i + 1; j < len(words); j++ {
						if words[i] !="" && words[j] != "" && !exceptions[ext] {
							wordList[words[i]+"-"+words[j]] = true
							if ext != "" {
								wordList[words[i]+"-"+words[j]+ext] = true
							}
							wordList[words[j]+"-"+words[i]] = true
							if ext != "" {
								wordList[words[j]+"-"+words[i]+ext] = true
							}
						}
					}
				}
			}
		}

		query := parsed.Query()
		for key, values := range query {
			for _, value := range values {
				if key != "" && !exceptions[ext] {
					wordList[key] = true
					if value != "" {
						wordList[value] = true
					}
				}
			}
		}
	}

	for word := range wordList {
		if !strings.HasSuffix(word, "_") && !printed[word] && !strings.HasPrefix(word, "http://") && !strings.HasPrefix(word, "https://") {
			fmt.Println(word)
			printed[word] = true
	
			if ext := filepath.Ext(word); ext != "" {
				wordWithoutExt := strings.TrimSuffix(word, ext)
				if !strings.HasSuffix(wordWithoutExt, "_") && !printed[wordWithoutExt] && !strings.HasPrefix(wordWithoutExt, "http://") && !strings.HasPrefix(wordWithoutExt, "https://") {
					fmt.Println(wordWithoutExt)
					printed[wordWithoutExt] = true
				}
			}
		}
	}
}