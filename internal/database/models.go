// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Book struct {
	ID          string
	Title       sql.NullString
	Author      sql.NullString
	Filesize    sql.NullString
	Extension   sql.NullString
	Md5         sql.NullString
	Year        sql.NullString
	Language    sql.NullString
	Pages       sql.NullString
	Publisher   sql.NullString
	Edition     sql.NullString
	Coverurl    sql.NullString
	Downloadurl sql.NullString
	Pageurl     sql.NullString
}
