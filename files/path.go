package files

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// return file path's absolute path
func AbsPath(path string) string {
	if path[0] == '~' {
		sep := strings.Index(path, string(os.PathSeparator))
		if sep < 0 {
			sep = len(path)
		}
		var err error
		var u *user.User
		username := path[1:sep]
		if len(username) == 0 {
			u, err = user.Current()
		} else {
			u, err = user.Lookup(username)
		}
		if err == nil {
			path = filepath.Join(u.HomeDir, path[sep:])
		}
	}
	path = os.ExpandEnv(path)
	abs, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abs
}
