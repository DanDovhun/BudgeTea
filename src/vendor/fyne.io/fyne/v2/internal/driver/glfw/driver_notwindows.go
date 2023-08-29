//go:build !windows
// +build !windows

package glfw

import "fyne.io/fyne/v2"

func logError(msg string, err error) {
	fyne.LogError(msg, err)
}
