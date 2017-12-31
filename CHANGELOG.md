# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.2] 2017-12-30

### Changed

- [Swarm](resources/schema/swarm.graphql#L1) updated fields.
- `UI_ENDPOINT` is now `API_ENDPOINT`
- Using personal fork of `graphql-go`

## [0.1.1] 2017-09-25

### Added

- Types:
  - [Swarm](resources/schema/swarm.graphql#L1) mostly working.
  - [SwarmSpec](resources/schema/swarm.graphql#L11) mostly working.
  - [OrchestrationConfig](resources/schema/swarm.graphql#L23) working.
  - [RaftConfig](resources/schema/swarm.graphql#L27) working.
  - [DispatcherConfig](resources/schema/swarm.graphql#L35) working.
  - [CaConfig](resources/schema/swarm.graphql#L39) mostly working.
  - [ExternalCa](resources/schema/swarm.graphql#L44), kinda working.
  - [EncryptionConfig](resources/schema/swarm.graphql#L51) working.
  - [TaskDefaults](resources/schema/swarm.graphql#L55) not working properly.
  - [Driver](resources/schema/swarm.graphql#L59) not working properly.
  - [JoinTokens](resources/schema/swarm.graphql#L65) working.
  - [SwarmInfo](resources/schema/swarm.graphql#L161) mostly working.
  - [SwarmCluster](resources/schema/swarm.graphql#L173) mostly working.
  - [TLSInfo](resources/schema/swarm.graphql#L187) fully working.
  - [RemoteManager](resources/schema/swarm.graphql#L193) fully working.
  - [SystemInfo](resources/schema/system.graphql#L19) mostly working.
  - [RegistryConfig](resources/schema/system.graphql#L93) mostly working.
  - [Plugins](resources/schema/system.graphql#L86) fully working.
  - [SystemCommit](resources/schema/system.graphql#L81) added.

## [0.1.0] 2017-09-13

- [SystemVersion type](resources/schema/system.graphql#L6) is the only query-ready field for now.

### Changed

- Using [neelance/graphql-go](https://gitlab.com/klud/graphql-go).
- `API_ENDPOINT` is now `UI_ENDPOINT`.
- `GQL_PORT` is available again.

### Removed

- No longer using [graphql-go/graphql](https://github.com/graphql-go/graphql)

## [0.1.0-rc] 2017-09-11

### Added

- Start using a changelog file in order to keep a neat track of every change made to the project.
- [Version type](resources/schema/system.graphql#L6) is the only query-ready field for now.

### Changed

- Moving from [graphql-go/graphql](https://github.com/graphql-go/graphql) to [neelance/graphql-go](https://gitlab.com/klud/graphql-go).
- `API_ENDPOINT`, `GQL_PORT` & `GRAPHIQL` are temporarily unavailable to be set in the Docker image.

### Removed

- No longer using [graphql-go/graphql](https://github.com/graphql-go/graphql)