//go:build windows

package bindings

import _ "embed"

//go:embed SimConnect.dll
var embeddedDLL []byte
