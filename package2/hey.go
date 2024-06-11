package package2

import (
	"github.com/google/uuid"
	"github.com/oktalz/goworktest/package1"
)

func UUID() string {
	u := uuid.New()
	return u.String()
}

func Hey() string {
	package1.UUID()
}
