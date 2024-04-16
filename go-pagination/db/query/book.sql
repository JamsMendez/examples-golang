-- name: FindBooksByOffset :many
SELECT id,
       title,
       author,
       created_at,
       updated_at
FROM books
ORDER BY id DESC
LIMIT $1 OFFSET $2;

-- name: CountBooks :one
SELECT COUNT(id) AS total
FROM books;