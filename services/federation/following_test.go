package federation_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/federation"
)

func ExampleService_Following() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Following(&federation.FollowingRequest{
		Limit: 100,
	})
	if err != nil {
		log.Printf("[Federation/Following] %s", err)

		return
	}

	log.Printf("[Federation/Following] %v listed", resp)
}
