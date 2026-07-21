// Command new-repo scaffolds a throwaway Go project in one step.
//
// Running `new-repo <name>` creates a directory named <name>, runs
// `go mod init <name>` inside it, drops in a minimal main.go, and opens
// the directory in your editor (Zed by default) so you can start hacking
// immediately.
package main

import (
	"fmt"
	"os"
	"os/exec"
)

const mainTemplate = `package main

import "fmt"

func main() {
	fmt.Println("hello from %s")
}
`

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "new-repo:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 1 || args[0] == "-h" || args[0] == "--help" {
		return fmt.Errorf("usage: new-repo <name>")
	}
	name := args[0]

	// Bail early if the directory already exists so we never clobber
	// anything you care about.
	if _, err := os.Stat(name); err == nil {
		return fmt.Errorf("%q already exists", name)
	} else if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(name, 0o755); err != nil {
		return err
	}

	if err := runIn(name, "go", "mod", "init", name); err != nil {
		return fmt.Errorf("go mod init: %w", err)
	}

	mainGo := name + "/main.go"
	if err := os.WriteFile(mainGo, []byte(fmt.Sprintf(mainTemplate, name)), 0o644); err != nil {
		return err
	}

	fmt.Printf("created %s (module %q)\n", name, name)
	return openEditor(name)
}

// runIn runs a command with dir as its working directory, streaming its
// output to the terminal.
func runIn(dir, cmd string, args ...string) error {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}

// openEditor opens dir in the configured editor. Set NEW_REPO_EDITOR to
// override the default of "zed". A missing editor is a soft failure: the
// project is already scaffolded, so we just say so and move on.
func openEditor(dir string) error {
	editor := os.Getenv("NEW_REPO_EDITOR")
	if editor == "" {
		editor = "zed"
	}

	if _, err := exec.LookPath(editor); err != nil {
		fmt.Printf("(couldn't find %q on PATH — open %s yourself)\n", editor, dir)
		return nil
	}

	c := exec.Command(editor, dir)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
