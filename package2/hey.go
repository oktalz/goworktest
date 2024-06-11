package main

import (
	"github.com/google/uuid"
	"github.com/oktalz/goworktest/package1"
)

func UUID() string {
	u := uuid.New()
	return u.String()
}

func Hey() string {
	return package1.UUID()
}

func main() {
	print(package1.UUID())
}
