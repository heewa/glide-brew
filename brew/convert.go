package brew

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/Masterminds/glide/action"
	"github.com/Masterminds/glide/cfg"
	gpath "github.com/Masterminds/glide/path"

	"github.com/heewa/glide-brew/resource"
)

// LoadLockFile reads the glide lock file, and verifies it's up to date
func LoadLockFile() (*cfg.Lockfile, error) {
	base := "."

	// Ensure GOPATH
	action.EnsureGopath()
	action.EnsureVendorDir()
	conf := action.EnsureConfig()

	// Lockfile exists
	if !gpath.HasLock(base) {
		return nil, errors.New("Lock file (glide.lock) does not exist.")
	}
	// Load lockfile
	lock, err := cfg.ReadLockFile(filepath.Join(base, gpath.LockFile))
	if err != nil {
		return nil, errors.New("Could not load lockfile.")
	}
	// Verify lockfile hasn't changed
	hash, err := conf.Hash()
	if err != nil {
		return nil, errors.New("Could not load lockfile.")
	} else if hash != lock.Hash {
		return nil, errors.New("Lock file may be out of date. Hash check of YAML failed. You may need to run 'update'")
	}

	return lock, nil
}

// ConvertLock converts dependencies to Homebrew resources and prints
// them, using the Lockfile.
func ConvertLock(lock *cfg.Lockfile) ([]resource.Resource, error) {
	resources := make([]resource.Resource, 0, len(lock.Imports))
	for _, lock := range lock.Imports {
		res, err := resource.FromLock(lock)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert a dependency: %v", err)
		}

		resources = append(resources, res)
	}

	return resources, nil
}
