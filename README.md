# new-repo

A tiny CLI for people who are tired of the "mkdir, `go mod init`, write main.go,
open the editor" dance every time they want to try something small in Go.

```
new-repo <name>
```

That one command:

1. Creates a directory named `<name>`
2. Runs `go mod init <name>` inside it
3. Writes a minimal `main.go`
4. Opens the directory in [Zed](https://zed.dev)

## Install

```sh
go install .          # from a clone of this repo
# or
go install github.com/piechutowski/new-repo@latest
```

Make sure your Go bin directory is on your `PATH` so the command is available
everywhere:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

(Add that line to your shell profile — `~/.zshrc`, `~/.bashrc`, etc.)

## Usage

```sh
new-repo scratch
```

```
created scratch (module "scratch")
```

Zed opens with `scratch/main.go` ready to go:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello from scratch")
}
```

## Using a different editor

Zed is the default. Point `NEW_REPO_EDITOR` at anything else that takes a
directory argument:

```sh
NEW_REPO_EDITOR=code new-repo scratch   # VS Code
NEW_REPO_EDITOR=nvim new-repo scratch   # Neovim
```

If the editor isn't found on your `PATH`, `new-repo` still scaffolds the
project and just tells you to open it yourself.

## License

Apache 2.0 — see [LICENSE](LICENSE).
