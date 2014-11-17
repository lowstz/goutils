package files

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"
)

func TestAbsPath(t *testing.T) {
	paths := []string{
		"foo",
		"~",
		"~foo",
		"~/foo",
		"foo/~/",
		"$HOME/foo",
	}

	cwd := os.Getenv("PWD")
	u, _ := user.Current()
	home := u.HomeDir

	expanded := []string{
		filepath.Join(cwd, "foo"), // ./foo
		home, // /home/current/
		filepath.Join(cwd, "~foo"),   // ./~foo
		filepath.Join(home, "foo"),   // /home/current/foo
		filepath.Join(cwd, "foo/~/"), // ./foo/~/
		filepath.Join(home, "foo"),   // /home/current/foo
		filepath.Join(home, "foo"),   // /home/current/foo
	}

	for i, path := range paths {
		res := AbsPath(path)
		if res != expanded[i] {
			t.Errorf("%d. Expected '%s' => '%s', got '%s'",
				i, path[i], expanded[i], res)
		}
	}
}
