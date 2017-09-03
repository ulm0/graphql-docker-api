package schema

import "github.com/graphql-go/graphql"

var (
	// Container related custom object types
	blkioDeviceRate = graphql.NewObject(graphql.ObjectConfig{
		Name: "BlkioDeviceRate",
		Fields: graphql.Fields{
			"path": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"rate": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
		},
	})
	blkioWeightDevice = graphql.NewObject(graphql.ObjectConfig{
		Name: "BlkioWieghtDevice",
		Fields: graphql.Fields{
			"path": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"weight": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
		},
	})
	container = graphql.NewObject(graphql.ObjectConfig{
		Name: "Container",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"names": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"changes": &graphql.Field{
				Type: containerChange,
				// Side Notes: this needs to return an slice of mount containerChange
				// so probably the type needs to be tweaked
				// Resolver
			},
			"command": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"details": &graphql.Field{
				Type: containerDetails,
				// Resolver
			},
			"image": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"imageId": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"logs": &graphql.Field{
				Type: graphql.String,
				// Side Note: this needs to receiver args
				// (args: follow bool, stdout bool, stderr bool, since int,
				// timestamps bool, lines int, custompaths stringanymap)
				// Resolver
			},
			"mounts": &graphql.Field{
				Type: containerMount,
				// Side Notes: this needs to return an slice of mount points
				// so probably the type needs to be tweaked
				// Resolver
			},
			"networkingSettings": &graphql.Field{
				Type: networkSettings,
				// Resovler
			},
			"ports": &graphql.Field{
				Type: portMapping,
				// Side Note: this needs to return an slice of ports
				// so probably the type needs to be tweaked
				// Resolver
			},
			"portsAsString": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"processes": &graphql.Field{
				Type: topResults,
				// Side Note: this needs to receive args (args: ps_args string)
				// Resolver
			},
			"sizeRootFs": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"sizeRw": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"state": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"stats": &graphql.Field{
				Type: containerStats,
				// Resolver
			},
			"status": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	containerChange = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerChange",
		Fields: graphql.Fields{
			"kind": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"path": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	containerConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerConfig",
		Fields: graphql.Fields{
			"attachStderr": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"attachStdin": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"attachStdout": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"cmd": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"domainname": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"entrypoint": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"env": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"exposedPorts": &graphql.Field{
				Type: stringSet,
				// Resolver
			},
			"healthcheck": &graphql.Field{
				Type: healthCheck,
				// Resolver
			},
			"hostConfig": &graphql.Field{
				Type: hostConfig,
				// Resolver
			},
			"hostname": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"image": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"labels": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"macAddress": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"networkDisabled": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"networkingConfig": &graphql.Field{
				Type: networkingConfig,
				// Resolver
			},
			"onBuild": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"openStdin": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"portSpecs": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"stdinOnce": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"stopSignal": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"tty": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"user": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"volumeNames": &graphql.Field{
				Type: stringSet,
				// Resolver
			},
			"volumes": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"workingDir": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	containerDetails = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerDetails",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"image": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"appArmorProfile": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"args": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"config": &graphql.Field{
				Type: containerConfig,
				// Resolver
			},
			"created": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"driver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"execDriver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"execIds": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"hostConfig": &graphql.Field{
				Type: hostConfig,
				// Resolver
			},
			"hostnamePath": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"hostsPath": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"logPath": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"mountLabel": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"mounts": &graphql.Field{
				Type: containerMount,
				// Side Notes: this needs to return an slice of containerMount
				// so probably the type needs to be tweaked
				// Resolver
			},
			"networkSettings": &graphql.Field{
				Type: networkSettings,
				// Resolver
			},
			"node": &graphql.Field{
				Type: node,
				// Resolver
			},
			"path": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"processLabel": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"resolvConfPath": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"restartCount": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"state": &graphql.Field{
				Type: containerState,
				// Resolver
			},
			"volumes": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"volumesRw": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
		},
	})
	containerFilter = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerFilter",
		Fields: graphql.Fields{
			"ancestor": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"before": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"exited": &graphql.Field{
				Type: graphql.Int,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"expose": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"health": &graphql.Field{
				Type: containerHealth,
				// Side Notes: this needs to return an slice of containerHealth
				// so probably the type needs to be tweaked
				// Resolver
			},
			"id": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"is_task": &graphql.Field{
				Type: graphql.Boolean,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"isolation": &graphql.Field{
				Type: containerIsolation,
				// Side Notes: this needs to return an slice of containerIsolation
				// so probably the type needs to be tweaked
				// Resolver
			},
			"label": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"network": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"publish": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"since": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"status": &graphql.Field{
				Type: containerStatus,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"volume": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	containerHealth = graphql.NewEnum(graphql.EnumConfig{
		Name: "ContainerHealth",
		Values: graphql.EnumValueConfigMap{
			"healthy": &graphql.EnumValueConfig{
				Value: "healthy",
			},
			"none": &graphql.EnumValueConfig{
				Value: "none",
			},
			"starting": &graphql.EnumValueConfig{
				Value: "starting",
			},
			"unhealthy": &graphql.EnumValueConfig{
				Value: "unhealthy",
			},
		},
	})
	containerIsolation = graphql.NewEnum(graphql.EnumConfig{
		Name: "ContainerIsolation",
		Values: graphql.EnumValueConfigMap{
			"default": &graphql.EnumValueConfig{
				Value: "default",
			},
			"hyperv": &graphql.EnumValueConfig{
				Value: "hyperv",
			},
			"process": &graphql.EnumValueConfig{
				Value: "process",
			},
		},
	})
	containerMount = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerMount",
		Fields: graphql.Fields{
			"destination": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"driver": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"mode": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"propagation": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"rw": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"source": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"type": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	containerState = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerState",
		Fields: graphql.Fields{
			"error": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"exitCode": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"finishedAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"health": &graphql.Field{
				Type: health,
				// Resolver
			},
			"oomKilled": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"paused": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"pid": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"restarting": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"running": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"startedAt": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"status": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	containerStats = graphql.NewObject(graphql.ObjectConfig{
		Name: "ContainerStats",
		Fields: graphql.Fields{
			"cpuStats": &graphql.Field{
				Type: cpuStats,
				// Resolver
			},
			"memoryStats": &graphql.Field{
				Type: memoryStats,
				// Resolver
			},
			"network": &graphql.Field{
				Type: networkStats,
				// Resolver
			},
			"networks": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"precpuStats": &graphql.Field{
				Type: cpuStats,
				// Resolver
			},
			"read": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
		},
	})
	containerStatus = graphql.NewEnum(graphql.EnumConfig{
		Name: "ContainerStatus",
		Values: graphql.EnumValueConfigMap{
			"created": &graphql.EnumValueConfig{
				Value: "created",
			},
			"dead": &graphql.EnumValueConfig{
				Value: "dead",
			},
			"exited": &graphql.EnumValueConfig{
				Value: "exited",
			},
			"paused": &graphql.EnumValueConfig{
				Value: "paused",
			},
			"removing": &graphql.EnumValueConfig{
				Value: "removing",
			},
			"restarting": &graphql.EnumValueConfig{
				Value: "restarting",
			},
			"running": &graphql.EnumValueConfig{
				Value: "running",
			},
		},
	})
	cpuStats = graphql.NewObject(graphql.ObjectConfig{
		Name: "CpuStats",
		Fields: graphql.Fields{
			"cpuUsage": &graphql.Field{
				Type: cpuUsage,
				// Resolver
			},
			"systemCpuUsage": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"throttlingData": &graphql.Field{
				Type: throttlingData,
				// Resolver
			},
		},
	})
	cpuUsage = graphql.NewObject(graphql.ObjectConfig{
		Name: "CpuUsage",
		Fields: graphql.Fields{
			"percpuUsage": &graphql.Field{
				Type: graphql.Float,
				// Side Notes: this needs to return an slice of float
				// so probably the type needs to be tweaked
				// Resolver
			},
			"totalUsage": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"usageInKernelMode": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"usageInUserMode": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	device = graphql.NewObject(graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.Fields{
			"cgroupPermissions": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"pathInContainer": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"pathOnHost": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	health = graphql.NewObject(graphql.ObjectConfig{
		Name: "Health",
		Fields: graphql.Fields{
			"failingStreak": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"log": &graphql.Field{
				Type: healthLog,
				// Side Notes: this needs to return an slice of healthLog
				// so probably the type needs to be tweaked
				// Resolver
			},
			"status": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	healthCheck = graphql.NewObject(graphql.ObjectConfig{
		Name: "HealthCheck",
		Fields: graphql.Fields{
			"interval": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"retries": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"startPeriod": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"test": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"timeout": &graphql.Field{
				Type: graphql.Float,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	healthLog = graphql.NewObject(graphql.ObjectConfig{
		Name: "HealthLog",
		Fields: graphql.Fields{
			"end": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
			"exitCode": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"output": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"start": &graphql.Field{
				Type: graphql.DateTime,
				// Resolver
			},
		},
	})
	hostConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "HostConfig",
		Fields: graphql.Fields{
			"autoRemove": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"binds": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioDeviceReadBps": &graphql.Field{
				Type: blkioDeviceRate,
				// Side Notes: this needs to return an slice of blkioDeviceRate
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioDeviceReadIOps": &graphql.Field{
				Type: blkioDeviceRate,
				// Side Notes: this needs to return an slice of blkioDeviceRate
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioDeviceWriteBps": &graphql.Field{
				Type: blkioDeviceRate,
				// Side Notes: this needs to return an slice of blkioDeviceRate
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioDeviceWriteIOps": &graphql.Field{
				Type: blkioDeviceRate,
				// Side Notes: this needs to return an slice of blkioDeviceRate
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioWeight": &graphql.Field{
				Type: graphql.Int,
				// Side Notes: this needs to return an slice of blkioDeviceRate
				// so probably the type needs to be tweaked
				// Resolver
			},
			"blkioWeightDevice": &graphql.Field{
				Type: blkioWeightDevice,
				// Side Notes: this needs to return an slice of blkioWeightDevice
				// so probably the type needs to be tweaked
				// Resolver
			},
			"capAdd": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"capDrop": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"cgroupParent": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"containerIdFile": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"cpuPeriod": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"cpuQuota": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"cpuShares": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"cpusetCpus": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"cpusetMems": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"devices": &graphql.Field{
				Type: device,
				// Side Notes: this needs to return an slice of device
				// so probably the type needs to be tweaked
				// Resolver
			},
			"dns": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"dnsOption": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"dnsSearch": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"extraHosts": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"ipcMode": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"links": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"logConfig": &graphql.Field{
				Type: logConfig,
				// Resolver
			},
			"lxcConf": &graphql.Field{
				Type: lxcConfParameter,
				// Side Notes: this needs to return an slice of lxcConfParameter
				// so probably the type needs to be tweaked
				// Resolver
			},
			"memory": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"memoryReservation": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"memorySwap": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"memorySwappiness": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"nanoCpus": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"networkMode": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"oomKillDisable": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"oomScoreAdj": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"pidMode": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"pidsLimit": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"portBindings": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"privileged": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"publishAllPorts": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"readonlyRootfs": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"restartPolicy": &graphql.Field{
				Type: restartPolicy,
				// Resolver
			},
			"securityOpt": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
			"shmSize": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"storageOpt": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"tmpfs": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"ulimits": &graphql.Field{
				Type: ulimit,
				// Side Notes: this needs to return an slice of ulimit
				// so probably the type needs to be tweaked
				// Resolver
			},
			"volumesFrom": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of string
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	logConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "LogConfig",
		Fields: graphql.Fields{
			"logOptions": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"logType": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	lxcConfParameter = graphql.NewObject(graphql.ObjectConfig{
		Name: "LxcConfParameter",
		Fields: graphql.Fields{
			"key": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"value": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	memoryStats = graphql.NewObject(graphql.ObjectConfig{
		Name: "MemoryStats",
		Fields: graphql.Fields{
			"failcnt": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"limit": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"maxUsage": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"stats": &graphql.Field{
				Type: stats,
				// Resolver
			},
			"usage": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	networkSettings = graphql.NewObject(graphql.ObjectConfig{
		Name: "NetworkSettings",
		Fields: graphql.Fields{
			"bridge": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"endpointId": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"gateway": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"globalIPv6Address": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"globalIPv6PrefixLen": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"hairpinMode": &graphql.Field{
				Type: graphql.Boolean,
				// Resolver
			},
			"ipAddress": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"ipPrefixLen": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"ipv6Gateway": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"linkLocalIPv6Address": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"linkLocalIPv6PrefixLen": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"macAddress": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"networks": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"portMapping": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"ports": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
			"sandboxId": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"sandboxKey": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	networkStats = graphql.NewObject(graphql.ObjectConfig{
		Name: "NetworkStats",
		Fields: graphql.Fields{
			"rxBytes": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"rxDropped": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"rxErrors": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"rxPackets": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"txBytes": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"txDropped": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"txErrors": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"txPackets": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	networkingConfig = graphql.NewObject(graphql.ObjectConfig{
		Name: "NetworkingConfig",
		Fields: graphql.Fields{
			"endpointsConfig": &graphql.Field{
				Type: stringAnyMap,
				// Resolver
			},
		},
	})
	node = graphql.NewObject(graphql.ObjectConfig{
		Name: "Node",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
				// Resolver
			},
			"addr": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"ip": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	portMapping = graphql.NewObject(graphql.ObjectConfig{
		Name: "PortMapping",
		Fields: graphql.Fields{
			"ip": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"privatePort": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"publicPort": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"type": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	restartPolicy = graphql.NewObject(graphql.ObjectConfig{
		Name: "RestartPolicy",
		Fields: graphql.Fields{
			"maxRetryCount": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
		},
	})
	stats = graphql.NewObject(graphql.ObjectConfig{
		Name: "Stats",
		Fields: graphql.Fields{
			"activeAnon": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"activeFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"cache": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"hierarchicalMemoryLimit": &graphql.Field{
				Type: graphql.Int,
				// Resolver
			},
			"inactiveAnon": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"inactiveFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"mappedFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"pgfault": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"pgmajfault": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"pgpgin": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"pgpgout": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"rss": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"rssHuge": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalActiveAnon": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalActiveFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalCache": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalInactiveAnon": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalInactiveFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalMappedFile": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalPgfault": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalPgmajfault": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalPgpgin": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalPgpgout": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalRss": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalRssHuge": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalUnevictable": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"totalWriteback": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"unevictable": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"writeBbck": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	throttlingData = graphql.NewObject(graphql.ObjectConfig{
		Name: "ThrottlingData",
		Fields: graphql.Fields{
			"periods": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"throttlePeriods": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"throttleTime": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
	topResults = graphql.NewObject(graphql.ObjectConfig{
		Name: "TopResults",
		Fields: graphql.Fields{
			"titles": &graphql.Field{
				Type: graphql.String,
				// Side Notes: this needs to return an slice of strings
				// so probably the type needs to be tweaked
				// Resolver
			},
		},
	})
	ulimit = graphql.NewObject(graphql.ObjectConfig{
		Name: "Ulimit",
		Fields: graphql.Fields{
			"hard": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
			"name": &graphql.Field{
				Type: graphql.String,
				// Resolver
			},
			"soft": &graphql.Field{
				Type: graphql.Float,
				// Resolver
			},
		},
	})
)
