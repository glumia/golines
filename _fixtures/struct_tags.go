package fixtures

type MyEntity struct {
	tableName struct{} `pg:"my_entity" pg:",discard_unknown_columns"`
}
