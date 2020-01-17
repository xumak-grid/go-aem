// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// DatastoreService handles communication with datastore related methods of Adobe
// AEM.
type DatastoreService service

func (s *WorkflowsService) GarbageCollect(ctx context.Context) error {
	// curl -u admin:admin -X POST http://localhost:4502/system/console/jmx/com.adobe.granite:type=Repository/op/runDataStoreGarbageCollection/java.lang.Boolean

	return _, nil
}
