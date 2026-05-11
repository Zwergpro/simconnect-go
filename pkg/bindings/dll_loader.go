//go:build windows

package bindings

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

// EnvDLLPath is the environment variable that overrides the DLL location.
// If set and the file exists, it is loaded instead of the embedded DLL.
const EnvDLLPath = "SIMCONNECT_DLL"

var (
	dllMu        sync.Mutex
	overridePath string
	loadOnce     sync.Once
	loadErr      error
)

// SetDLLPath sets a process-global override for the SimConnect.dll location.
// If the file at path exists when LoadDLL runs, it is used instead of the
// embedded DLL. If it does not exist, the loader logs a notice and falls back
// to the env-var override or the embedded copy.
//
// Must be called before the first SimConnect call (i.e. before client.Open).
// Calls after the DLL has been loaded are ignored with a warning.
func SetDLLPath(path string) {
	dllMu.Lock()
	defer dllMu.Unlock()
	if loadErr != nil || dllLoaded() {
		log.Printf("simconnect-go: SetDLLPath(%q) ignored — DLL already loaded", path)
		return
	}
	overridePath = path
}

func dllLoaded() bool {
	return dll != nil && dll.Name != ""
}

// LoadDLL resolves the SimConnect.dll path (override → env var → embedded)
// and loads it eagerly so failures surface here rather than as opaque HRESULTs
// later. Safe to call repeatedly; the first call wins.
func LoadDLL() error {
	loadOnce.Do(func() {
		path, err := resolveDLLPath()
		if err != nil {
			loadErr = err
			return
		}
		dll.Name = path
		if err := dll.Load(); err != nil {
			loadErr = fmt.Errorf("simconnect-go: load %s: %w", path, err)
		}
	})
	return loadErr
}

func resolveDLLPath() (string, error) {
	dllMu.Lock()
	override := overridePath
	dllMu.Unlock()

	if override != "" {
		if fileExists(override) {
			return override, nil
		}
		log.Printf("simconnect-go: WithDLLPath(%q) not found, falling back to embedded DLL", override)
	}

	if envPath := os.Getenv(EnvDLLPath); envPath != "" {
		if fileExists(envPath) {
			return envPath, nil
		}
		log.Printf("simconnect-go: %s=%q not found, falling back to embedded DLL", EnvDLLPath, envPath)
	}

	return extractEmbeddedDLL()
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func extractEmbeddedDLL() (string, error) {
	sum := sha256.Sum256(embeddedDLL)
	hash := hex.EncodeToString(sum[:])[:16]
	path := filepath.Join(os.TempDir(), "simconnect-go-"+hash+".dll")

	if info, err := os.Stat(path); err == nil && !info.IsDir() && info.Size() == int64(len(embeddedDLL)) {
		return path, nil
	}

	tmp := path + ".tmp." + strconv.Itoa(os.Getpid())
	if err := os.WriteFile(tmp, embeddedDLL, 0o644); err != nil {
		return "", fmt.Errorf("simconnect-go: write embedded DLL to %s: %w", tmp, err)
	}

	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		if errors.Is(err, fs.ErrExist) {
			return path, nil
		}
		if info, statErr := os.Stat(path); statErr == nil && !info.IsDir() && info.Size() == int64(len(embeddedDLL)) {
			return path, nil
		}
		return "", fmt.Errorf("simconnect-go: rename embedded DLL to %s: %w", path, err)
	}

	return path, nil
}
