package schema

var (
	system = `
	type System {
		info: SystemInfo
		version: SystemVersion
	}
	type SystemVersion {
		apiVersion: String!
		arch: String!
		buildTime: String!
		experimental: Boolean!
		gitCommit: String!
		goVersion: String!
		kernelVersion: String!
		minApiVersion: String!
		os: String!
		version: String!
	}
	type SystemInfo {
		### TBI
		# driverStatus: [[String]] # [][2]string
		# genericResources: [SwarmGenericResource]
		# runtimes: [Runtime] # Returns a map[string]Runtime
		# systemStatus: [[String]] # [][2]string

		architecture: String
		bridgeNfIp6tables: Boolean
		bridgeNfIptables: Boolean
		cgroupDriver: String
		clusterAdvertise: String
		clusterStore: String
		containerdCommit: SystemCommit
		containers: Int
		containersPaused: Int
		containersRunning: Int
		containersStopped: Int
		cpuCfsPeriod: Boolean
		cpuCfsQuota: Boolean
		cpuSet: Boolean
		cpuShares: Boolean
		debug: Boolean
		defaultRuntime: String
		dockerRootDir: String
		driver: String
		experimentalBuild: Boolean
		httpProxy: String
		httpsProxy: String
		id: ID!
		images: Int
		indexServerAddress: String
		initBinary: String
		initCommit: SystemCommit
		ipv4Forwarding: Boolean
		isolation: String #ContainerIsolation
		kernelMemory: Boolean
		kernelVersion: String
		labels: [String!]
		liveRestoreEnabled: Boolean
		loggingDriver: String
		memTotal: Int
		memoryLimit: Boolean
		nCpu: Int
		nEventsListener: Int
		nFd: Int
		nGoRoutines: Int
		name: String
		noProxy: String
		oomKillDisable: Boolean
		operatingSystem: String
		osType: String
		plugins: Plugins
		registryConfig: RegistryConfig
		runcCommit: SystemCommit
		securityOptions: [String!]
		serverVersion: String
		swapLimit: Boolean
		# swarm: SwarmInfo
		systemTime: String
	}

	type SystemCommit {
		id: ID!
		expected: String!
	}

	type Plugins{
		authorizations: [String!]
		logs: [String!]
		networks: [String!]
		volumes: [String!]
	}
	type RegistryConfig {
		allowNondistributableArtifactsCidrs: [String!]
		allowNondistributableArtifactsHostnames: [String!]
		indexConfigs: [String!]
		insecureRegistryCidrs: [String!]
		mirrors: [String!]
	}
	`
)
