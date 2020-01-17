// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// PackagesService handles communication with the Packages related methods of Adobe
// AEM.
type PackagesService service

type Package struct {
	Pid             string        `json:"pid"`
	Path            string        `json:"path"`
	Name            string        `json:"name"`
	DownloadName    string        `json:"downloadName"`
	Group           string        `json:"group"`
	GroupTitle      string        `json:"groupTitle"`
	Version         string        `json:"version"`
	Description     string        `json:"description"`
	Thumbnail       string        `json:"thumbnail"`
	BuildCount      int           `json:"buildCount"`
	Created         int64         `json:"created"`
	CreatedBy       string        `json:"createdBy"`
	LastUnpacked    int64         `json:"lastUnpacked,omitempty"`
	LastUnpackedBy  string        `json:"lastUnpackedBy,omitempty"`
	LastUnwrapped   int64         `json:"lastUnwrapped"`
	Size            int           `json:"size"`
	HasSnapshot     bool          `json:"hasSnapshot"`
	NeedsRewrap     bool          `json:"needsRewrap"`
	RequiresRoot    bool          `json:"requiresRoot"`
	RequiresRestart bool          `json:"requiresRestart"`
	AcHandling      string        `json:"acHandling"`
	Dependencies    []interface{} `json:"dependencies"`
	Resolved        bool          `json:"resolved"`
	Filter          []struct {
		Root  string        `json:"root"`
		Rules []interface{} `json:"rules"`
	} `json:"filter"`
	Screenshots     []interface{} `json:"screenshots"`
	LastUnwrappedBy string        `json:"lastUnwrappedBy,omitempty"`
	BuiltWith       string        `json:"builtWith,omitempty"`
	LastModified    int64         `json:"lastModified,omitempty"`
	LastModifiedBy  string        `json:"lastModifiedBy,omitempty"`
	LastWrapped     int64         `json:"lastWrapped,omitempty"`
	LastWrappedBy   string        `json:"lastWrappedBy,omitempty"`
	TestedWith      string        `json:"testedWith,omitempty"`
	FixedBugs       string        `json:"fixedBugs,omitempty"`
	ProviderName    string        `json:"providerName,omitempty"`
	ProviderURL     string        `json:"providerUrl,omitempty"`
	ProviderLink    string        `json:"providerLink,omitempty"`
}

type APIResponse struct {
	Results []Package `json:"results"`
	Total   int       `json:"total"`
}

func (s *PackagesService) List(ctx context.Context) ([]Bundle, error) {
	// curl -u admin:admin http://localhost:4502/crx/packmgr/list.jsp

	return _, nil
}

func (s *PackagesService) Build(ctx context.Context) error {
	return _, nil
}

func (s *PackagesService) Install(ctx context.Context) error {
	// curl -u admin:admin -X POST http://localhost:4502/crx/packmgr/service/.json/etc/packages/export/name-of-package?cmd=install
	return _, nil
}

func (s *PackagesService) Reinstall(ctx context.Context) error {
	return _, nil
}

func (s *PackagesService) Delete(ctx context.Context) error {
	// curl -u admin:admin -X POST -F cmd=delete http://localhost:4502/crx/packmgr/service/script.html/etc/packages/package-group/package-name.zip
	return _, nil
}

func (p *Package) IfExist() bool {
	// curl -u admin:admin -s -I -w %{http_code} http://localhost:4502/etc/packages/package/group/path/name_of_package.zip

}
