-- name: FindBooksByOffset :many
SELECT
	id,
	title,
	author,
	created_at,
	updated_at
FROM
	books
ORDER BY
	id DESC
LIMIT
	$1
OFFSET
	$2
;

-- name: CountBooks :one
SELECT
	COUNT(id) AS total
FROM
	books
;

-- name: FindLastBooks :many
SELECT
	id,
	title,
	author,
	created_at,
	updated_at
FROM
	books
ORDER BY
	id DESC
LIMIT
	$1
;

-- name: FindBooksByID :many
SELECT
	id,
	title,
	author,
	created_at,
	updated_at
FROM
	books
WHERE
	id <= $1 and updated_at <=
ORDER BY
	id DESC
LIMIT
	$2
;