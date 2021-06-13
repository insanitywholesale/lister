package cassandra

var createListerKeyspaceQuery = `
CREATE KEYSPACE lister WITH replication = {
	'class': 'SimpleStrategy',
	'replication_factor' : 1};`

var createListTableQuery = `
CREATE TABLE lister.List (
	Id INT PRIMARY KEY,
	Title TEXT,
	Items list<TEXT>
);`

var listRetrieveAllQuery = `SELECT id, title, items FROM List;`
var listRetrievalQuery = `SELECT id, title, items FROM List WHERE Id=? LIMIT 1;`
var listInsertQuery = `INSERT INTO List (
	Id,
	Title,
	Items
) VALUES (?, ?, ?);`
