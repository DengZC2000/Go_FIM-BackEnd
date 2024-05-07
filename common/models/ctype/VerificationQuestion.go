package ctype

type VerificationQuestion struct {
	Problem1 *string `json:"problem1"`
	Problem2 *string `json:"problem2"`
	Problem3 *string `json:"problem3"`

	Answer1 *string `json:"answer1"`
	Answer2 *string `json:"answer2"`
	Answer3 *string `json:"answer3"`
}
