package mocks

type MockReplicatedTable struct {
	Name     string
	Cols     []string
	ColTypes map[string]string
}

func (m *MockReplicatedTable) GetTableName() string              { return m.Name }
func (m *MockReplicatedTable) GetColumns() []string              { return m.Cols }
func (m *MockReplicatedTable) GetColumnTypes() map[string]string { return m.ColTypes }
