package apiservice

import "time"

type API interface {
	Request(endpoint string) string
}

type RateLimiter struct {
	apiService API
	timestamps []int64
}

const (
	MAX_REQUEST = 3
	TIME_WINDOW_MS = 10000
)

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		apiService: NewAPIService(),
		timestamps: []int64{},
	}
}

func (this *RateLimiter) Request(endpoint string) string {
	this.updateTimestamp()

	if len(this.timestamps) < MAX_REQUEST {
		this.timestamps = append(this.timestamps, time.Now().UnixMilli())
		return this.apiService.Request(endpoint) 
	}

	return "Rate limit exceded"
}

func (this *RateLimiter) updateTimestamp() {
	currTime := time.Now().UnixMilli()
	expireBoundary := currTime - TIME_WINDOW_MS

	validIndex := 0

	for i := 0; i < len(this.timestamps); i++ {
		if this.timestamps[i] > expireBoundary {
			validIndex = i
			break
		}

		if i == len(this.timestamps) - 1 {
			validIndex = len(this.timestamps)
		}
	}

	this.timestamps = this.timestamps[validIndex:]
}
