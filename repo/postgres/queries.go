package postgres

var createListTableQuery = `
CREATE TABLE if not exists List (
	ListId SERIAL PRIMARY KEY,
	Title VARCHAR,
	Items TEXT[]
);`

var listRetrieveAllQuery = `SELECT * FROM List`
var listRetrievalQuery = `SELECT * FROM List WHERE Id=$1;`
var listInsertQuery = `INSERT INTO List (
	Title,
	Items
) VALUES ($1, $2) RETURNING ListId;`
