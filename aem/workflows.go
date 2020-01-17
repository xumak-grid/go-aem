// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// WorkflowsService handles communication with the Workflows related methods of Adobe
// AEM.
type WorkflowsService service

func (s *WorkflowsService) Enable(ctx context.Context) error {
	// curl -u admin:admin -X POST 'http://127.0.0.1:4502/libs/cq/workflow/launcher' --data '_charset_=utf-8&edit=%2Fetc%2Fworkflow%2Flauncher%2Fconfig%2Fupdate_asset_create&%3Astatus=browser&eventType=1&nodetype=nt%3Afile&glob=%2Fcontent%2Fdam(%2F.*%2F)renditions%2Foriginal&condition=&workflow=%2Fetc%2Fworkflow%2Fmodels%2Fdam%2Fupdate_asset%2Fjcr%3Acontent%2Fmodel&description=&enabled=true&excludeList=&runModes=author'

	return _, nil
}

func (s *WorkflowsService) Disable(ctx context.Context) error {
	// curl -u admin:admin -X POST 'http://127.0.0.1:4502/libs/cq/workflow/launcher' --data '_charset_=utf-8&edit=%2Fetc%2Fworkflow%2Flauncher%2Fconfig%2Fupdate_asset_create&%3Astatus=browser&eventType=1&nodetype=nt%3Afile&glob=%2Fcontent%2Fdam(%2F.*%2F)renditions%2Foriginal&condition=&workflow=%2Fetc%2Fworkflow%2Fmodels%2Fdam%2Fupdate_asset%2Fjcr%3Acontent%2Fmodel&description=&enabled=false&excludeList=&runModes=author'

	return _, nil
}
