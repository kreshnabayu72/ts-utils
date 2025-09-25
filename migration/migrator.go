package migration

type Model interface {
	TableName() string
	Up(builder *MigrationBuilder)
}
