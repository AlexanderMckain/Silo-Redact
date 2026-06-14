//go:build !windows

package main

import "github.com/atotto/clipboard"

func ghostPollReadClipboard() (string, bool) {
	s, err := clipboard.ReadAll()
	if err != nil || s == "" {
		return "", false
	}
	return s, true
}

func ghostPollWriteClipboard(text string) bool {
	return clipboard.WriteAll(text) == nil
}
