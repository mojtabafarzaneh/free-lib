-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE books(
    ID varchar(256) PRIMARY KEY,
    title varchar(256) ,
    Author  varchar(256),
	Filesize  varchar(256),
	Extension  varchar(256),
	Md5  varchar(256),
	Year    varchar(256),
	Language  varchar(256),
	Pages     varchar(256),
	Publisher  varchar(256),
	Edition    varchar(256),
	CoverURL  varchar(256),
	DownloadURL varchar(256),
	PageURL varchar(256)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE books;