package main

import (
	"fmt"
	"time"

	"github.com/JeffersonTupinamba/api-social-meli/internal/follow"
)

func main() {
	follow := follow.Follow{
		UserID:     "1",
		FollowedID: "2",
		CreatedAt:  time.Now(),
	}
	fmt.Println("Hello, World!")
}
