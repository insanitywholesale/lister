package postgres

var createListTableQuery = `
CREATE TABLE if not exists List (
	Id SERIAL PRIMARY KEY,
	Title VARCHAR,
	Items TEXT[]
);`

var listRetrieveAllQuery = `SELECT Id, Title, Items FROM List;`
var listRetrievalQuery = `SELECT Id, Title, Items FROM List WHERE Id=$1;`
var listInsertQuery = `INSERT INTO List (
	Title,
	Items
) VALUES ($1, $2) RETURNING Id;`
var listDeleteQuery = `DELETE FROM List WhERE Id=$1;`
