package resolver

import graphql "github.com/neelance/graphql-go"

/*
SystemVersion type resolvers
*/

func (r *systemVersionResolver) ApiVersion() string {
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

func (r *systemVersionResolver) MinApiVersion() string {
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

func (r *systemInfoResolver) CpuCfsPeriod() *bool {
	return &r.info.CPUCfsPeriod
}

func (r *systemInfoResolver) CpuCfsQuota() *bool {
	return &r.info.CPUCfsQuota
}

func (r *systemInfoResolver) CpuSet() *bool {
	return &r.info.CPUSet
}

func (r *systemInfoResolver) CpuShares() *bool {
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

func (r *systemInfoResolver) HttpProxy() *string {
	return &r.info.HTTPProxy
}

func (r *systemInfoResolver) HttpsProxy() *string {
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
