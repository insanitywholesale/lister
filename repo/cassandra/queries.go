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

var listRetrieveAllQuery = `SELECT Id, Title, Items FROM List;`
var listRetrievalQuery = `SELECT Id, Title, Items FROM List WHERE Id=? LIMIT 1;`
var listInsertQuery = `INSERT INTO List (
	Id,
	Title,
	Items
) VALUES (?, ?, ?);`
var listMaxIdQuery = `SELECT MAX(id) FROM List WHERE id > 0 ALLOW FILTERING;`
