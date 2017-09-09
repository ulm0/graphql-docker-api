package schema

import "github.com/graphql-go/graphql"

var (
	swInfo = sysInfo.Swarm
	// Swarm related custom object types
	caConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "CaConfig",
		Fields: graphql.Fields{
			"externalCas": &graphql.Field{
				Type: graphql.NewList(externalCa),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.CAConfig.ExternalCAs, nil
				},
			},
			"nodeCertExpiry": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.CAConfig.NodeCertExpiry), nil
				},
			},
		},
	})
	dispatcherConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "DispatcherConfig",
		Fields: graphql.Fields{
			"heartbeatPeriod": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.Dispatcher.HeartbeatPeriod), nil
				},
			},
		},
	})
	driver = graphql.NewObject(graphql.ObjectConfig{
		Name: "Driver",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.TaskDefaults.LogDriver.Name, nil
				},
			},
			// TBI
			// "options": &graphql.Field{
			// 	Type: stringAnyMap,
			// 	// Resolver
			// },
		},
	})
	encryptionConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "EncryptionConfig",
		Fields: graphql.Fields{
			"autoLockManagers": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.EncryptionConfig.AutoLockManagers, nil
				},
			},
		},
	})
	engineConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "EngineConfig",
		Fields: graphql.Fields{
			"engineVersion": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"plugins": &graphql.Field{
				Type: enginePlugin,
				// This must return a enginePlugin slice
				// Resolver
			},
		},
	})
	enginePlugin = graphql.NewObject(graphql.ObjectConfig{
		Name: "EnginePlugin",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"type": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	externalCa = graphql.NewObject(graphql.ObjectConfig{
		Name: "ExternalCa",
		Fields: graphql.Fields{
			"caCert": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					for _, exCa := range swInfo.Cluster.Spec.CAConfig.ExternalCAs {
						return exCa.CACert, nil
					}
					return nil, nil
				},
			},
			// TBI
			// "options": &graphql.Field{
			// 	Type: stringAnyMap,
			// 	// Resolver
			// },
			"protocols": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					for _, exCa := range swInfo.Cluster.Spec.CAConfig.ExternalCAs {
						return exCa.Protocol, nil
					}
					return nil, nil
				},
			},
			"url": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					for _, exCa := range swInfo.Cluster.Spec.CAConfig.ExternalCAs {
						return exCa.URL, nil
					}
					return nil, nil
				},
			},
		},
	})
	joinTokens = graphql.NewObject(graphql.ObjectConfig{
		Name: "JoinTokens",
		Fields: graphql.Fields{
			"manager": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"worker": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	managerStatus = graphql.NewObject(graphql.ObjectConfig{
		Name: "ManagerStatus",
		Fields: graphql.Fields{
			"addr": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"leader": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"reachability": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	nodeDescription = graphql.NewObject(graphql.ObjectConfig{
		Name: "NodeDescription",
		Fields: graphql.Fields{
			"engine": &graphql.Field{
				Type: engineConfig,
				// Resolver
			},
			"hostname": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"platform": &graphql.Field{
				Type: platform,
				// Resolver
			},
			"resources": &graphql.Field{
				Type: resources,
				// Resolver
			},
		},
	})
	nodeInfo = graphql.NewObject(graphql.ObjectConfig{
		Name: "NodeInfo",
		Fields: graphql.Fields{
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"description": &graphql.Field{
				Type: nodeDescription,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"managerStatus": &graphql.Field{
				Type: managerStatus,
				// Resolver
			},
			"spec": &graphql.Field{
				Type: nodeSpec,
				// Resolver
			},
			"status": &graphql.Field{
				Type: nodeStatus,
				// Resolver
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"version": &graphql.Field{
				Type: swarmVersion,
				// Resolver
			},
		},
	})
	nodeMembership = graphql.NewEnum(graphql.EnumConfig{
		Name: "NodeMembership",
		Values: graphql.EnumValueConfigMap{
			"accepted": &graphql.EnumValueConfig{
				Value: "accepted",
			},
			"pending": &graphql.EnumValueConfig{
				Value: "pending",
			},
		},
	})
	nodeRole = graphql.NewEnum(graphql.EnumConfig{
		Name: "NodeRole",
		Values: graphql.EnumValueConfigMap{
			"manager": &graphql.EnumValueConfig{
				Value: "manager",
			},
			"worker": &graphql.EnumValueConfig{
				Value: "worker",
			},
		},
	})
	nodeSpec = graphql.NewObject(graphql.ObjectConfig{
		Name: "NodeSpec",
		Fields: graphql.Fields{
			"availability": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"role": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	nodeStatus = graphql.NewObject(graphql.ObjectConfig{
		Name: "NodeStatus",
		Fields: graphql.Fields{
			"addr": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"state": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	nodefilter = graphql.NewObject(graphql.ObjectConfig{
		Name: "NodeFilter",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
				// This must return a string slice
				// Resolver
			},
			"label": &graphql.Field{
				Type: graphql.String,
				// This must return a string slice
				// Resolver
			},
			"membership": &graphql.Field{
				Type: nodeMembership,
				// This must return a string slice
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// This must return a string slice
				// Resolver
			},
			"role": &graphql.Field{
				Type: nodeRole,
				// This must return a string slice
				// Resolver
			},
		},
	})
	orchestrationConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "OrchestrationConfig",
		Fields: graphql.Fields{
			"taskHistoryRetentionLimit": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(*swInfo.Cluster.Spec.Orchestration.TaskHistoryRetentionLimit), nil
				},
			},
		},
	})
	platform = graphql.NewObject(graphql.ObjectConfig{
		Name: "Platform",
		Fields: graphql.Fields{
			"architecture": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"os": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	raftconfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "RaftConfig",
		Fields: graphql.Fields{
			"electionTick": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.Raft.ElectionTick), nil
				},
			},
			"heartbeatTick": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.Raft.HeartbeatTick), nil
				},
			},
			"keepOldSnapshots": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(*swInfo.Cluster.Spec.Raft.KeepOldSnapshots), nil
				},
			},
			"logEntriesForSlowFollowers": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.Raft.LogEntriesForSlowFollowers), nil
				},
			},
			"snapshotInterval": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Spec.Raft.SnapshotInterval), nil
				},
			},
		},
	})
	remoteManager = graphql.NewObject(graphql.ObjectConfig{
		Name: "RemoteManager",
		Fields: graphql.Fields{
			"addr": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					for _, peer := range swInfo.RemoteManagers {
						return peer.Addr, nil
					}
					return nil, nil
				},
			},
			"nodeId": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					for _, peer := range swInfo.RemoteManagers {
						return peer.NodeID, nil
					}
					return nil, nil
				},
			},
		},
	})
	resources = graphql.NewObject(graphql.ObjectConfig{
		Name: "Resources",
		Fields: graphql.Fields{
			"memoryBytes": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"nanoCpus": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	swarm = graphql.NewObject(graphql.ObjectConfig{
		Name: "Swarm",
		Fields: graphql.Fields{
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"joinTokens": &graphql.Field{
				Type: joinTokens,
				// Resolver
			},
			"node": &graphql.Field{
				Type: swarmNode,
				// It receives (args: filter nodeFilter)
				// This must return a swarmNode slice
				// Resolver
			},
			"swarmSpec": &graphql.Field{
				Type: swarmSpec,
				// Resolver
			},
			"unlockKey": &graphql.Field{
				Type: unlockKey,
				// Resolver
			},
			"version": &graphql.Field{
				Type: swarmVersion,
				// Resolver
			},
		},
	})
	swarmCluster = graphql.NewObject(graphql.ObjectConfig{
		Name: "SwarmCluster",
		Fields: graphql.Fields{
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.CreatedAt, nil
				},
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.ID, nil
				},
			},
			"rootRotationInProgress": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.RootRotationInProgress, nil
				},
			},
			"swarmSpec": &graphql.Field{
				Type: swarmSpec,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec, nil
				},
			},
			"tlsInfo": &graphql.Field{
				Type: tlsInfo,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.TLSInfo, nil
				},
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.UpdatedAt, nil
				},
			},
			"version": &graphql.Field{
				Type: swarmVersion,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Version, nil
				},
			},
		},
	})
	swarmInfo = graphql.NewObject(graphql.ObjectConfig{
		Name: "SwarmInfo",
		Fields: graphql.Fields{
			"cluster": &graphql.Field{
				Type: swarmCluster,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster, nil
				},
			},
			"controlAvailable": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.ControlAvailable, nil
				},
			},
			"error": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Error, nil
				},
			},
			"localNodeState": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.LocalNodeState, nil
				},
			},
			"managers": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Managers, nil
				},
			},
			"nodeAddr": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.NodeAddr, nil
				},
			},
			"nodeId": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.NodeID, nil
				},
			},
			"nodes": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Nodes, nil
				},
			},
			"remoteManagers": &graphql.Field{
				Type: graphql.NewList(remoteManager),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.RemoteManagers, nil
				},
			},
		},
	})
	swarmNode = graphql.NewObject(graphql.ObjectConfig{
		Name: "SwarmNode",
		Fields: graphql.Fields{
			"createdAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"description": &graphql.Field{
				Type: nodeDescription,
				// Resolver
			},
			"details": &graphql.Field{
				Type: nodeInfo,
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"managerStatus": &graphql.Field{
				Type: managerStatus,
				// Resolver
			},
			"spec": &graphql.Field{
				Type: nodeSpec,
				// Resolver
			},
			"status": &graphql.Field{
				Type: nodeStatus,
				// Resolver
			},
			"updatedAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"version": &graphql.Field{
				Type: swarmVersion,
				// Resolver
			},
		},
	})
	swarmSpec = graphql.NewObject(graphql.ObjectConfig{
		Name: "SwarmSpec",
		Fields: graphql.Fields{
			"caConfig": &graphql.Field{
				Type: caConfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.CAConfig, nil
				},
			},
			"dispatcher": &graphql.Field{
				Type: dispatcherConfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.Dispatcher, nil
				},
			},
			"encryptionConfig": &graphql.Field{
				Type: encryptionConfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.EncryptionConfig, nil
				},
			},
			// TBI
			// "labels": &graphql.Field{
			// 	Type: stringAnyMap,
			// 	// Resolver
			// },
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.Annotations.Name, nil
				},
			},
			"orchestration": &graphql.Field{
				Type: orchestrationConfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.Orchestration, nil
				},
			},
			"raft": &graphql.Field{
				Type: raftconfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.Raft, nil
				},
			},
			"taskDefaults": &graphql.Field{
				Type: taskDefaults,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.TaskDefaults, nil
				},
			},
		},
	})
	swarmVersion = graphql.NewObject(graphql.ObjectConfig{
		Name: "SwarmVersion",
		Fields: graphql.Fields{
			"index": &graphql.Field{
				Type: graphql.Float,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(swInfo.Cluster.Version.Index), nil
				},
			},
		},
	})
	taskDefaults = graphql.NewObject(graphql.ObjectConfig{
		Name: "TaskDefault",
		Fields: graphql.Fields{
			"logDriver": &graphql.Field{
				Type: driver,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.Spec.TaskDefaults.LogDriver, nil
				},
			},
		},
	})
	tlsInfo = graphql.NewObject(graphql.ObjectConfig{
		Name: "tlsInfo",
		Fields: graphql.Fields{
			"certIssuerPublicKey": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return string(swInfo.Cluster.TLSInfo.CertIssuerPublicKey), nil
				},
			},
			"certIssuerSubject": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return string(swInfo.Cluster.TLSInfo.CertIssuerSubject), nil
				},
			},
			"trustRoot": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return swInfo.Cluster.TLSInfo.TrustRoot, nil
				},
			},
		},
	})
	unlockKey = graphql.NewObject(graphql.ObjectConfig{
		Name: "UnlockKey",
		Fields: graphql.Fields{
			"unlockKey": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
)
