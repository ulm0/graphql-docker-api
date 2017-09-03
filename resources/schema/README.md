# Schema resources used in the wrapper

It's important to keep this schema synced with the types implemented in the wrapper, best efforts are made in order to keep it so. That being said, you need to know *this schema is subject to change* as the time goes by and new changes are introduced in the Docker API and this wrapper as well.

You can check the query-ready fields in the [TODO list](../../README.md#todo)

## Some conventions

* Commented out types with `TBI` comment above means `to be implement`, therefore these are not ready to be queried yet.
* Unused types (at all) are commented out with `# @deprecated`, this means these fields were not present at the time this wrapper was written.

## Schema example

```graphql
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
```