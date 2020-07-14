package schema

import (
	"context"
	"database/sql"
	d3pgx "github.com/godzie44/d3/adapter/pgx"
	"github.com/godzie44/d3/orm"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestTypeConversion(t *testing.T) {
	pgDb, d3orm := initDb(t)

	ctx := d3orm.CtxWithSession(context.Background())
	rep, err := d3orm.MakeRepository(&allTypeStruct{})
	assert.NoError(t, err)

	uuidVal, _ := uuid.FromString("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")

	currTime := time.Unix(time.Now().Unix(), 0)
	entity := &allTypeStruct{
		Uuid:             uuidVal,
		BoolField:        true,
		IntField:         1,
		Int32Field:       2,
		Int64Field:       3,
		Float32Field:     4,
		Float64Field:     5,
		StringField:      "d3",
		TimeField:        currTime,
		NullBoolField:    sql.NullBool{Bool: true, Valid: true},
		NullI32Field:     sql.NullInt32{Int32: 1, Valid: true},
		NullI64Field:     sql.NullInt64{Int64: 2, Valid: true},
		NullFloat64Field: sql.NullFloat64{Float64: 1.1, Valid: true},
		NullStringField:  sql.NullString{Valid: false},
		NullTimeField:    sql.NullTime{Time: currTime, Valid: true},
	}

	assert.NoError(t, rep.Persists(ctx, entity))
	assert.NoError(t, orm.Session(ctx).Flush())

	ctx2 := d3orm.CtxWithSession(context.Background())
	rep, err = d3orm.MakeRepository(&allTypeStruct{})
	assert.NoError(t, err)

	fetchedEntity, err := rep.FindOne(ctx2, rep.Select().Where("id", "=", entity.ID))
	assert.NoError(t, err)

	assert.Equal(t, entity, fetchedEntity)
	dropDb(t, pgDb)
}

func initDb(t *testing.T) (*pgx.Conn, *orm.Orm) {
	cfg, _ := pgx.ParseConfig(os.Getenv("D3_PG_TEST_DB"))
	driver, err := d3pgx.NewPgxDriver(cfg)
	assert.NoError(t, err)

	d3orm := orm.New(driver)
	assert.NoError(t, d3orm.Register((*allTypeStruct)(nil), (*entityWithAliases)(nil)))

	sqlSchema, err := d3orm.GenerateSchema()
	assert.NoError(t, err)

	_, err = driver.UnwrapConn().(*pgx.Conn).Exec(context.Background(), sqlSchema)
	assert.NoError(t, err)

	return driver.UnwrapConn().(*pgx.Conn), d3orm
}

func dropDb(t *testing.T, db *pgx.Conn) {
	_, err := db.Exec(context.Background(), `
DROP TABLE IF EXISTS all_types;
DROP TABLE IF EXISTS test_aliases;
`)
	assert.NoError(t, err)
}

func TestCustomTypeConversion(t *testing.T) {
	pgDb, d3orm := initDb(t)

	ctx := d3orm.CtxWithSession(context.Background())
	rep, err := d3orm.MakeRepository(&entityWithAliases{})
	assert.NoError(t, err)

	entity := &entityWithAliases{
		email:       Email("mail"),
		secretEmail: myEmail("mail"),
	}

	assert.NoError(t, rep.Persists(ctx, entity))
	assert.NoError(t, orm.Session(ctx).Flush())

	ctx2 := d3orm.CtxWithSession(context.Background())
	rep, err = d3orm.MakeRepository(&entityWithAliases{})
	assert.NoError(t, err)

	fetchedEntity, err := rep.FindOne(ctx2, rep.Select().Where("id", "=", entity.ID))
	assert.NoError(t, err)

	assert.Equal(t, entity, fetchedEntity)
	dropDb(t, pgDb)
}
