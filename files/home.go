package files

import (
	"os"
	"runtime"
)

var Path string

func ReadyDir(port string) {
	home := UserHomeDir()
	Path = home + "/wolf" + port
	os.Mkdir(Path, 0755)
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
