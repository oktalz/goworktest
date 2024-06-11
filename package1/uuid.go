package package1

import "github.com/google/uuid"

func UUID() string {
	u := uuid.New()
	return u.String() + "*****************" + u.String()
}
