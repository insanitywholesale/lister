package postgres

var createListTableQuery = `
CREATE TABLE if not exists List (
	Id SERIAL PRIMARY KEY,
	Title VARCHAR(100),
	Items TEXT
);`

var listRetrieveAllQuery = `SELECT * FROM List`
var listRetrievalQuery = `SELECT * FROM List WHERE Id=$1;`
var listInsertQuery = `INSERT INTO List`
