CREATE TABLE bar (id serial not null);
CREATE TABLE foo (id serial not null, bar serial);

-- name: TableName :one
SELECT foo.id
FROM foo
JOIN bar ON foo.bar = bar.id
WHERE bar.id = $1 AND foo.id = $2;
