package schema

import "github.com/graphql-go/graphql"

var (
	swInfo = sysInfo.Swarm
	// Swarm related custom object types
	caConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "CaConfig",
		Fields: graphql.Fields{
			"externalCas": &graphql.Field{
				Type: externalCa,
				// This must return a externalCa slice
				// Resolver
			},
			"nodeCertExpiry": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	dispatcherConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "DispatcherConfig",
		Fields: graphql.Fields{
			"heartbeatPeriod": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	driver = graphql.NewObject(graphql.ObjectConfig{
		Name: "Driver",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"options": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
		},
	})
	encryptionConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "EncryptionConfig",
		Fields: graphql.Fields{
			"autoLockManagers": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
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
			"options": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"protocols": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"url": &graphql.Field{
				Type: graphql.String,
				// Resolver
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
				// Resolver
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
				// Resolver
			},
			"heartbeatTick": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"keepOldSnapshots": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"logEntriesForSlowFollowers": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"snapshotInterval": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
		},
	})
	// TBI
	// remoteManager = graphql.NewObject(graphql.ObjectConfig{
	// 	Name: "RemoteManager",
	// 	Fields: graphql.Fields{
	// 		"addr": &graphql.Field{
	// 			Type: graphql.String,
	// 			// Resolver
	// 		},
	// 		"nodeId": &graphql.Field{
	// 			Type: graphql.String,
	// 			// Resolver
	// 		},
	// 	},
	// })
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
			// TBI
			// "swarmSpec": &graphql.Field{
			// 	Type: swarmSpec,
			// 	// Resolver
			// },
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
			// TBI
			// "remoteManagers": &graphql.Field{
			// 	Type: graphql.NewList(remoteManager),
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return swInfo.RemoteManagers, nil
			// 	},
			// },
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
				// Resolver
			},
			"dispatcher": &graphql.Field{
				Type: dispatcherConfig,
				// Resolver
			},
			"encryptionConfig": &graphql.Field{
				Type: encryptionConfig,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"orchestration": &graphql.Field{
				Type: orchestrationConfig,
				// Resolver
			},
			"raft": &graphql.Field{
				Type: raftconfig,
				// Resolver
			},
			"taskDefaults": &graphql.Field{
				Type: taskDefaults,
				// Resolver
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
				// Resolver
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
