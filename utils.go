package main

import uuid "github.com/satori/go.uuid"

// GetUUID returns a unique ID
func GetUUID() string {
	u := uuid.NewV1().String()
	return u
}
