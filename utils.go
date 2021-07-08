package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// isValidURL tests a string to determine if it is a well-structured url or not.
func isValidURL(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func Assert(t *testing.T, got interface{}, expected interface{}) {
	assert.Equal(t, expected, got)
}

// expandHomeDir expands the "~/" part of a path to the current user's home directory
func expandHomeDir(p string) string {
	usr, _ := user.Current()
	homedir := usr.HomeDir
	if p == "~" {
		// In case of "~", which won't be caught by the "else if"
		p = homedir
	} else if strings.HasPrefix(p, "~/") {
		// Use strings.HasPrefix so we don't match paths like
		// "/something/~/something/"
		p = filepath.Join(homedir, p[2:])
	}
	return p
}

// GetConfigDir returns the absolute path of ffcss's configuration directory
func GetConfigDir() string {
	return expandHomeDir("~/.config/ffcss")
}

// GetCacheDir returns the temporary path for cloned repos and other stuff
func GetCacheDir() string {
	return expandHomeDir("~/.cache/ffcss/")
}

// CacheDir joins the cache directory with the given path segments
func CacheDir(pathSegments ...string) string {
	return filepath.Join(GetCacheDir(), filepath.Join(pathSegments...))
}

// ConfigDir joins the config directory with the given path segments
func ConfigDir(pathSegments ...string) string {
	return filepath.Join(GetConfigDir(), filepath.Join(pathSegments...))
}

// GetManifestPath returns the path of a theme's manifest file
func GetManifestPath(themeRoot string) string {
	return filepath.Join(themeRoot, "ffcss.yaml")
}

// ProfileDirsPaths returns an array of profile directories from the profile folder.
// 1 arguments: the profiles folder is assumed to be the current OS's default.
// 2 argument: use the given profiles folder
// more arguments: panic.
func ProfileDirsPaths(operatingSystem string, optionalProfilesDir ...string) ([]string, error) {
	var profilesFolder string
	if len(optionalProfilesDir) == 0 {
		// XXX: Weird golang thing, if I assign to profilesFolder directly, it tells me the variable is unused
		_profilesFolder, err := DefaultProfilesDir(operatingSystem)
		profilesFolder = _profilesFolder
		if err != nil {
			return []string{}, fmt.Errorf("couldn't get the profiles folder: %w. Try to use --profiles-dir", err)
		}
	} else if len(optionalProfilesDir) == 1 {
		profilesFolder = optionalProfilesDir[0]
	} else {
		panic(fmt.Sprintf("received %d arguments, expected 1 or 2", len(optionalProfilesDir)+1))
	}
	directories, err := os.ReadDir(profilesFolder)
	releasesPaths := make([]string, 0)
	patternReleaseID := regexp.MustCompile(`[a-z0-9]{8}\.\w+`)
	if err != nil {
		return []string{}, fmt.Errorf("couldn't read %s: %w", profilesFolder, err)
	}
	for _, releasePath := range directories {
		if patternReleaseID.MatchString(releasePath.Name()) {
			stat, err := os.Stat(filepath.Join(profilesFolder, releasePath.Name()))
			if err != nil {
				continue
			}
			if stat.IsDir() {
				releasesPaths = append(releasesPaths, filepath.Join(profilesFolder, releasePath.Name()))
			}
		}
	}
	return releasesPaths, nil
}

func DefaultProfilesDir(operatingSystem string) (string, error) {
	switch operatingSystem {
	case "linux":
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homedir, ".mozilla", "firefox"), nil
	case "macos":
		user, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("couldn't get the current user: %w", err)
		}

		return filepath.Join("/Users", user.Username, "Library", "Application Support", "Firefox", "Profiles"), nil
	case "windows":
		homedir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		return filepath.Join(homedir, "AppData", "Roaming", "Mozilla", "Firefox", "Profiles"), nil
	}
	return "", fmt.Errorf("unknown operating system %s", operatingSystem)
}

func cwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}

// isURLClonable determines if the given URL points to a git repository
func isURLClonable(URL string) (clonable bool, err error) {
	output, err := exec.Command("git", "ls-remote", URL).CombinedOutput()
	if err == nil {
		return true, nil
	}
	switch err.(type) {
	case *exec.ExitError:
		if err.(*exec.ExitError).ExitCode() == 128 {
			return false, nil
		}
	}
	return false, fmt.Errorf("while running git-ls-remote: %w: %s", err, output)
}

// RenameIfExists renames from to to if from exists. If it doesn't, don't attempt renaming.
func RenameIfExists(from string, to string) error {
	if _, err := os.Stat(from); os.IsNotExist(err) {
		return nil
	}
	if _, err := os.Stat(to); os.IsNotExist(err) {
		return os.Rename(from, to)
	}
	err := os.RemoveAll(to)
	if err != nil {
		return err
	}
	return os.Rename(from, to)
}
