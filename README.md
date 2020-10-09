# Misskey Go SDK

WIP ;)

Official Misskey API Documentation: https://misskey.io/api-doc

## How to use

For detailed examples, check the `example` directory.

```go
package main

import (
  "log"

  "github.com/yitsushi/go-misskey"
  "github.com/yitsushi/go-misskey/core"
  "github.com/yitsushi/go-misskey/services/meta"
)

func main() {
  client := misskey.NewClient("https://slippy.xyz", "my misskey token")

  stats, err := client.Meta().Stats()
  if err != nil {
    log.Printf("[Meta] Error happened: %s", err)
    return
  }

  log.Printf("[Stats] Instances:          %d", stats.Instances)
  log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
  log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
  log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
  log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
}
```
