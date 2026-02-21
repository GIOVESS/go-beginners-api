package main

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Hello Giovanni endpoint
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, I'm Giovanni!",
		})
	})

	// --- Joke API ---
	jokes := []struct {
		Setup string `json:"setup"`
		Punch string `json:"punch"`
		Topic string `json:"topic"`
	}{
		{"Why do Go developers prefer the beach?", "Because of the shore (sure) types.", "go"},
		{"How do you comfort a JavaScript bug?", "You console it.", "programming"},
		{"Why did the API go to therapy?", "Too many internal server errors.", "api"},
		{"What's a programmer's favorite hangout?", "The Foo Bar.", "programming"},
		{"Why did the function break up with the variable?", "It had too many arguments.", "go"},
	}
	r.GET("/joke", func(c *gin.Context) {
		topic := strings.ToLower(c.DefaultQuery("topic", ""))
		var pool []int
		for i := range jokes {
			if topic == "" || jokes[i].Topic == topic {
				pool = append(pool, i)
			}
		}
		if len(pool) == 0 {
			pool = []int{0, 1, 2, 3, 4}
		}
		j := jokes[pool[rand.Intn(len(pool))]]
		c.JSON(http.StatusOK, gin.H{"setup": j.Setup, "punch": j.Punch, "topic": j.Topic})
	})

	// --- Minimal chatbot (canned responses) ---
	responses := map[string]string{
		"hello":       "Hi! I'm a tiny API bot. Try /joke for a joke, or ask 'who are you?'",
		"hi":          "Hey! Ask me for a joke with GET /joke",
		"joke":        "Use the API: GET http://localhost:8080/joke",
		"who are you": "I'm the go-beginners-api bot. I say hello and tell bad jokes.",
		"bye":         "Bye! Don't forget to run .\\run.ps1 again.",
	}
	r.GET("/chat", func(c *gin.Context) {
		q := strings.ToLower(strings.TrimSpace(c.DefaultQuery("q", "")))
		if q == "" {
			c.JSON(http.StatusOK, gin.H{
				"reply":        "Send ?q=hello or ?q=joke to chat.",
				"your_message": "",
			})
			return
		}
		reply, ok := responses[q]
		if !ok {
			reply = "I only know: hello, hi, joke, who are you, bye. Try one of those!"
		}
		c.JSON(http.StatusOK, gin.H{"reply": reply, "your_message": q})
	})

	if err := r.Run(); err != nil {
		return
	}
}
