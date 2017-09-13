package schema

var (
	query = `
	type Query {
			system: System
	}
	`
	schema = `
	schema {
		query: Query
	}
	`
	// Schema to be exported
	Schema = schema + query + system
)
