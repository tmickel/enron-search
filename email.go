package main

import (
	"strings"
)

type Email struct {
	Filename string
	From     string
	To       string
	Subject  string
	Body     string
}

func NewEmail(filename, raw string) *Email {
	email := &Email{Filename: filename}
	lines := strings.Split(raw, "\n")
	headers := true
	for _, line := range lines {
		stripped := strings.TrimSpace(line)
		if stripped == "" {
			headers = false
			continue
		}
		if headers {
			kv := strings.Split(stripped, ": ")
			if len(kv) < 2 { // empty header
				continue
			}
			k, v := kv[0], kv[1]
			if k == "From" {
				email.From = v
			}
			if k == "To" {
				email.To = v
			}
			if k == "Subject" {
				email.Subject = v
			}
			// skip other headers for now
		} else {
			email.Body += line + "\n"
		}
	}
	return email
}
