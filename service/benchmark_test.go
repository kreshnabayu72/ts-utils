package service

import (
	"context"
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var allStructs = []any{
	new_model.User{},
	new_model.UserSystemRole{},
	new_model.UserTransactionRole{},
	new_model.CustomField{},
	new_model.CustomFieldCategory{},
	new_model.Asset{},
	new_model.AssetName{},
	new_model.Brand{},
	new_model.Model{},
	new_model.Group{},
	new_model.Category{},
	new_model.DamageReport{},
	new_model.MissingReport{},
	new_model.DisposalReport{},
	new_model.BorrowingRequest{},
	new_model.AssignmentRequest{},
	new_model.TransferRequest{},
	new_model.Attachment{},
}

func callSelectBuilder(v any) (string, []interface{}, error) {
	switch v.(type) {
	case new_model.User:
		return sql_query.NewSQLSelectBuilder[new_model.User]("test").Build()
	case new_model.UserSystemRole:
		return sql_query.NewSQLSelectBuilder[new_model.UserSystemRole]("test").Build()
	case new_model.UserTransactionRole:
		return sql_query.NewSQLSelectBuilder[new_model.UserTransactionRole]("test").Build()
	case new_model.CustomField:
		return sql_query.NewSQLSelectBuilder[new_model.CustomField]("test").Build()
	case new_model.CustomFieldCategory:
		return sql_query.NewSQLSelectBuilder[new_model.CustomFieldCategory]("test").Build()
	case new_model.Asset:
		return sql_query.NewSQLSelectBuilder[new_model.Asset]("test").Build()
	case new_model.AssetName:
		return sql_query.NewSQLSelectBuilder[new_model.AssetName]("test").Build()
	case new_model.Brand:
		return sql_query.NewSQLSelectBuilder[new_model.Brand]("test").Build()
	case new_model.Model:
		return sql_query.NewSQLSelectBuilder[new_model.Model]("test").Build()
	case new_model.Group:
		return sql_query.NewSQLSelectBuilder[new_model.Group]("test").Build()
	case new_model.Category:
		return sql_query.NewSQLSelectBuilder[new_model.Category]("test").Build()
	case new_model.DamageReport:
		return sql_query.NewSQLSelectBuilder[new_model.DamageReport]("test").Build()
	case new_model.MissingReport:
		return sql_query.NewSQLSelectBuilder[new_model.MissingReport]("test").Build()
	case new_model.DisposalReport:
		return sql_query.NewSQLSelectBuilder[new_model.DisposalReport]("test").Build()
	case new_model.BorrowingRequest:
		return sql_query.NewSQLSelectBuilder[new_model.BorrowingRequest]("test").Build()
	case new_model.AssignmentRequest:
		return sql_query.NewSQLSelectBuilder[new_model.AssignmentRequest]("test").Build()
	case new_model.TransferRequest:
		return sql_query.NewSQLSelectBuilder[new_model.TransferRequest]("test").Build()
	case new_model.Attachment:
		return sql_query.NewSQLSelectBuilder[new_model.Attachment]("test").Build()
	}

	return "", nil, nil
}

func callInsertBuilder(v any) (string, []interface{}, error) {
	switch v.(type) {
	case new_model.User:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.User{}).Build()
	case new_model.UserSystemRole:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.UserSystemRole{}).Build()
	case new_model.UserTransactionRole:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.UserTransactionRole{}).Build()
	case new_model.CustomField:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.CustomField{}).Build()
	case new_model.CustomFieldCategory:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.CustomFieldCategory{}).Build()
	case new_model.Asset:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Asset{}).Build()
	case new_model.AssetName:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.AssetName{}).Build()
	case new_model.Brand:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Brand{}).Build()
	case new_model.Model:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Model{}).Build()
	case new_model.Group:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Group{}).Build()
	case new_model.Category:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Category{}).Build()
	case new_model.DamageReport:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.DamageReport{}).Build()
	case new_model.MissingReport:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.MissingReport{}).Build()
	case new_model.DisposalReport:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.DisposalReport{}).Build()
	case new_model.BorrowingRequest:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.BorrowingRequest{}).Build()
	case new_model.AssignmentRequest:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.AssignmentRequest{}).Build()
	case new_model.TransferRequest:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.TransferRequest{}).Build()
	case new_model.Attachment:
		return sql_query.NewSQLInsertBuilder("test").Insert(new_model.Attachment{}).Build()
	}

	return "", nil, nil
}

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

type productDTO struct {
	ID    int     `json:"id" column:"id"`
	Name  string  `json:"name" column:"name"`
	Price float64 `json:"price" column:"price"`
}

type orderDTO struct {
	ID        int    `json:"id" column:"id"`
	UserID    int    `json:"userId" column:"user_id"`
	ProductID int    `json:"productId" column:"product_id"`
	Status    string `json:"status" column:"status"`
}

// MemoryUsage
func memUsageKB() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc / 1024
}

func TestCacheMemoryUsage(t *testing.T) {
	runtime.GC() // biar baseline bersih
	before := memUsageKB()

	sql_query.ExtractJSONTags[userDTO]()
	sql_query.ExtractJSONTags[productDTO]()
	sql_query.ExtractJSONTags[orderDTO]()

	runtime.GC() // pastikan semua yg bisa di-GC udah dibuang
	after := memUsageKB()

	fmt.Printf("Cache retained memory = %d KB (for 3 dummy structs)\n", after-before)
}

func TestCacheMemoryUsage_NewModel(t *testing.T) {
	runtime.GC()
	before := memUsageKB()

	for _, v := range allStructs {
		callSelectBuilder(v)
	}

	runtime.GC()
	after := memUsageKB()

	fmt.Printf("Cache retained memory = %d KB (for %d structs)\n", after-before, len(allStructs))
}

// --- Benchmarks ---

func BenchmarkExtractJSONTags(b *testing.B) {
	for i, v := range allStructs {
		name := fmt.Sprintf("%T", v)

		b.Run(name+fmt.Sprintf("%s %d", "/ColdCache", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sql_query.ClearCache()
				sql_query.ExtractJSONTags[any](reflect.TypeOf(v))
			}
		})

		sql_query.ExtractJSONTags[any](reflect.TypeOf(v))
		b.Run(name+fmt.Sprintf("%s %d", "/WarmCache", i+1), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// for _, v := range allStructs {
				sql_query.ExtractJSONTags[any](reflect.TypeOf(v))
				// }
			}
		})
	}
}

// func BenchmarkMakeMapJSONTagsFromValue(b *testing.B) {
// 	for i, v := range allStructs {
// 		name := fmt.Sprintf("%T", v)

// 		b.Run(name+fmt.Sprintf("%s %d", "/ColdCache", i+1), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				sql_query.ClearCache()
// 				sql_query.MakeMapJSONTagsFromValue(reflect.TypeOf(v))
// 			}
// 		})

// 		sql_query.MakeMapJSONTagsFromValue(reflect.TypeOf(v))
// 		b.Run(name+fmt.Sprintf("%s %d", "/WarmCache", i+1), func(b *testing.B) {
// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				// for _, v := range allStructs {
// 				sql_query.MakeMapJSONTagsFromValue(reflect.TypeOf(v))
// 				// }
// 			}
// 		})
// 	}
// }

func BenchmarkNewSQLSelectBuilder(b *testing.B) {
	for i, v := range allStructs {
		name := fmt.Sprintf("%T", v)

		b.Run(name+fmt.Sprintf("%s %d", "/ColdCache", i+1), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sql_query.ClearCache() // clear cache every iteration
				callSelectBuilder(v)
			}
		})

		b.Run(name+fmt.Sprintf("%s %d", "/WarmCache", i+1), func(b *testing.B) {
			// warm up once
			callSelectBuilder(v)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				// for _, v := range allStructs {
				callSelectBuilder(v)
				// }
			}
		})
	}
}

// func BenchmarkNewSQLInsertBuilder(b *testing.B) {
// 	for i, v := range allStructs {
// 		name := fmt.Sprintf("%T", v)

// 		b.Run(name+fmt.Sprintf("%s %d", "/ColdCache", i+1), func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				sql_query.ClearCache() // clear cache every iteration
// 				callInsertBuilder(v)
// 			}
// 		})

// 		b.Run(name+fmt.Sprintf("%s %d", "/WarmCache", i+1), func(b *testing.B) {
// 			// warm up once
// 			callInsertBuilder(v)

// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				// for _, v := range allStructs {
// 				callInsertBuilder(v)
// 				// }
// 			}
// 		})
// 	}
// }

// // Benchmark SelectOne (always returns 1 row)
// func BenchmarkSelectOne(b *testing.B) {
// 	for _, v := range allStructs {
// 		name := fmt.Sprintf("%T", v)
// 		ctx := context.Background()
// 		u := reflect.New(reflect.TypeOf(v)).Interface()

// 		b.Run(name+"/ColdCache", func(b *testing.B) {
// 			for i := 0; i < b.N; i++ {
// 				query, args, _ := callSelectBuilder(v) // rebuild every iteration (cold)
// 				svc := &BasePostgreSqlService{
// 					Pool:      newFakePool(1),
// 					TableName: "users",
// 				}
// 				_ = svc.SelectOne(u, ctx, query, args...)
// 			}
// 		})

// 		b.Run(name+"/WarmCache", func(b *testing.B) {
// 			// build once, then reuse
// 			query, args, _ := callSelectBuilder(v)
// 			svc := &BasePostgreSqlService{
// 				Pool:      newFakePool(1),
// 				TableName: "users",
// 			}

// 			b.ResetTimer() // ensure timer starts after cache is ready
// 			for i := 0; i < b.N; i++ {
// 				_ = svc.SelectOne(u, ctx, query, args...)
// 			}
// 		})
// 	}
// }

// // Benchmark SelectMany with various row counts
// func BenchmarkSelectMany(b *testing.B) {
// 	for _, v := range allStructs {
// 		ctx := context.Background()
// 		sliceType := reflect.SliceOf(reflect.TypeOf(v))
// 		slicePtr := reflect.New(sliceType).Interface()
// 		b.Run(fmt.Sprintf("%T/ColdCache", v), func(b *testing.B) {
// 			for _, n := range []int{1, 10, 100, 1000} {
// 				b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
// 					svc := &BasePostgreSqlService{
// 						Pool:      newFakePool(n),
// 						TableName: "users",
// 					}

// 					b.ResetTimer()
// 					for i := 0; i < b.N; i++ {
// 						// ColdCache → builder dipanggil setiap iterasi
// 						query, args, _ := callSelectBuilder(v)
// 						_ = svc.SelectMany(slicePtr, ctx, query, args...)
// 					}
// 				})
// 			}
// 		})

// 		b.Run(fmt.Sprintf("%T/WarmCache", v), func(b *testing.B) {
// 			query, args, _ := callSelectBuilder(v) // WarmCache → builder dipanggil sekali

// 			for _, n := range []int{1, 10, 100, 1000} {
// 				b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
// 					svc := &BasePostgreSqlService{
// 						Pool:      newFakePool(n),
// 						TableName: "users",
// 					}

// 					b.ResetTimer()
// 					for i := 0; i < b.N; i++ {
// 						_ = svc.SelectMany(slicePtr, ctx, query, args...)
// 					}
// 				})
// 			}
// 		})
// 	}
// }

// // Benchmark InsertOne (always returns 1 row)
// func BenchmarkInsertOne(b *testing.B) {
// 	u := userDTO{
// 		ID:    1,
// 		Name:  "A",
// 		Email: "a",
// 	}
// 	query, args, _ := sql_query.NewSQLInsertBuilder("users").Insert(u).Build()
// 	svc := &BasePostgreSqlService{
// 		Pool:      newFakePool(1),
// 		TableName: "users",
// 	}
// 	ctx := context.Background()
// 	for i := 0; i < b.N; i++ {
// 		_, _ = svc.InsertOne(ctx, query, args...)
// 	}
// }

// // Benchmark InsertMany with various row counts
// func BenchmarkInsertMany(b *testing.B) {
// 	batches := map[int][]userDTO{}
// 	for _, n := range []int{1, 10, 100, 1000} {
// 		slice := make([]userDTO, n)
// 		for i := 0; i < n; i++ {
// 			slice[i] = userDTO{ID: i, Name: "A", Email: "a"}
// 		}
// 		batches[n] = slice
// 	}
// 	for n, slice := range batches {
// 		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
// 			// Build once per sub-benchmark
// 			query, args, _ := sql_query.NewSQLInsertBuilder("users").
// 				Insert(slice).
// 				Build()

// 			svc := &BasePostgreSqlService{
// 				Pool:      newFakePool(n),
// 				TableName: "users",
// 			}
// 			ctx := context.Background()

// 			b.ResetTimer()
// 			for i := 0; i < b.N; i++ {
// 				_, _ = svc.InsertMany(ctx, query, args...)
// 			}
// 		})
// 	}
// }
