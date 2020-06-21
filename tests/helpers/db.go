package helpers

import (
	"context"
	"fmt"
	"github.com/godzie44/d3/orm"
	"github.com/godzie44/d3/orm/persistence"
	"github.com/godzie44/d3/orm/query"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DbAdapterWithQueryCounter struct {
	queryCounter, insertCounter, updateCounter, deleteCounter int
	dbAdapter                                                 orm.Driver
}

func (d *DbAdapterWithQueryCounter) MakeScalarDataMapper() orm.ScalarDataMapper {
	return d.dbAdapter.MakeScalarDataMapper()
}

func (d *DbAdapterWithQueryCounter) MakePusher(tx orm.Transaction) persistence.Pusher {
	ps := d.dbAdapter.MakePusher(tx)
	return &persistStoreWithCounters{
		ps: ps,
		onInsert: func() {
			d.insertCounter++
		},
		onUpdate: func() {
			d.updateCounter++
		},
		onDelete: func() {
			d.deleteCounter++
		},
	}
}

func (d *DbAdapterWithQueryCounter) BeginTx() (orm.Transaction, error) {
	return d.dbAdapter.BeginTx()
}

func NewDbAdapterWithQueryCounter(dbAdapter orm.Driver) *DbAdapterWithQueryCounter {
	wrappedAdapter := &DbAdapterWithQueryCounter{dbAdapter: dbAdapter}

	dbAdapter.AfterQuery(func(_ string, _ ...interface{}) {
		wrappedAdapter.queryCounter++
	})

	return wrappedAdapter
}

func (d *DbAdapterWithQueryCounter) ExecuteQuery(query *query.Query) ([]map[string]interface{}, error) {
	return d.dbAdapter.ExecuteQuery(query)
}

func (d *DbAdapterWithQueryCounter) BeforeQuery(fn func(query string, args ...interface{})) {
	d.dbAdapter.BeforeQuery(fn)
}

func (d *DbAdapterWithQueryCounter) AfterQuery(fn func(query string, args ...interface{})) {
	d.dbAdapter.AfterQuery(fn)
}

func (d *DbAdapterWithQueryCounter) QueryCounter() int {
	return d.queryCounter
}

func (d *DbAdapterWithQueryCounter) UpdateCounter() int {
	return d.updateCounter
}

func (d *DbAdapterWithQueryCounter) InsertCounter() int {
	return d.insertCounter
}

func (d *DbAdapterWithQueryCounter) DeleteCounter() int {
	return d.deleteCounter
}
func (d *DbAdapterWithQueryCounter) ResetCounters() {
	d.deleteCounter = 0
	d.updateCounter = 0
	d.insertCounter = 0
	d.queryCounter = 0
}

type persistStoreWithCounters struct {
	ps                           persistence.Pusher
	onInsert, onUpdate, onDelete func()
}

func (p *persistStoreWithCounters) Insert(table string, cols []string, values []interface{}, onConflict persistence.OnConflict) error {
	p.onInsert()
	return p.ps.Insert(table, cols, values, onConflict)
}

func (p *persistStoreWithCounters) InsertWithReturn(table string, cols []string, values []interface{}, returnCols []string, withReturned func(scanner persistence.Scanner) error) error {
	p.onInsert()
	return p.ps.InsertWithReturn(table, cols, values, returnCols, withReturned)
}

func (p *persistStoreWithCounters) Update(table string, cols []string, values []interface{}, identityCond map[string]interface{}) error {
	p.onUpdate()
	return p.ps.Update(table, cols, values, identityCond)
}

func (p *persistStoreWithCounters) Remove(table string, identityCond map[string]interface{}) error {
	p.onDelete()
	return p.ps.Remove(table, identityCond)
}

type pgTester struct {
	t    *testing.T
	conn *pgx.Conn
}

func NewPgTester(t *testing.T, conn *pgx.Conn) *pgTester {
	return &pgTester{t, conn}
}

func (p *pgTester) SeeOne(sql string, args ...interface{}) *pgTester {
	return p.See(1, sql, args...)
}

func (p *pgTester) SeeTwo(sql string, args ...interface{}) *pgTester {
	return p.See(2, sql, args...)
}

func (p *pgTester) SeeThree(sql string, args ...interface{}) *pgTester {
	return p.See(3, sql, args...)
}

func (p *pgTester) SeeFour(sql string, args ...interface{}) *pgTester {
	return p.See(4, sql, args...)
}

func (p *pgTester) See(count int, sql string, args ...interface{}) *pgTester {
	var cnt int
	err := p.conn.QueryRow(context.Background(), fmt.Sprintf("SELECT count(*) cnt FROM (%s) t", sql), args...).Scan(&cnt)
	assert.NoError(p.t, err)

	assert.Equal(p.t, count, cnt)
	return p
}

func (p *pgTester) SeeTable(tableName string) *pgTester {
	var tableSql = "SELECT * FROM pg_tables where schemaname = 'public' and tablename=$1"

	var cnt int
	err := p.conn.QueryRow(context.Background(), fmt.Sprintf("SELECT count(*) cnt FROM (%s) t", tableSql), tableName).Scan(&cnt)
	assert.NoError(p.t, err)

	assert.GreaterOrEqual(p.t, cnt, 1)
	return p
}
