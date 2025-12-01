package provider

import (
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/db"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
)

// DBProviderer provides a way to create database service for various supported databases.
//
// Can be extended to create new support for other database.
type DBProviderer interface {
	Init()
	MakeMongoDBService(dbName db.DBName, collection string) service.MongoDBServicer
	MakePostgreSQLService(dbName db.DBName, options ...db.PostgreSQLServiceOptions) service.PostgreSQLServicer
	MakeRedisService(options ...db.RedisServiceOptions) service.RedisServicer
	Close()
}

// DBProvider implements DBProviderer
type DBProvider struct{}

func (dp *DBProvider) Init() {
	// Pass
}

// MakeMongoDBService creates a mongodb service.
func (dp *DBProvider) MakeMongoDBService(dbName db.DBName, collection string) service.MongoDBServicer {
	return service.MakeMongoDBService(dbName, collection)
}

// MakePostgreSQLService creates a postgresql service
func (dp *DBProvider) MakePostgreSQLService(dbName db.DBName, options ...db.PostgreSQLServiceOptions) service.PostgreSQLServicer {
	return service.MakePostgreSQLService(dbName, options...)
}

// MakeRedisService creates a redis service
func (dp *DBProvider) MakeRedisService(options ...db.RedisServiceOptions) service.RedisServicer {
	return service.MakeRedisService(options...)
}

func (dp *DBProvider) Close() {
	// Pass
}
