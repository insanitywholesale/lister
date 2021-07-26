package cassandra

var createListerKeyspaceQuery = `
CREATE KEYSPACE lister WITH replication = {
	'class': 'SimpleStrategy',
	'replication_factor' : 1};`

var createListTableQuery = `
CREATE TABLE IF NOT EXISTS lister.List (
	Id INT PRIMARY KEY,
	Title TEXT,
	Items list<TEXT>
);`

var listRetrieveAllQuery = `SELECT Id, Title, Items FROM lister.List;`
var listRetrievalQuery = `SELECT Id, Title, Items FROM lister.List WHERE Id=? LIMIT 1;`
var listInsertQuery = `INSERT INTO lister.List (
	Id,
	Title,
	Items
) VALUES (?, ?, ?);`
var listMaxIdQuery = `SELECT MAX(id) FROM lister.List WHERE id > 0 ALLOW FILTERING;`
var listDeleteQuery = `DELETE FROM lister.List WHERE Id=?;`
