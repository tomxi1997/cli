//go:build !windows
// +build !windows

package manager

var defaultSystemPluginDirs = []string{
	"/data/docker/android/lib/docker/cli-plugins", "/data/docker/android/libexec/docker/cli-plugins",
}
