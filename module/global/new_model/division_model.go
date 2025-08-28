package new_model

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gitea.qwertysystem.net/BETS/ts-utils/migration"
)

type Division struct {
	ID        int       `json:"_id" column:"id" sql:"BIGINT PRIMARY KEY"`
	Name      string    `json:"name" column:"name" sql:"VARCHAR NOT NULL"`
	CreatedAt time.Time `json:"createdAt" column:"created_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
	UpdatedAt time.Time `json:"updatedAt" column:"updated_at" sql:"TIMESTAMPTZ NOT NULL DEFAULT NOW()"`
}

func (m *Division) TableName() string {
	return "divisions"
}

func (m *Division) Up(builder *migration.MigrationBuilder) {
	builder.AddColumn(m)

	// Insert default data
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)
	path := filepath.Join(baseDir, "migration/division_post_query.sql")

	postQuery, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error path: " + path)
		log.Fatal(err)
	}
	builder.SetPostQuery(string(postQuery))
}
