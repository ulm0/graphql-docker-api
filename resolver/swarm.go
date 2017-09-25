package resolver

import graphql "github.com/neelance/graphql-go"

/* Swarm type resolver */

func (r *swarmResolver) CreatedAt() *string { return ptrString(r.Swarm.CreatedAt.String()) }
func (r *swarmResolver) Id() graphql.ID     { return graphql.ID(r.Swarm.ID) }
func (r *swarmResolver) JoinTokens() *jointTokensResolver {
	tokens := r.Swarm.JoinTokens
	return &jointTokensResolver{JoinTokens: tokens}
}
func (r *swarmResolver) RootRotationInProgress() *bool {
	return &r.Swarm.RootRotationInProgress
}
func (r *swarmResolver) SwarmSpec() *swarmSpecResolver {
	spec := r.Swarm.Spec
	return &swarmSpecResolver{SwarmSpec: spec}
}
func (r *swarmResolver) TLSInfo() *tlsInfoResolver {
	tls := r.Swarm.TLSInfo
	return &tlsInfoResolver{TLSInfo: tls}
}
func (r *swarmResolver) UnlockKey() *unlockKeyResolver {
	cli := clientInit()
	k, _ := cli.SwarmGetUnlockKey(ctx)
	key := unlockKeyResolver(k)
	return &key
}
func (r *swarmResolver) UpdatedAt() *string { return ptrString(r.Swarm.CreatedAt.String()) }
func (r *swarmResolver) Version() *swarmVersionResolver {
	ver := swarmVersionResolver(r.Swarm.Version)
	return &ver
}

/*	SwarmSpec type resolver */

func (r *swarmSpecResolver) CaConfig() *caConfigResolver {
	ca := r.SwarmSpec.CAConfig
	return &caConfigResolver{CaConfig: ca}
}
func (r *swarmSpecResolver) Dispatcher() *dispatcherConfigResolver {
	dis := r.SwarmSpec.Dispatcher
	return &dispatcherConfigResolver{DispatcherConfig: dis}
}
func (r *swarmSpecResolver) EncryptionConfig() *encryptionConfigResolver {
	enc := encryptionConfigResolver(r.SwarmSpec.EncryptionConfig)
	return &enc
}
func (r *swarmSpecResolver) Name() *string { return &r.SwarmSpec.Name }
func (r *swarmSpecResolver) Orchestration() *orchestrationConfigResolver {
	orc := orchestrationConfigResolver(r.SwarmSpec.Orchestration)
	return &orc
}
func (r *swarmSpecResolver) Raft() *raftConfigResolver {
	raft := r.SwarmSpec.Raft
	return &raftConfigResolver{RaftConfig: raft}
}
func (r *swarmSpecResolver) TaskDefaults() *taskDefaultsResolver {
	task := r.SwarmSpec.TaskDefaults
	return &taskDefaultsResolver{TaskDefault: task}
}

/*	OrchestrationConfig type resolver */

func (r *orchestrationConfigResolver) TASKHistoryRetentionLimit() *int32 {
	return ptrInt32(r.TaskHistoryRetentionLimit)
}

/*	RaftConfig type resolver */

func (r *raftConfigResolver) ElectionTick() *int32     { return ptrInt32(r.RaftConfig.ElectionTick) }
func (r *raftConfigResolver) HeartbeatTick() *int32    { return ptrInt32(r.RaftConfig.HeartbeatTick) }
func (r *raftConfigResolver) KeepOldSnapshots() *int32 { return ptrInt32(r.RaftConfig.KeepOldSnapshots) }
func (r *raftConfigResolver) LogEntriesForSlowFollowers() *int32 {
	return ptrInt32(r.RaftConfig.LogEntriesForSlowFollowers)
}
func (r *raftConfigResolver) SnapshotInterval() *int32 { return ptrInt32(r.RaftConfig.SnapshotInterval) }

/*	DispatcherConfig type resolver */

func (r *dispatcherConfigResolver) HeartbeatPeriod() *string {
	return ptrString(r.DispatcherConfig.HeartbeatPeriod.String())
}

/*	CaConfig Resolver */

func (r *caConfigResolver) ExternalCas() *[]*externalCaResolver {
	cas := make([]*externalCaResolver, 0)
	for _, v := range r.CaConfig.ExternalCAs {
		cas = append(cas, &externalCaResolver{ExternalCa: v})
	}
	return &cas
}
func (r *caConfigResolver) NodeCertExpiry() *string {
	return ptrString(r.CaConfig.NodeCertExpiry.String())
}

/*	ExternalCa type resolver */

func (r *externalCaResolver) CaCert() *string   { return &r.ExternalCa.CACert }
func (r *externalCaResolver) Protocol() *string { return ptrString(r.ExternalCa.Protocol) }
func (r *externalCaResolver) Url() *string      { return &r.ExternalCa.URL }

/* EncryptionConfig type resolver */

func (r *encryptionConfigResolver) AutoLOCKManagers() *bool { return &r.AutoLockManagers }

/* TaskDefault type resolver */

func (r *taskDefaultsResolver) LogDriver() *swarmDriverResolver {
	driver := r.TaskDefault.LogDriver
	return &swarmDriverResolver{LogDriver: driver}
}

/* SwarmDriver type resolver */

// NOTE: This throws "graphql: panic occurred: runtime error: invalid memory address or nil pointer dereference"
func (r *swarmDriverResolver) NAME() *string { return &r.LogDriver.Name }

/* Jointtokens type resolver */

func (r *jointTokensResolver) Manager() *string { return &r.JoinTokens.Manager }
func (r *jointTokensResolver) Worker() *string  { return &r.JoinTokens.Worker }

/* UnlockKey type resolver */

func (r *unlockKeyResolver) UnlockKEY() *string { return ptrString(r.UnlockKey) }

/*	SwarmInfo type resolver */

func (r *swarmInfoResolver) Cluster() *swarmClusterResolver {
	cluster := r.Swarm.Cluster
	return &swarmClusterResolver{SwarmCluster: cluster}
}
func (r *swarmInfoResolver) ControlAvailable() *bool { return &r.Swarm.ControlAvailable }
func (r *swarmInfoResolver) Error() *string          { return &r.Swarm.Error }
func (r *swarmInfoResolver) LocalNodeState() *string { return ptrString(r.Swarm.LocalNodeState) }
func (r *swarmInfoResolver) Managers() *int32        { return ptrInt32(r.Swarm.Managers) }
func (r *swarmInfoResolver) NodeAddr() *string       { return &r.Swarm.NodeAddr }
func (r *swarmInfoResolver) NodeID() *string         { return &r.Swarm.NodeID }
func (r *swarmInfoResolver) Nodes() *int32           { return ptrInt32(r.Swarm.Nodes) }
func (r *swarmInfoResolver) RemoteManagers() *[]*remoteManagerResolver {
	peers := make([]*remoteManagerResolver, 0)
	for _, v := range r.Swarm.RemoteManagers {
		peers = append(peers, &remoteManagerResolver{RemoteManager: v})
	}
	return &peers
}

/*	SwarmCluster type resolver */

func (r *swarmClusterResolver) CreatedAt() *string {
	return ptrString(r.SwarmCluster.CreatedAt.String())
}
func (r *swarmClusterResolver) ID() graphql.ID { return graphql.ID(r.SwarmCluster.ID) }
func (r *swarmClusterResolver) RootRotationInProgress() *bool {
	return &r.SwarmCluster.RootRotationInProgress
}
func (r *swarmClusterResolver) SwarmSpec() *swarmSpecResolver {
	spec := r.SwarmCluster.Spec
	return &swarmSpecResolver{SwarmSpec: spec}
}
func (r *swarmClusterResolver) TLSInfo() *tlsInfoResolver {
	tls := r.SwarmCluster.TLSInfo
	return &tlsInfoResolver{TLSInfo: tls}
}
func (r *swarmClusterResolver) UpdatedAt() *string {
	return ptrString(r.SwarmCluster.UpdatedAt.String())
}
func (r *swarmClusterResolver) Version() *swarmVersionResolver {
	return ptrSwVer(r.SwarmCluster.Version)
}

/*	SwarmVersion type resolver */

func (r *swarmVersionResolver) INDEX() *int32 { return ptrInt32(r.Index) }

/*	TLSInfo type resolver */

func (r *tlsInfoResolver) CertIssuerPublicKey() *string {
	return ptrString(r.TLSInfo.CertIssuerPublicKey)
}
func (r *tlsInfoResolver) CertIssuerSubject() *string { return ptrString(r.TLSInfo.CertIssuerSubject) }
func (r *tlsInfoResolver) TrustRoot() *string         { return &r.TLSInfo.TrustRoot }

/*	RemoteManager type resolver */

func (r *remoteManagerResolver) Addr() *string   { return &r.RemoteManager.Addr }
func (r *remoteManagerResolver) NodeID() *string { return &r.RemoteManager.NodeID }
