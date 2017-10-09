package schema

var image = `
	type Image {
		created: String
		details: ImageDetails
		history: [ImageHistory]
		id: ID!
		labels: StringAnyMap
		parentId: String
		repoDigests: [String]
		repoTags: [String]
		size: Long
		virtualSize: Long
	}

	input ImageFilter {
		before: [String]
		dangling: [Boolean]
		label: [String]
		reference: [String]
		since: [String]
	}

	type ImageDetails {
		architecture: String
		author: String
		comment: String
		config: ContainerConfig
		container: String
		containerConfig: ContainerConfig
		created: Date
		dockerVersion: String
		id: ID!
		os: String
		parent: String
		rootFs: RootFs
		size: Long
		virtualSize: String
	}

	type RootFs {
		layers: [String]
		type: String
	}

	type ImageHistory {
		comment: String
		created: Long
		createdBy: String
		id: ID!
		size: Long
		tags: [String]
	}

	type ImageSearchResult {
		automated: Boolean
		description: String
		name: String
		official: Boolean
		starCount: Int
	}
	`
