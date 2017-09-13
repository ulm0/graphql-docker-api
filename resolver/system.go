package resolver

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
