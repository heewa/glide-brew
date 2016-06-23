package resource

import (
	"fmt"

	"github.com/Masterminds/glide/cfg"
)

// Resource represents a Homebrew resource definition for a Go dependency.
// See: http://www.rubydoc.info/github/Homebrew/homebrew/master/Resource/Go
type Resource struct {
	Name             string
	URL              string
	Revision         string
	DownloadStrategy string
}

// FromLock converts a Glide Lock to a Resource
func FromLock(lock *cfg.Lock) (Resource, error) {
	// Get repo info about the locked dependency and convert to homebrew's
	// resource attributes
	dep := cfg.DependencyFromLock(lock)

	repo, err := dep.GetRepo("")
	if err != nil {
		return Resource{}, err
	}

	br := Resource{
		Name:             dep.Name,
		URL:              repo.Remote(),
		Revision:         lock.Version,
		DownloadStrategy: string(repo.Vcs()),
	}

	return br, nil
}

// String serializes a Resource into Homebrew's syntax, for inclusion in a
// formula.
func (res Resource) String() string {
	return fmt.Sprintf(`go_resource "%s" do
  url "%s", :using => :%s, :revision => "%s"
end`, res.Name, res.URL, res.DownloadStrategy, res.Revision)
}
