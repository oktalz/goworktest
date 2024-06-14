package main

import (
	"github.com/google/uuid"
	"github.com/oktalz/goworktest/common"
)

func UUID() string {
	u := uuid.New()
	return u.String()
}

func Hey() string {
	return common.UUID()
}

func main() {
	print(common.UUID())
}
