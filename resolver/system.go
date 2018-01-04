package resolver

import (
	"encoding/json"
	"strings"

	graphql "gitlab.com/ulm0/graphql-go"
)

/* Version retrieves server version related info */

func (r *systemResolver) Version() *systemVersionResolver {
	v, _ := r.system.ServerVersion(ctx)
	return &systemVersionResolver{version: v}
}

/*	SystemVersion type resolvers */

func (r *systemVersionResolver) APIVersion() string    { return r.version.APIVersion }
func (r *systemVersionResolver) Arch() string          { return r.version.Arch }
func (r *systemVersionResolver) BuildTime() string     { return r.version.BuildTime }
func (r *systemVersionResolver) Experimental() bool    { return r.version.Experimental }
func (r *systemVersionResolver) GitCommit() string     { return r.version.GitCommit }
func (r *systemVersionResolver) GoVersion() string     { return r.version.GoVersion }
func (r *systemVersionResolver) KernelVersion() string { return r.version.KernelVersion }
func (r *systemVersionResolver) MinAPIVersion() string { return r.version.MinAPIVersion }
func (r *systemVersionResolver) Os() string            { return r.version.Os }
func (r *systemVersionResolver) Version() string       { return r.version.Version }

/* Info retrieves a docker host info */

func (r *systemResolver) Info() *systemInfoResolver {
	i, _ := r.system.Info(ctx)
	return &systemInfoResolver{info: i}
}

/*	SystemInfo type resolvers */

func (r *systemInfoResolver) Architecture() *string     { return &r.info.Architecture }
func (r *systemInfoResolver) BridgeNfIP6tables() *bool  { return &r.info.BridgeNfIP6tables }
func (r *systemInfoResolver) BridgeNfIPtables() *bool   { return &r.info.BridgeNfIptables }
func (r *systemInfoResolver) CgroupDriver() *string     { return &r.info.CgroupDriver }
func (r *systemInfoResolver) ClusterAdvertise() *string { return &r.info.ClusterAdvertise }
func (r *systemInfoResolver) ClusterStore() *string     { return &r.info.ClusterStore }
func (r *systemInfoResolver) ContainerdCommit() *systemInfoCommitResolver {
	return ptrCommit(r.info.ContainerdCommit)
}
func (r *systemInfoResolver) Containers() *int32          { return ptrInt32(r.info.Containers) }
func (r *systemInfoResolver) ContainersPaused() *int32    { return ptrInt32(r.info.ContainersPaused) }
func (r *systemInfoResolver) ContainersRunning() *int32   { return ptrInt32(r.info.ContainersRunning) }
func (r *systemInfoResolver) ContainersStopped() *int32   { return ptrInt32(r.info.ContainersStopped) }
func (r *systemInfoResolver) CPUCfsPeriod() *bool         { return &r.info.CPUCfsPeriod }
func (r *systemInfoResolver) CPUCfsQuota() *bool          { return &r.info.CPUCfsQuota }
func (r *systemInfoResolver) CPUSet() *bool               { return &r.info.CPUSet }
func (r *systemInfoResolver) CPUShares() *bool            { return &r.info.CPUShares }
func (r *systemInfoResolver) Debug() *bool                { return &r.info.Debug }
func (r *systemInfoResolver) DefaultRuntime() *string     { return &r.info.DefaultRuntime }
func (r *systemInfoResolver) DockerRootDir() *string      { return &r.info.DockerRootDir }
func (r *systemInfoResolver) Driver() *string             { return &r.info.Driver }
func (r *systemInfoResolver) ExperimentalBuild() *bool    { return &r.info.ExperimentalBuild }
func (r *systemInfoResolver) HTTPProxy() *string          { return &r.info.HTTPProxy }
func (r *systemInfoResolver) HTTPSProxy() *string         { return &r.info.HTTPSProxy }
func (r *systemInfoResolver) ID() graphql.ID              { return graphql.ID(r.info.ID) }
func (r *systemInfoResolver) Images() *int32              { return ptrInt32(r.info.Images) }
func (r *systemInfoResolver) IndexServerAddress() *string { return &r.info.IndexServerAddress }
func (r *systemInfoResolver) InitBinary() *string         { return &r.info.InitBinary }
func (r *systemInfoResolver) InitCommit() *systemInfoCommitResolver {
	return ptrCommit(r.info.InitCommit)
}
func (r *systemInfoResolver) Isolation() *string {
	i := string(r.info.Isolation)
	if (strings.ToLower(string(i)) == "default" || string(i) == "") && r.info.OSType == "linux" {
		def := "default"
		return &def
	}
	return nil
}
func (r *systemInfoResolver) IPv4Forwarding() *bool     { return &r.info.IPv4Forwarding }
func (r *systemInfoResolver) KernelMemory() *bool       { return &r.info.KernelMemory }
func (r *systemInfoResolver) KernelVersion() *string    { return &r.info.KernelVersion }
func (r *systemInfoResolver) Labels() *[]string         { return &r.info.Labels }
func (r *systemInfoResolver) LiveRestoreEnabled() *bool { return &r.info.LiveRestoreEnabled }
func (r *systemInfoResolver) LoggingDriver() *string    { return &r.info.LoggingDriver }
func (r *systemInfoResolver) MemTotal() *int32          { return ptrInt32(r.info.MemTotal) }
func (r *systemInfoResolver) MemoryLimit() *bool        { return &r.info.MemoryLimit }
func (r *systemInfoResolver) NCPU() *int32              { return ptrInt32(r.info.NCPU) }
func (r *systemInfoResolver) NEventsListener() *int32   { return ptrInt32(r.info.NEventsListener) }
func (r *systemInfoResolver) NFd() *int32               { return ptrInt32(r.info.NFd) }
func (r *systemInfoResolver) NGoRoutines() *int32       { return ptrInt32(r.info.NGoroutines) }
func (r *systemInfoResolver) Name() *string             { return &r.info.Name }
func (r *systemInfoResolver) NoProxy() *string          { return &r.info.NoProxy }
func (r *systemInfoResolver) OomKillDisable() *bool     { return &r.info.OomKillDisable }
func (r *systemInfoResolver) OperatingSystem() *string  { return &r.info.OperatingSystem }
func (r *systemInfoResolver) OsType() *string           { return &r.info.OSType }
func (r *systemInfoResolver) Plugins() *systemInfoPluginResolver {
	return &systemInfoPluginResolver{plugins: r.info.Plugins}
}
func (r *systemInfoResolver) RegistryConfig() *systemInfoRegistryResolver {
	// reg := r.info.RegistryConfig
	return &systemInfoRegistryResolver{registryConfig: *r.info.RegistryConfig}
}
func (r *systemInfoResolver) RuncCommit() *systemInfoCommitResolver {
	return ptrCommit(r.info.RuncCommit)
}
func (r *systemInfoResolver) SecurityOptions() *[]string { return &r.info.SecurityOptions }
func (r *systemInfoResolver) ServerVersion() *string     { return &r.info.ServerVersion }
func (r *systemInfoResolver) SwapLimit() *bool           { return &r.info.SwapLimit }
func (r *systemInfoResolver) Swarm() *swarmInfoResolver {
	// swarm := r.info.Swarm
	return &swarmInfoResolver{Swarm: r.info.Swarm}
}
func (r *systemInfoResolver) SystemTime() *string { return &r.info.SystemTime }

/*	SystemInfoCommit type resolver */

func (c *systemInfoCommitResolver) Id() graphql.ID   { return graphql.ID(c.ID) }
func (c *systemInfoCommitResolver) EXPECTED() string { return c.Expected }

/*	SystemInfoPlugin type resolver */

func (r *systemInfoPluginResolver) Authorizations() *[]string { return &r.plugins.Authorization }
func (r *systemInfoPluginResolver) Logs() *[]string           { return &r.plugins.Log }
func (r *systemInfoPluginResolver) Networks() *[]string       { return &r.plugins.Network }
func (r *systemInfoPluginResolver) Volumes() *[]string        { return &r.plugins.Volume }

/*	SystemInfoRegistry type resolver */

func (r *systemInfoRegistryResolver) AllowNondistributableArtifactsCidrs() *[]string {
	return ptrIPNetString(r.registryConfig.AllowNondistributableArtifactsCIDRs)
}
func (r *systemInfoRegistryResolver) AllowNondistributableArtifactsHostnames() *[]string {
	return &r.registryConfig.AllowNondistributableArtifactsHostnames
}

// WIP
func (r *systemInfoRegistryResolver) IndexConfigs() *[]string {
	configs := make([]string, 0)
	for n, c := range r.registryConfig.IndexConfigs {
		name, _ := json.Marshal(n)
		config, _ := json.Marshal(c)
		index := string(name) + ":" + string(config)
		configs = append(configs, index)
	}
	return &configs
}
func (r *systemInfoRegistryResolver) InsecureRegistryCidrs() *[]string {
	return ptrIPNetString(r.registryConfig.InsecureRegistryCIDRs)
}
func (r *systemInfoRegistryResolver) Mirrors() *[]string { return &r.registryConfig.Mirrors }
