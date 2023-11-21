// ratelimiter.go
package ratelimiter

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter defines the rate limiting configuration
type RateLimiter struct {
	capacity     int        // Maximum number of tokens in the bucket
	refillAmount int        // Number of tokens added to the bucket per second
	tokens       int        // Current number of tokens in the bucket
	mu           sync.Mutex // Mutex for synchronization
	lastRefill   time.Time  // Time of the last token refill
}

// NewRateLimiter creates a new RateLimiter instance with the specified capacity and refill amount
func NewRateLimiter(capacity, refillAmount int) *RateLimiter {
	return &RateLimiter{
		capacity:     capacity,
		refillAmount: refillAmount,
		tokens:       0,
		lastRefill:   time.Now(),
	}
}

// Middleware enforces rate limits using the provided RateLimiter instance
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Token refill logic
		rl.mu.Lock()
		defer rl.mu.Unlock()

		rl.refillTokens()

		// Check if the user has enough tokens
		if !rl.consumeToken() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Continue with the next handler if the user has enough tokens
		next.ServeHTTP(w, r)
	})
}

// Helper function to refill tokens in the bucket
func (rl *RateLimiter) refillTokens() {
	elapsed := time.Since(rl.lastRefill)
	tokensToAdd := int(elapsed.Seconds()) * rl.refillAmount

	// Refill the tokens up to the capacity
	rl.tokens = min(rl.capacity, rl.tokens+tokensToAdd)
	rl.lastRefill = time.Now()
}

// Helper function to consume tokens from the bucket
func (rl *RateLimiter) consumeToken() bool {
	if rl.tokens >= 1 {
		rl.tokens -= 1
		return true
	}
	return false
}

// Helper function to get the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
