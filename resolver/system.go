package resolver

import (
	graphql "github.com/neelance/graphql-go"
)

/*
	SystemVersion type resolvers
*/

func (r *systemVersionResolver) APIVersion() string {
	return r.version.APIVersion
}

func (r *systemVersionResolver) Arch() string {
	return r.version.Arch
}

func (r *systemVersionResolver) BuildTime() string {
	return r.version.BuildTime
}

func (r *systemVersionResolver) Experimental() bool {
	return r.version.Experimental
}

func (r *systemVersionResolver) GitCommit() string {
	return r.version.GitCommit
}

func (r *systemVersionResolver) GoVersion() string {
	return r.version.GoVersion
}

func (r *systemVersionResolver) KernelVersion() string {
	return r.version.KernelVersion
}

func (r *systemVersionResolver) MinAPIVersion() string {
	return r.version.MinAPIVersion
}

func (r *systemVersionResolver) Os() string {
	return r.version.Os
}

func (r *systemVersionResolver) Version() string {
	return r.version.Version
}

/*
	SystemInfo type resolvers
*/

func (r *systemInfoResolver) Architecture() *string {
	return &r.info.Architecture
}

func (r *systemInfoResolver) BridgeNfIP6tables() *bool {
	return &r.info.BridgeNfIP6tables
}

func (r *systemInfoResolver) BridgeNfIPtables() *bool {
	return &r.info.BridgeNfIptables
}

func (r *systemInfoResolver) CgroupDriver() *string {
	return &r.info.CgroupDriver
}

func (r *systemInfoResolver) ClusterAdvertise() *string {
	return &r.info.ClusterAdvertise
}

func (r *systemInfoResolver) ClusterStore() *string {
	return &r.info.ClusterStore
}

func (r *systemInfoResolver) ContainerdCommit() *systemInfoCommitResolver {
	c := systemInfoCommitResolver(r.info.ContainerdCommit)
	return &c
}

func (r *systemInfoResolver) Containers() *int32 {
	c := int32(r.info.Containers)
	return &c
}

func (r *systemInfoResolver) ContainersPaused() *int32 {
	cp := int32(r.info.ContainersPaused)
	return &cp
}

func (r *systemInfoResolver) ContainersRunning() *int32 {
	cr := int32(r.info.ContainersRunning)
	return &cr
}

func (r *systemInfoResolver) ContainersStopped() *int32 {
	cs := int32(r.info.ContainersStopped)
	return &cs
}

func (r *systemInfoResolver) CPUCfsPeriod() *bool {
	return &r.info.CPUCfsPeriod
}

func (r *systemInfoResolver) CPUCfsQuota() *bool {
	return &r.info.CPUCfsQuota
}

func (r *systemInfoResolver) CPUSet() *bool {
	return &r.info.CPUSet
}

func (r *systemInfoResolver) CPUShares() *bool {
	return &r.info.CPUShares
}

func (r *systemInfoResolver) Debug() *bool {
	return &r.info.Debug
}

func (r *systemInfoResolver) DefaultRuntime() *string {
	return &r.info.DefaultRuntime
}

func (r *systemInfoResolver) DockerRootDir() *string {
	return &r.info.DockerRootDir
}

func (r *systemInfoResolver) Driver() *string {
	return &r.info.Driver
}

func (r *systemInfoResolver) ExperimentalBuild() *bool {
	return &r.info.ExperimentalBuild
}

func (r *systemInfoResolver) HTTPProxy() *string {
	return &r.info.HTTPProxy
}

func (r *systemInfoResolver) HTTPSProxy() *string {
	return &r.info.HTTPSProxy
}

func (r *systemInfoResolver) ID() graphql.ID {
	return graphql.ID(r.info.ID)
}

func (r *systemInfoResolver) Images() *int32 {
	i := int32(r.info.Images)
	return &i
}

func (r *systemInfoResolver) IndexServerAddress() *string {
	return &r.info.IndexServerAddress
}

func (r *systemInfoResolver) InitBinary() *string {
	return &r.info.InitBinary
}

func (r *systemInfoResolver) InitCommit() *systemInfoCommitResolver {
	c := systemInfoCommitResolver(r.info.InitCommit)
	return &c
}

func (r *systemInfoResolver) IPv4Forwarding() *bool {
	return &r.info.IPv4Forwarding
}

func (r *systemInfoResolver) KernelMemory() *bool {
	return &r.info.KernelMemory
}

func (r *systemInfoResolver) KernelVersion() *string {
	return &r.info.KernelVersion
}

func (r *systemInfoResolver) Labels() *[]string {
	return &r.info.Labels
}

func (r *systemInfoResolver) LiveRestoreEnabled() *bool {
	return &r.info.LiveRestoreEnabled
}

func (r *systemInfoResolver) LoggingDriver() *string {
	return &r.info.LoggingDriver
}

func (r *systemInfoResolver) MemTotal() *int32 {
	m := int32(r.info.MemTotal)
	return &m
}

func (r *systemInfoResolver) MemoryLimit() *bool {
	return &r.info.MemoryLimit
}

func (r *systemInfoResolver) NCPU() *int32 {
	n := int32(r.info.NCPU)
	return &n
}

func (r *systemInfoResolver) NEventsListener() *int32 {
	n := int32(r.info.NEventsListener)
	return &n
}

func (r *systemInfoResolver) NFd() *int32 {
	n := int32(r.info.NFd)
	return &n
}

func (r *systemInfoResolver) NGoRoutines() *int32 {
	n := int32(r.info.NGoroutines)
	return &n
}

func (r *systemInfoResolver) Name() *string {
	return &r.info.Name
}

func (r *systemInfoResolver) NoProxy() *string {
	return &r.info.NoProxy
}

func (r *systemInfoResolver) OomKillDisable() *bool {
	return &r.info.OomKillDisable
}

func (r *systemInfoResolver) OperatingSystem() *string {
	return &r.info.OperatingSystem
}

func (r *systemInfoResolver) OsType() *string {
	return &r.info.OSType
}

func (r *systemInfoResolver) Plugins() *systemInfoPluginResolver {
	p := r.info.Plugins
	return &systemInfoPluginResolver{plugins: p}
}

func (r *systemInfoResolver) RegistryConfig() *systemInfoRegistryResolver {
	reg := r.info.RegistryConfig
	return &systemInfoRegistryResolver{registryConfig: *reg}
}

func (r *systemInfoResolver) RuncCommit() *systemInfoCommitResolver {
	c := systemInfoCommitResolver(r.info.RuncCommit)
	return &c
}

func (r *systemInfoResolver) SecurityOptions() *[]string {
	return &r.info.SecurityOptions
}

func (r *systemInfoResolver) ServerVersion() *string {
	return &r.info.ServerVersion
}

func (r *systemInfoResolver) SwapLimit() *bool {
	return &r.info.SwapLimit
}

func (r *systemInfoResolver) SystemTime() *string {
	return &r.info.SystemTime
}

/*
	SystemInfoCommit type resolver
*/

func (c *systemInfoCommitResolver) Id() graphql.ID {
	return graphql.ID(c.ID)
}

func (c *systemInfoCommitResolver) ExpecteD() string {
	return c.Expected
}

/*
	SystemInfoPlugin type resolver
*/

func (r *systemInfoPluginResolver) Authorizations() *[]string {
	return &r.plugins.Authorization
}

func (r *systemInfoPluginResolver) Logs() *[]string {
	return &r.plugins.Log
}

func (r *systemInfoPluginResolver) Networks() *[]string {
	return &r.plugins.Network
}

func (r *systemInfoPluginResolver) Volumes() *[]string {
	return &r.plugins.Volume
}

/*
	SystemInfoRegistry type resolver
*/

func (r *systemInfoRegistryResolver) AllowNondistributableArtifactsCidrs() *[]string {
	// Starts an slice with len 0 and then appends elements to it
	// So it will not output empty string
	// Such as:
	// "registryConfig": {
	//      "allowNondistributableArtifactsCidrs": [
	//        "", # This occurs when the slice is started with len 1
	//      ]
	cidrs := make([]string, 0)
	for _, cidr := range r.registryConfig.AllowNondistributableArtifactsCIDRs {
		cidrs = append(cidrs, cidr.String())
	}
	return &cidrs
}

func (r *systemInfoRegistryResolver) AllowNondistributableArtifactsHostnames() *[]string {
	return &r.registryConfig.AllowNondistributableArtifactsHostnames
}

// WIP
// func (r *systemInfoRegistryResolver) IndexConfigs() *[]string {
// 	i := r.registryConfig.IndexConfigs
// 	indexes := make([]string, 0)
// 	for _, index := range i {
// 		indexes = append(indexes, index.Name)
// 	}
// 	return &indexes
// }

func (r *systemInfoRegistryResolver) InsecureRegistryCidrs() *[]string {
	// Starts an slice with len 0 and then appends elements to it
	// So it will not output empty string
	// Such as:
	// "registryConfig": {
	//      "insecureRegistryCidrs": [
	//        "", # This occurs when the slice is started with len 1
	//        "127.0.0.0/8"
	//      ]
	cidrs := make([]string, 0)
	for _, cidr := range r.registryConfig.InsecureRegistryCIDRs {
		cidrs = append(cidrs, cidr.String())
	}
	return &cidrs
}

func (r *systemInfoRegistryResolver) Mirrors() *[]string {
	return &r.registryConfig.Mirrors
}
