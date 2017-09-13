# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] 2017-09-13

- [SystemVersion type](resources/schema/system.graphql#L6) is the only query-ready field for now.

### Changed

- Using [neelance/graphql-go](https://github.com/neelance/graphql-go).
- `API_ENDPOINT` is now `UI_ENDPOINT`.
- `GQL_PORT` is available again.

### Removed

- No longer using [graphql-go/graphql](https://github.com/graphql-go/graphql)

## [0.1.0-rc] 2017-09-11

### Added

- Start using a changelog file in order to keep a neat track of every change made to the project.
- [Version type](resources/schema/system.graphql#L6) is the only query-ready field for now.

### Changed

- Moving from [graphql-go/graphql](https://github.com/graphql-go/graphql) to [neelance/graphql-go](https://github.com/neelance/graphql-go).
- `API_ENDPOINT`, `GQL_PORT` & `GRAPHIQL` are temporarily unavailable to be set in the Docker image.

### Removed

- No longer using [graphql-go/graphql](https://github.com/graphql-go/graphql)