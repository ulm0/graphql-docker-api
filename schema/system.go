package schema

import "github.com/graphql-go/graphql"

var (
	// Struct to retrive data from in SystemVersion
	sysVer, _ = cli.ServerVersion(ctx)

	sysInfo, _ = cli.Info(ctx)
	// System related custom object types
	plugins = graphql.NewObject(graphql.ObjectConfig{
		Name: "Plugins",
		Fields: graphql.Fields{
			"authorizations": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Plugins.Authorization, nil
				},
			},
			// "logs": &graphql.Field{
			// 	Type: graphql.NewList(graphql.String),
			// 	// Resolver
			// },
			"networks": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Plugins.Network, nil
				},
			},
			"volumes": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Plugins.Volume, nil
				},
			},
		},
	})
	registryConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "RegistryConfig",
		Fields: graphql.Fields{
			// TBI
			// "indexConfigs": &graphql.Field{
			// 	Type: graphql.NewList(graphql.String),
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return sysInfo.RegistryConfig.IndexConfigs, nil
			// 	},
			// },
			"insecureRegistryCidrs": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.RegistryConfig.InsecureRegistryCIDRs, nil
				},
			},
			"mirrors": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.RegistryConfig.Mirrors, nil
				},
			},
		},
	})
	system = graphql.NewObject(graphql.ObjectConfig{
		Name: "System",
		Fields: graphql.Fields{
			"info": &graphql.Field{
				Type: graphql.NewNonNull(systemInfo),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo, nil
				},
			},
			"version": &graphql.Field{
				Type: graphql.NewNonNull(systemVersion),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer, nil
				},
			},
		},
	})
	systemInfo = graphql.NewObject(graphql.ObjectConfig{
		Name: "SystemInfo",
		Fields: graphql.Fields{
			"architecture": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Architecture, nil
				},
			},
			"bridgeNfIp6tables": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.BridgeNfIP6tables, nil
				},
			},
			"bridgeNfIptables": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.BridgeNfIptables, nil
				},
			},
			"cgroupDriver": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.CgroupDriver, nil
				},
			},
			"clusterAdvertise": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ClusterAdvertise, nil
				},
			},
			"clusterStore": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ClusterStore, nil
				},
			},
			"containers": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Containers, nil
				},
			},
			"containersPaused": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ContainersPaused, nil
				},
			},
			"containersRunning": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ContainersRunning, nil
				},
			},
			"containersStopped": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ContainersStopped, nil
				},
			},
			"cpuCfsPeriod": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.CPUCfsPeriod, nil
				},
			},
			"cpuCfsQuota": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.CPUCfsQuota, nil
				},
			},
			"cpuSet": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.CPUSet, nil
				},
			},
			"cpuShares": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.CPUShares, nil
				},
			},
			"debug": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Debug, nil
				},
			},
			"defaultRuntime": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.DefaultRuntime, nil
				},
			},
			"dockerRootDir": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.DockerRootDir, nil
				},
			},
			"driver": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Driver, nil
				},
			},
			"experimentalBuild": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ExperimentalBuild, nil
				},
			},
			"httpProxy": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.HTTPProxy, nil
				},
			},
			"httpsProxy": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.HTTPSProxy, nil
				},
			},
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ID, nil
				},
			},
			"images": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Images, nil
				},
			},
			"indexServerAddress": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.IndexServerAddress, nil
				},
			},
			"initBinary": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.InitBinary, nil
				},
			},
			"ipv4Forwarding": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.IPv4Forwarding, nil
				},
			},
			// TBI
			// "isolation": &graphql.Field{
			// 	Type: containerIsolation,
			// 	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// 		return sysInfo.Isolation, nil
			// 	},
			// },
			"kernelMemory": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.KernelMemory, nil
				},
			},
			"kernelVersion": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.KernelVersion, nil
				},
			},
			"labels": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Labels, nil
				},
			},
			"liveRestoreEnabled": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.LiveRestoreEnabled, nil
				},
			},
			"loggingDriver": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.LoggingDriver, nil
				},
			},
			"memTotal": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return int32(sysInfo.MemTotal), nil
				},
			},
			"memoryLimit": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.MemoryLimit, nil
				},
			},
			"nCpu": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.NCPU, nil
				},
			},
			"nEventsListener": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.NEventsListener, nil
				},
			},
			"nFd": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.NFd, nil
				},
			},
			"nGoRoutines": &graphql.Field{
				Type: graphql.Int,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.NGoroutines, nil
				},
			},
			"name": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Name, nil
				},
			},
			"noProxy": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.NoProxy, nil
				},
			},
			"oomKillDisable": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.OomKillDisable, nil
				},
			},
			"operatingSystem": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.OperatingSystem, nil
				},
			},
			"osType": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.OSType, nil
				},
			},
			"plugins": &graphql.Field{
				Type: plugins,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Plugins, nil
				},
			},
			"registryConfig": &graphql.Field{
				Type: registryConfig,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.RegistryConfig, nil
				},
			},
			"securityOptions": &graphql.Field{
				Type: graphql.NewList(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.SecurityOptions, nil
				},
			},
			"serverVersion": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.ServerVersion, nil
				},
			},
			"swapLimit": &graphql.Field{
				Type: graphql.Boolean,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.SwapLimit, nil
				},
			},
			"swarm": &graphql.Field{
				Type: swarmInfo,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.Swarm, nil
				},
			},
			"systemTime": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysInfo.SystemTime, nil
				},
			},
		},
	})
	systemVersion = graphql.NewObject(graphql.ObjectConfig{
		Name: "SystemVersion",
		Fields: graphql.Fields{
			"apiVersion": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.APIVersion, nil
				},
			},
			"arch": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.Arch, nil
				},
			},
			"buildTime": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.BuildTime, nil
				},
			},
			"experimental": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.Experimental, nil
				},
			},
			"gitCommit": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.GitCommit, nil
				},
			},
			"goVersion": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.GoVersion, nil
				},
			},
			"kernelVersion": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.KernelVersion, nil
				},
			},
			"minApiVersion": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.MinAPIVersion, nil
				},
			},
			"os": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.Os, nil
				},
			},
			"version": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return sysVer.Version, nil
				},
			},
		},
	})
)
