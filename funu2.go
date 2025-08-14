package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var charToCode = map[rune]string{}

func initMapping() {
    mapping := map[rune]string{
        '2': "ABC",
        '3': "DEF",
        '4': "GHI",
        '5': "JKL",
        '6': "MNO",
        '7': "PQRS",
        '8': "TUV",
        '9': "WXYZ",
    }

    for digit, letters := range mapping {
        for i, ch := range letters {
            charToCode[ch] = strings.Repeat(string(digit), i+1)
        }
    }
}

func wordToCode(word string) string {
    var b strings.Builder
    for _, ch := range word {
        b.WriteString(charToCode[ch])
    }
    return b.String()
}

func main() {
    initMapping()

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    s := scanner.Text()

    scanner.Scan()
    n := 0
    fmt.Sscanf(scanner.Text(), "%d", &n)

    words := make([]string, n)
    for i := 0; i < n; i++ {
        scanner.Scan()
        words[i] = scanner.Text()
    }


    codeToWord := make(map[string]string)
    wordLengths := make(map[int]bool) 
    for _, w := range words {
        code := wordToCode(w)
        codeToWord[code] = w
        wordLengths[len(code)] = true
    }

    dp := make([]string, len(s)+1)
    dp[len(s)] = "" 

    for i := len(s) - 1; i >= 0; i-- {
        for length := range wordLengths {
            if i+length <= len(s) {
                segment := s[i : i+length]
                if word, ok := codeToWord[segment]; ok {
                    if dp[i+length] != "" || i+length == len(s) {
                        if dp[i+length] != "" {
                            dp[i] = word + " " + dp[i+length]
                        } else {
                            dp[i] = word
                        }
                        break
                    }
                }
            }
        }
    }

    fmt.Println(dp[0])
}