[![Go Report Card](https://goreportcard.com/badge/github.com/yitsushi/go-misskey)](https://goreportcard.com/report/github.com/yitsushi/go-misskey)
[![Coverage Status](https://coveralls.io/repos/github/yitsushi/go-misskey/badge.svg?branch=coverage)](https://coveralls.io/github/yitsushi/go-misskey?branch=coverage)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/yitsushi/go-misskey)

# Misskey Go SDK

WIP ;)

Official Misskey API Documentation: https://misskey.io/api-doc

## Progress

| Status | Endpoint Group | Implementation Issue | Note |
|--------|----------------|----------------------|------|
| :x: | [admin](https://misskey.io/api-doc#tag/admin) | [#21](https://github.com/yitsushi/go-misskey/issues/21) ||
| :white_check_mark: | [meta](https://misskey.io/api-doc#tag/meta) | ||
| :white_check_mark: | [antennas](https://misskey.io/api-doc#tag/antennas) | [#3](https://github.com/yitsushi/go-misskey/issues/3) ||
| :x: | [federation](https://misskey.io/api-doc#tag/federation) | [#4](https://github.com/yitsushi/go-misskey/issues/4) ||
| :x: | [app](https://misskey.io/api-doc#tag/app) | [#5](https://github.com/yitsushi/go-misskey/issues/5) ||
| :x: | [auth](https://misskey.io/api-doc#tag/auth) |||
| :x: | [account](https://misskey.io/api-doc#tag/account) |||
| :x: | [channels](https://misskey.io/api-doc#tag/channels) |||
| :x: | [notes](https://misskey.io/api-doc#tag/notes) | [#6](https://github.com/yitsushi/go-misskey/issues/6) ||
| :x: | [charts](https://misskey.io/api-doc#tag/charts) | [#7](https://github.com/yitsushi/go-misskey/issues/7) ||
| :x: | [clips](https://misskey.io/api-doc#tag/clips) | [#8](https://github.com/yitsushi/go-misskey/issues/8) ||
| :white_check_mark: | [drive](https://misskey.io/api-doc#tag/drive) | [#9](https://github.com/yitsushi/go-misskey/issues/9) ||
| :x: | [following](https://misskey.io/api-doc#tag/following) | [#10](https://github.com/yitsushi/go-misskey/issues/10) ||
| :x: | [games](https://misskey.io/api-doc#tag/games) | [#11](https://github.com/yitsushi/go-misskey/issues/11) ||
| :x: | [hashtags](https://misskey.io/api-doc#tag/hashtags) | [#12](https://github.com/yitsushi/go-misskey/issues/12) ||
| :x: | [messaging](https://misskey.io/api-doc#tag/messaging) | [#13](https://github.com/yitsushi/go-misskey/issues/13) ||
| :x: | [reactions](https://misskey.io/api-doc#tag/reactions) | [#14](https://github.com/yitsushi/go-misskey/issues/14) ||
| :x: | [notifications](https://misskey.io/api-doc#tag/notifications) | [#15](https://github.com/yitsushi/go-misskey/issues/15) ||
| :x: | [pages](https://misskey.io/api-doc#tag/pages) | [#16](https://github.com/yitsushi/go-misskey/issues/16) ||
| :x: | [users](https://misskey.io/api-doc#tag/users) | [#17](https://github.com/yitsushi/go-misskey/issues/17) ||
| :x: | [room](https://misskey.io/api-doc#tag/room) | [#18](https://github.com/yitsushi/go-misskey/issues/18) ||
| :x: | [groups](https://misskey.io/api-doc#tag/groups) | [#19](https://github.com/yitsushi/go-misskey/issues/19) ||
| :x: | [list](https://misskey.io/api-doc#tag/lists) | [#20](https://github.com/yitsushi/go-misskey/issues/20) ||


## How to use

For detailed examples, check the `example` directory.

```go
package main

import (
  "log"

  "github.com/sirupsen/logrus"
  "github.com/yitsushi/go-misskey"
  "github.com/yitsushi/go-misskey/core"
  "github.com/yitsushi/go-misskey/services/meta"
)

func main() {
  client := misskey.NewClient("https://slippy.xyz", "my misskey token")
  client.LogLevel(logrus.DebugLevel)

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
