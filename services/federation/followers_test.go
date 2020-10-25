package federation_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/federation"
)

func ExampleService_Followers() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Followers(federation.FollowersRequest{
		Limit: 100,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Followers] %s", err)

		return
	}

	log.Printf("[Federation/Followers] %v listed", resp)
}
