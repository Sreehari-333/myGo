package main

import (
	"log"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

type Session struct {
	UserID    string
	CreatedAt time.Time
}

var (
	sessionStore = make(map[string]Session)
	mu           sync.Mutex
)

func AddSession(sessionID, userID string) {
	mu.Lock()
	defer mu.Unlock()

	sessionStore[sessionID] = Session{
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	log.Printf("Session added : %s for user %s\n", sessionID, userID)

}

func CleanExpiredSessions() {
	mu.Lock()
	defer mu.Unlock()

	now := time.Now()
	for id, session := range sessionStore {
		if now.Sub(session.CreatedAt) > 1*time.Minute {
			delete(sessionStore, id)
			log.Printf("Session %s expired and removed\n", id)
		}
	}
}

func main() {
	AddSession("sess1", "user123")
	AddSession("sess2", "user456")

	c := cron.New()

	c.AddFunc("@every 10s", CleanExpiredSessions)
	c.Start()

	log.Println("Cron job started. Cleaning sessions every 10 second.")

	select {}
}
