// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// BundlesService handles communication with the Users related methods of Adobe
// AEM.
type BundlesService service

type Bundle struct {
	SymbolicName     string
	Version          string
	Location         string
	LastModifiction  string
	Documentation    string
	Vendor           string
	Description      string
	StartLevel       int
	ExportedPackages []string
	ImportedPackages []string
	ImportingBundles []string
	ManifestHeaders  []string
}

// Get Bundle config
// curl -u admin:admin http://localhost:4502/system/console/bundles/org.apache.jackrabbit.oak-core.json

func (s *BundlesService) List(ctx context.Context) ([]Bundle, error) {
	return _, nil
}

func (s *BundlesService) Uninstall(ctx context.Context) error {
	// curl -u admin:admin -daction=uninstall http://localhost:4502/system/console/bundles/name-of-bundle

	return _, nil
}

func (s *BundlesService) Stop(ctx context.Context) error {
	// curl -u admin:admin http://localhost:4502/system/console/bundles/org.apache.sling.scripting.jsp -F action=stop
	return _, nil
}

func (s *BundlesService) Start(ctx context.Context) error {
	//	curl -u admin:admin http://localhost:4502/system/console/bundles/org.apache.sling.scripting.jsp -F action=start
	return _, nil
}

func (s *BundlesService) RefreshImport(ctx context.Context) error {
	return _, nil
}

func (s *PackagesService) Update(ctx context.Context) error {
	return _, nil
}
