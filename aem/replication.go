// © Copyright 2017 XumaK, LLC. All rights reserved.
// Do not distribute.

package aem

import "context"

// ReplicationService handles communication with the Replication related methods of Adobe
// AEM.
type ReplicationService service

type Agent struct {
	Name               string
	Description        string
	Enabled            bool
	SerializationType  string
	RetryDelay         int
	AgentUserID        string
	LogLevel           string
	ReverseReplication bool
	AliasUpdate        bool
	Transport          AgentTransport
	Proxy              AgentProxy
	Extended           AgentExtended
	Triggers           AgentTriggers
	Batch              AgentBatch
}

type AgentTransport struct {
	URI               string
	User              string
	Password          string
	NTLMDomain        string
	NTLMHost          string
	SSL               string
	AllowExpiredCerts bool
}

type AgentProxy struct {
	Host       string
	Port       string
	User       string
	NTLMDomain string
	NTLMHost   string
}

type AgentExtended struct {
	Interface       string
	HTTPMethod      string
	HTTPHeaders     []string
	CloseConnection bool
	ConnectTimeout  int
	SocketTimeout   int
	ProtocolVersion string
}

type AgentTriggers struct {
	IgnoreDefault    bool
	OnModification   bool
	OnDistribute     bool
	OnOffTimeReached bool
	OnReceive        bool
	NoStatusUpdate   bool
	// NoVersioning will not force versioning of activated pages if true.
	NoVersioning bool
}

type AgentBatch struct {
	// EnableBatchMode enables sending replications in batches.
	EnableBatchMode bool
	// MaxWaitTime is the maximum wait time in seconds until a batch request is
	// started.
	MaxWaitTime string
	// TriggerSize starts batch replication when this size limit is reached.
	TriggerSize string
}

func (s *ReplicationService) ActivateTree(ctx context.Context) error {
	// curl -u admin:admin -F cmd=activate -F ignoredeactivated=true -F onlymodified=true -F path=/content/geometrixxhttp://localhost:4502/etc/replication/treeactivation.html

	return _, nil
}

func (s *ReplicationService) Build(ctx context.Context) error {
	return _, nil
}

func (s *ReplicationService) Reinstall(ctx context.Context) error {
	return _, nil
}
