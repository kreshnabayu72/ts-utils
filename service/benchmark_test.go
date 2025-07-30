package service

import (
	"context"
	"fmt"
	"testing"

	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// fakeRows implements pgx.Rows.
type fakeRows struct {
	cols []string
	data [][]interface{}
	idx  int
}

func newFakeRows(n int) *fakeRows {
	cols := []string{"id", "name", "email"}
	data := make([][]interface{}, n)
	for i := 0; i < n; i++ {
		data[i] = []interface{}{i, fmt.Sprintf("User%d", i), fmt.Sprintf("user%d@example.com", i)}
	}
	return &fakeRows{cols: cols, data: data}
}

func (r *fakeRows) Next() bool {
	if r.idx < len(r.data) {
		return true
	}
	return false
}

func (r *fakeRows) Values() ([]interface{}, error) {
	vals := r.data[r.idx]
	r.idx++
	return vals, nil
}

func (r *fakeRows) RawValues() [][]byte { return nil } // <— required

func (r *fakeRows) FieldDescriptions() []pgconn.FieldDescription {
	desc := make([]pgconn.FieldDescription, len(r.cols))
	for i, c := range r.cols {
		desc[i].Name = c // pgx.FieldDescription.Name is a string
	}
	return desc
}

func (r *fakeRows) Err() error                    { return nil }
func (r *fakeRows) Close()                        {}
func (r *fakeRows) CommandTag() pgconn.CommandTag { return pgconn.CommandTag{} }

// These two complete the interface:
func (r *fakeRows) Conn() *pgx.Conn                { return nil } // stub
func (r *fakeRows) Scan(dest ...interface{}) error { return nil } // stub

type fakeRow struct{}

func (r *fakeRow) Scan(dest ...interface{}) error {
	if len(dest) == 1 {
		if idPtr, ok := dest[0].(*int); ok {
			*idPtr = 1 // Pretend DB returns ID=1
		}
	}
	return nil
}

// fakePool as before…
type fakePool struct{ rowsPerQuery int }

func newFakePool(n int) *fakePool                             { return &fakePool{rowsPerQuery: n} }
func (p *fakePool) Begin(ctx context.Context) (pgx.Tx, error) { panic("unused") }
func (p *fakePool) Query(ctx context.Context, _ string, _ ...any) (pgx.Rows, error) {
	return newFakeRows(p.rowsPerQuery), nil
}
func (p *fakePool) QueryRow(ctx context.Context, _ string, _ ...any) pgx.Row { return &fakeRow{} }
func (p *fakePool) Exec(ctx context.Context, _ string, _ ...any) (pgconn.CommandTag, error) {
	return pgconn.NewCommandTag("INSERT 1"), nil
}
func (p *fakePool) Close() {}

// --- DTO for scanning ---
type userDTO struct {
	ID    int    `json:"id" column:"id"`
	Name  string `json:"name" column:"name"`
	Email string `json:"email" column:"email"`
}

// --- Benchmarks ---

// Benchmark SelectOne (always returns 1 row)
func BenchmarkSelectOne(b *testing.B) {
	query, args, _ := sql_query.NewSQLSelectBuilder[userDTO]("users").Build()
	svc := &BasePostgreSqlService{
		Pool:      newFakePool(1),
		TableName: "users",
	}
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		var u userDTO
		_ = svc.SelectOne(&u, ctx, query, args...)
	}
}

// Benchmark SelectMany with various row counts
func BenchmarkSelectMany(b *testing.B) {
	query, args, _ := sql_query.NewSQLSelectBuilder[userDTO]("users").Build()
	for _, n := range []int{1, 10, 100, 1000} {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			svc := &BasePostgreSqlService{
				Pool:      newFakePool(n),
				TableName: "users",
			}
			ctx := context.Background()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				var slice []userDTO
				_ = svc.SelectMany(&slice, ctx, query, args...)
			}
		})
	}
}

// Benchmark InsertOne (always returns 1 row)
func BenchmarkInsertOne(b *testing.B) {
	u := userDTO{
		ID:    1,
		Name:  "A",
		Email: "a",
	}
	query, args, _ := sql_query.NewSQLInsertBuilder("users").Insert(u).Build()
	svc := &BasePostgreSqlService{
		Pool:      newFakePool(1),
		TableName: "users",
	}
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		_, _ = svc.InsertOne(ctx, query, args...)
	}
}

// Benchmark InsertMany with various row counts
func BenchmarkInsertMany(b *testing.B) {
	batches := map[int][]userDTO{}
	for _, n := range []int{1, 10, 100, 1000} {
		slice := make([]userDTO, n)
		for i := 0; i < n; i++ {
			slice[i] = userDTO{ID: i, Name: "A", Email: "a"}
		}
		batches[n] = slice
	}
	for n, slice := range batches {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			// Build once per sub-benchmark
			query, args, _ := sql_query.NewSQLInsertBuilder("users").
				Insert(slice).
				Build()

			svc := &BasePostgreSqlService{
				Pool:      newFakePool(n),
				TableName: "users",
			}
			ctx := context.Background()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = svc.InsertMany(ctx, query, args...)
			}
		})
	}
}
