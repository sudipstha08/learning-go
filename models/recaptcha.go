package models

import "time"

//RecaptchaResponse structure.
type RecaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

//RecaptchaRequest structure
type RecaptchaRequest struct {
	Response string `json:"response" binding:"required"`
}
