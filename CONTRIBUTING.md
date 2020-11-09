# Contributing to Go Misskey

🎉 First off, thanks for taking the time to contribute! 🎉

The following is a set of guidelines for contributing to Go Misskey.
These are mostly guidelines, not rules. Use your best judgment,
and feel free to propose changes to this document in a pull request.

Please read our [Code of Conduct][code-of-conduct].

#### Table Of Contents

[What should I know before I get started?](#what-should-i-know-before-i-get-started)

[How Can I Contribute?](#how-can-i-contribute)

* [Reporting Bugs](#reporting-bugs)
* [Suggesting Enhancements](#suggesting-enhancements)
* [Your First Code Contribution](#your-first-code-contribution)
* [Pull Requests](#pull-requests)

[Styleguides](#styleguides)

* [Git Commit Messages](#git-commit-messages)
* [Documentation Styleguide](#documentation-styleguide)

[Additional Notes](#additional-notes)
  * [Issue and Pull Request Labels](#issue-and-pull-request-labels)

## What should I know before I get started

Other then Go, not much. If you are present on any of the Misskey instances,
it may be a bit of an advantage, but you can contribute without any experiences
with Misskey.

However, implementing and testing endpoints requires an API token, as the
official documentation is not accurate. In some places it says there is no
response body with `201 No Content` status code, but in reality it's
`200 OK` and it has a response body.

We try to map 1:1 on endpoints with services under the `services` directory,
but sometimes it's better to create sub-packages or to rename functions.

## How Can I Contribute

### Reporting Bugs

This section guides you through submitting a bug report.
Following these guidelines helps maintainers and
the community to understand your report 📝, reproduce the behavior 💻 💻,
and find related reports 🔎.

We have a template for bugs, if you are reporting one, please
give us as much detail as you can and use the template as a guideline.
The `Bug report` template is not a strict template, you can format your report
in different ways as long as it helps understanding the issue better.

#### Before Submitting A Bug Report

Perform a [cursory search][issue-search] to see if the problem has already
been reported. If it has and the issue is still open, add a comment to
the existing issue instead of opening a new one.

### Suggesting Enhancements

This section guides you through submitting an enhancement suggestion,
including completely new features and minor improvements to existing
functionality. Following these guidelines helps maintainers and the
community to understand your suggestion 📝 and find related suggestions 🔎.

We have a template for `Feature requests`, if you are submitting an
enhancement suggestion, please give us as much detail as you can
and use the template as a guideline. The `Feature requests` template
is not a strict template, you can format your report in different ways
as long as it helps understanding the issue better.

#### Before Submitting An Enhancement Suggestion

Perform a [cursory search][issue-search] to see if the enhancement
has already been suggested. If it has, add a comment to the existing
issue instead of opening a new one.

### Your First Code Contribution

Unsure where to begin contributing?
You can start by looking through these `❤️ Good first issue`,
`📖 Documentation` and `🏳️ Help wanted` issues:

* [❤️ Good first issue][good-first-issue] - issues which should be easy to
    to implement based on existing code. They are mostly an endpoint or
    endpoint group with straightforward implementation.
* [🏳️ Help wanted][help-wanted] - issues which should be a bit more involved
    than `❤️ Good first issue` issues.
* [📖 Documentation][documentation] - issues which may not require deep understanding
    of the system, but requires to understand go.

#### Local development

As a requirement, you will need Go to be installed and a Misskey API token
to test service endpoints which need authentication.
If you don't have an API Token yet, you can create a new user on Slippy.
The registration requires only a `username` and a `password`
(yes, no `email` or `phone`).

To test your local changes, simply create a small Go program inside the
repository so it compiles with the code, running your modifications:

```go
package main

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func main() {
	client := misskey.NewClient(
        "https://slippy.xyz",
        os.Getenv("MISSKEY_TOKEN"),
    )
	client.LogLevel(logrus.DebugLevel)

	response, err := client.Notes().Show("8ebr49xcpj")
	if err != nil {
		log.Printf("Error happened: %s", err)

		return
	}

	for index, choice := range response.Poll.Choices {
		log.Printf("%2d %s", index, choice.Text)
	}
}
```

If you are using [Visual Studio Code][vscode] or any other
editors that has a REST Client functionality (built-in or plugin),
you can create a request file to test endpoints.

For [Vim][vim], I use [diepm/vim-rest-console][vim-rest-console]:

```rest
# Global scope.
https://slippy.xyz/api
Content-Type: yourtoken

-s
-i

apiToken = zU6RqyRosMtkJlXPdkjMUoMSzZr8h0bv
--

# List lists
--
POST /users/lists/list

{
  "i": ":apiToken"
}
```

For `vscode`, you can use [REST Client by Huachao Mao][humao-rest-client]:

```http
@host = slippy.xyz
@token = yourtoken

# @name create
POST https://{{host}}/api/notes/create
Content-Type: application/json

{
    "i": "{{token}}",
    "text": "Test message"
}

###

@noteID = {{create.response.body.$.createdNote.id}}

POST https://{{host}}/api/notes/delete
Content-Type: application/json

{
    "i": "{{token}}",
    "noteId": "{{noteID}}"
}
```

### Pull Requests

The process described here has several goals:

* Maintain quality
* Fix problems that are important to users
* Enable a sustainable system for maintainers to review contributions

Please follow these steps to have your contribution considered by the maintainers:

1. Follow all instructions in the [template][pr-template]
2. Follow the [styleguides](#styleguides).
3. Follow the [Patterns and Structure][patterns-and-structure] guide
4. After you submit your pull request, verify that all status checks
    are passing
    <details>
    <summary>What if the status checks are failing?</summary>
    If a status check is failing, and you believe that the failure is unrelated
    to your change, please leave a comment on the pull request explaining why
    you believe the failure is unrelated. A maintainer will re-run the
    status check for you. If we conclude that the failure was a false positive,
    then we will open an issue to track that problem with our status
    check suite.
    </details>

While the prerequisites above must be satisfied prior to having your pull
request reviewed, the reviewer(s) may ask you to complete additional
design work, tests, or other changes before your pull request can
be ultimately accepted.

A Pull Request will be rejected without a related issue.

## Styleguides

### Git Commit Messages

* Use the present tense ("Add feature" not "Added feature")
* Use the imperative mood ("Set the header..." not "Sets the header...")
* Limit the first line to 72 characters or less
* Reference issues and pull requests liberally after the first line

### Documentation Styleguide

For Go `package`, `function`, `struct`, `const`, `var` or `interface`
documentation, follow the [Go guidelines for documentation][godoc]. Some of them
are tested with linter.

Any other documentation lives under the `docs/` directory and
uses [Markdown][markdown] format.

## Additional Notes

### Issue and Pull Request Labels

This section lists the labels we use to help track and manage issues and pull requests.

The labels are loosely grouped by their purpose, but it's not required that every issue
have a label from every group or that an issue can't have more than one label from the same group.

### Type of Issue and Issue State

| Label | Issues | Description |
|-------|--------|-------------|
| 🐛 Bug                | [Link][label-bug] | Something isn't working |
| ✨ Feature            | [Link][label-feature] | New feature or request |
| ♻️ Duplicate          | [Link][label-duplicate] | This issue or pull request already exists |
| ❤️ Good first issue   | [Link][label-good-first-issue] | New feature request or enhancement suggestion |
| 🏳️ Help wanted        | [Link][label-help-wanted] | Extra attention is needed |
| 🚫 Invalid            | [Link][label-invalid] | This doesn't seem right |
| 💬 Question           | [Link][label-question] | Further information is requested |
| 🤡 Wontfix            | [Link][label-wontfix] | The team has decided not to fix these issues for now |

### Topic Categories

| Label | Issues | Description |
|-------|--------|-------------|
| 👮 Admin          | [Link][label-admin]           |  Admin endpoint related |
| ⚙️ Core           | [Link][label-core]            | Core related |
| 📖 Documentation  | [Link][label-documentation]   | Improvements or additions to documentation |
| 🔒 Security       | [Link][label-security]        | Security issues |
| 🛠️ Service        | [Link][label-service]         | Related to any of the services/endpoints |

[code-of-conduct]: https://github.com/yitsushi/go-misskey/blob/main/CODE_OF_CONDUCT.md
[issue-search]: https://github.com/search?q=is%3Aissue+repo%3Ayitsushi%2Fgo-misskey
[good-first-issue]: https://github.com/yitsushi/go-misskey/labels/%E2%9D%A4%EF%B8%8F%20Good%20first%20issue
[help-wanted]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%8F%B3%EF%B8%8F%20Help%20wanted
[documentation]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%93%96%20Documentation
[pr-template]: https://github.com/yitsushi/go-misskey/blob/main/.github/PULL_REQUEST_TEMPLATE.md
[patterns-and-structure]: https://github.com/yitsushi/go-misskey/blob/main/docs/patterns-and-structure.md
[godoc]: https://blog.golang.org/godoc
[markdown]: https://daringfireball.net/projects/markdown/
[vscode]: https://code.visualstudio.com/
[vim]: https://www.vim.org/
[vim-rest-console]: https://github.com/diepm/vim-rest-console
[humao-rest-client]: https://marketplace.visualstudio.com/items?itemName=humao.rest-client

[label-bug]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%90%9B%20Bug
[label-feature]: https://github.com/yitsushi/go-misskey/labels/%E2%9C%A8%20Feature
[label-duplicate]: https://github.com/yitsushi/go-misskey/labels/%E2%99%BB%EF%B8%8F%20Duplicate
[label-good-first-issue]: https://github.com/yitsushi/go-misskey/labels/%E2%9D%A4%EF%B8%8F%20Good%20first%20issue
[label-help-wanted]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%8F%B3%EF%B8%8F%20Help%20wanted
[label-invalid]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%9A%AB%20Invalid
[label-question]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%92%AC%20Question
[label-wontfix]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%A4%A1%20Wontfix
[label-admin]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%91%AE%20Admin
[label-core]: https://github.com/yitsushi/go-misskey/labels/%E2%9A%99%EF%B8%8F%20Core
[label-documentation]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%93%96%20Documentation
[label-security]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%94%92%20Security
[label-service]: https://github.com/yitsushi/go-misskey/labels/%F0%9F%9B%A0%EF%B8%8F%20Service
