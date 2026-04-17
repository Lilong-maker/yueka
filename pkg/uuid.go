package pkg

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func XueHua() {
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
