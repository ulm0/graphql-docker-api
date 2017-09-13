package schema

var (
	system = `
	type System {
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
	`
)
