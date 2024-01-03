package platform

import (
	"golang.org/x/sys/windows"
	"os"
)

// GetSurgeDir returns the surge dir
func GetAppDir() string {
	return os.Getenv("APPDATA") + string(os.PathSeparator) + "AnyNoteX"
}

// GetRemoteFolder returns the download dir
func GetRemoteFolder() (string, error) {
	homedir, _ := windows.KnownFolderPath(windows.FOLDERID_Downloads, 0)
	return homedir + string(os.PathSeparator) + "surge_downloads", nil
}
