package schema

var (
	swarm = `
	type Swarm {
	 	createdAt: String
	 	id: ID!
		joinTokens: JoinTokens
		rootRotationInProgress: Boolean
		swarmSpec: SwarmSpec
		tlsInfo: TLSInfo
		unlockKey: UnlockKey
		updatedAt: String
		version: SwarmVersion
	}
	
	type SwarmSpec {
	 	caConfig: CaConfig
	 	dispatcher: DispatcherConfig
	 	encryptionConfig: EncryptionConfig
	 	# TBI
	 	# labels: StringAnyMap
	 	name: String
	 	orchestration: OrchestrationConfig
	 	raft: RaftConfig
	 	taskDefaults: TaskDefaults
	}
	
	type OrchestrationConfig {
		taskHistoryRetentionLimit: Int
	}
	
	type RaftConfig {
		electionTick: Int
		heartbeatTick: Int
		keepOldSnapshots: Int
		logEntriesForSlowFollowers: Int
		snapshotInterval: Int
	}
	
	type DispatcherConfig {
	 	heartbeatPeriod: String
	}
	
	type CaConfig {
		externalCas: [ExternalCa]
		nodeCertExpiry: String
	}
	
	type ExternalCa {
		caCert: String
		# options: StringAnyMap
		protocol: String
		url: String
	}
	
	type EncryptionConfig {
		autoLockManagers: Boolean
	}
	
	type TaskDefaults {
	 	logDriver: Driver
	}
	
	type Driver {
		name: String
		# TBI
		# options: StringAnyMap
	}
	
	type JoinTokens {
		manager: String
		worker: String
	}
	
	type UnlockKey {
	 	unlockKey: String
	}
	
	type SwarmInfo {
		cluster: SwarmCluster
		controlAvailable: Boolean
		error: String
		localNodeState: String
		managers: Int
		nodeAddr: String
		nodeId: String
		nodes: Int
		remoteManagers: [RemoteManager]
	}
	
	type SwarmCluster {
		createdAt: String
	 	id: ID!
	 	rootRotationInProgress: Boolean
	 	swarmSpec: SwarmSpec
	 	tlsInfo: TLSInfo
	 	updatedAt: String
	 	version: SwarmVersion
	}
	
	type SwarmVersion {
		index: Int
	}

	type TLSInfo{
		certIssuerPublicKey: String
		certIssuerSubject: String
		trustRoot: String
	}
	
	type RemoteManager {
		addr: String
		nodeId: String
	}
	`
)
