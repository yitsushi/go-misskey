[![Go Report Card](https://goreportcard.com/badge/github.com/yitsushi/go-misskey)](https://goreportcard.com/report/github.com/yitsushi/go-misskey)
[![Coverage Status](https://coveralls.io/repos/github/yitsushi/go-misskey/badge.svg?branch=main)](https://coveralls.io/github/yitsushi/go-misskey?branch=main)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/yitsushi/go-misskey)
[![Chat on Matrix](https://matrix.to/img/matrix-badge.svg)](https://matrix.to/#/#go-misskey:matrix.org)

# Misskey Go SDK

Official Misskey API Documentation: https://misskey.io/api-doc

Check the `docs` directory for more information.

## Progress

| Status | Endpoint Group | Implementation Issue | Note |
|--------|----------------|----------------------|------|
| :warning: | [admin](https://misskey.io/api-doc#tag/admin) | [#21](https://github.com/yitsushi/go-misskey/issues/21) | In Progress (74%) |
| :white_check_mark: | [meta](https://misskey.io/api-doc#tag/meta) | ||
| :white_check_mark: | [antennas](https://misskey.io/api-doc#tag/antennas) | [#3](https://github.com/yitsushi/go-misskey/issues/3) ||
| :white_check_mark: | [federation](https://misskey.io/api-doc#tag/federation) | [#4](https://github.com/yitsushi/go-misskey/issues/4) ||
| :x: | [app](https://misskey.io/api-doc#tag/app) | [#5](https://github.com/yitsushi/go-misskey/issues/5) ||
| :x: | [auth](https://misskey.io/api-doc#tag/auth) |||
| :x: | [account](https://misskey.io/api-doc#tag/account) |||
| :x: | [channels](https://misskey.io/api-doc#tag/channels) |||
| :white_check_mark: | [notes](https://misskey.io/api-doc#tag/notes) | [#6](https://github.com/yitsushi/go-misskey/issues/6) ||
| :x: | [charts](https://misskey.io/api-doc#tag/charts) | [#7](https://github.com/yitsushi/go-misskey/issues/7) ||
| :white_check_mark: | [clips](https://misskey.io/api-doc#tag/clips) | [#8](https://github.com/yitsushi/go-misskey/issues/8) ||
| :white_check_mark: | [drive](https://misskey.io/api-doc#tag/drive) | [#9](https://github.com/yitsushi/go-misskey/issues/9) ||
| :white_check_mark: | [following](https://misskey.io/api-doc#tag/following) | [#10](https://github.com/yitsushi/go-misskey/issues/10) ||
| :white_check_mark: | [hashtags](https://misskey.io/api-doc#tag/hashtags) | [#12](https://github.com/yitsushi/go-misskey/issues/12) ||
| :x: | [messaging](https://misskey.io/api-doc#tag/messaging) | [#13](https://github.com/yitsushi/go-misskey/issues/13) ||
| :white_check_mark: | [reactions](https://misskey.io/api-doc#tag/reactions) | [#14](https://github.com/yitsushi/go-misskey/issues/14) ||
| :white_check_mark: | [notifications](https://misskey.io/api-doc#tag/notifications) | [#15](https://github.com/yitsushi/go-misskey/issues/15) ||
| :x: | [pages](https://misskey.io/api-doc#tag/pages) | [#16](https://github.com/yitsushi/go-misskey/issues/16) ||
| :x: | [users](https://misskey.io/api-doc#tag/users) | [#17](https://github.com/yitsushi/go-misskey/issues/17) ||
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

## How can I get a Misskey Token?

Navigate to `Settings > API` and there you generate a new token.

## How can I debug what's wrong?

There is a logging system, right now it's not very wide spread
in the codebase, but if you turn it on, you will be able to see:
 - all request with method, endpoint and body
 - all responds with status code, from what endpoint told and the body

To enable debug mode, just change the `LogLevel` to `DebugLevel`:

```go
client := misskey.NewClient("https://slippy.xyz", "my misskey token")
client.LogLevel(logrus.DebugLevel)
```

The output should look like this:
```
DEBU[0000] POST https://slippy.xyz/api/antennas/show     _type=request
DEBU[0000] {"antennaId":"8dbpybhulw","i":"my misskey token"}  _type=request
DEBU[0000] {"id":"8dbpybhulw","createdAt":"2020-10-13T16:03:22.674Z","name":"Genshin Impact","keywords":[["genshin"]],"excludeKeywords":[[""]],"src":"all","userListId":null,"userGroupId":null,"users":[""],"caseSensitive":false,"notify":false,"withReplies":true,"withFile":false,"hasUnreadNote":false}  _type=response code=200 from="https://slippy.xyz/api/antennas/show"
```
