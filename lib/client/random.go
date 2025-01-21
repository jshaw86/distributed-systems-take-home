package client 

import (
    "fmt"
    "math/rand"
    "strings"
)

var userAgentList = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7)",
	"Mozilla/5.0 (X11; Linux x86_64)",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X)",
    "Mozilla/5.0 (Linux; Android 13; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
    "Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
    "Mozilla/5.0 (Linux; Android 13; SM-S901B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
}

const (
    charset = "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
)

func RandomUserAgent() string {
	return userAgentList[rand.Intn(len(userAgentList))]
}

func RandomIP() string {
    return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func RandomBadActor() int {
    return rand.Intn(100)
}

func DecisionBasedOnProbability(probability float64) bool {
	randomValue := rand.Float64() * 100
	return randomValue < probability
}

func RandomUserID() string {
    length := 10 
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset)-1)]
    }

    return string(b)
}

func AnomolyProbability(badActor int) bool {
    return DecisionBasedOnProbability(float64(badActor) * rand.Float64())
}

func StringWithCharset(words int) string {
    var sentence []string 
    for i := 0; i < words; i++ {
        length := rand.Intn(10)
        b := make([]byte, length)
        for i := range b {
            b[i] = charset[rand.Intn(len(charset)-1)]
        }
        sentence = append(sentence, string(b)) 
    }
    return strings.Join(sentence, " ") 
}
