package helpers

import (
	"encoding/base64"
	"fmt"
)

// CopyToClipboardOSC52 prints the OSC52 escape sequence for clipboard integration.
// This is for use in modern terminals that support it (iTerm2, Kitty, WezTerm, etc.),
// allowing clipboard access over SSH.
func CopyToClipboardOSC52(text string) {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	// OSC 52 sequence: \x1b]52;c;BASE64_STRING\x07
	fmt.Printf("\x1b]52;c;%s\x07", encoded)
}