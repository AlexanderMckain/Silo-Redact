package main

import (
	"fmt"
	"os"
	"strings"
)

// siloRedactDebug is true when SILOREDACT_DEBUG=1 (developer diagnostics only; never log secrets in release).
var siloRedactDebug = strings.TrimSpace(os.Getenv("SILOREDACT_DEBUG")) == "1"

func debugLog(format string, args ...interface{}) {
	if !siloRedactDebug {
		return
	}
	fmt.Printf(format, args...)
}

func debugLogln(msg string) {
	if !siloRedactDebug {
		return
	}
	fmt.Println(msg)
}
