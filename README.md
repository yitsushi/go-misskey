# Misskey Go SDK

WIP ;)

Official Misskey API Documentation: https://misskey.io/api-doc

## Progress

 - [ ] [admin](https://misskey.io/api-doc#tag/admin) [#21](https://github.com/yitsushi/go-misskey/issues/21)
 - [X] [meta](https://misskey.io/api-doc#tag/meta)
 - [ ] [antennas](https://misskey.io/api-doc#tag/antennas) [#3](https://github.com/yitsushi/go-misskey/issues/3)
        Done: Create, List, Delete
 - [ ] [federation](https://misskey.io/api-doc#tag/federation) [#4](https://github.com/yitsushi/go-misskey/issues/4)
 - [ ] [app](https://misskey.io/api-doc#tag/app) [#5](https://github.com/yitsushi/go-misskey/issues/5)
 - [ ] [auth](https://misskey.io/api-doc#tag/auth) [noissue created]
 - [ ] [account](https://misskey.io/api-doc#tag/account) [no issue created]
 - [ ] [channels](https://misskey.io/api-doc#tag/channels) [no issue created]
 - [ ] [notes](https://misskey.io/api-doc#tag/notes) [#6](https://github.com/yitsushi/go-misskey/issues/6)
 - [ ] [charts](https://misskey.io/api-doc#tag/charts) [#7](https://github.com/yitsushi/go-misskey/issues/7)
 - [ ] [clips](https://misskey.io/api-doc#tag/clips) [#8](https://github.com/yitsushi/go-misskey/issues/8)
 - [ ] [drive](https://misskey.io/api-doc#tag/drive) [#9](https://github.com/yitsushi/go-misskey/issues/9)
 - [ ] [following](https://misskey.io/api-doc#tag/following) [#10](https://github.com/yitsushi/go-misskey/issues/10)
 - [ ] [games](https://misskey.io/api-doc#tag/games) [#11](https://github.com/yitsushi/go-misskey/issues/11)
 - [ ] [hashtags](https://misskey.io/api-doc#tag/hashtags) [#12](https://github.com/yitsushi/go-misskey/issues/12)
 - [ ] [messaging](https://misskey.io/api-doc#tag/messaging) [#13](https://github.com/yitsushi/go-misskey/issues/13)
 - [ ] [reactions](https://misskey.io/api-doc#tag/reactions) [#14](https://github.com/yitsushi/go-misskey/issues/14)
 - [ ] [notifications](https://misskey.io/api-doc#tag/notifications) [#15](https://github.com/yitsushi/go-misskey/issues/15)
 - [ ] [pages](https://misskey.io/api-doc#tag/pages) [#16](https://github.com/yitsushi/go-misskey/issues/16)
 - [ ] [users](https://misskey.io/api-doc#tag/users) [#17](https://github.com/yitsushi/go-misskey/issues/17)
 - [ ] [room](https://misskey.io/api-doc#tag/room) [#18](https://github.com/yitsushi/go-misskey/issues/18)
 - [ ] [groups](https://misskey.io/api-doc#tag/groups) [#19](https://github.com/yitsushi/go-misskey/issues/19)
 - [ ] [list](https://misskey.io/api-doc#tag/lists) [#20](https://github.com/yitsushi/go-misskey/issues/20)


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
