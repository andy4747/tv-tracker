-- name: ListTokens :many
SELECT *
FROM tokens;

-- name: GetToken :one
SELECT *
FROM tokens
WHERE id = $1;

-- name: GetTokenByUser :one
SELECT *
FROM tokens
WHERE user_id = $1;

-- name: GetTokenByToken :one
SELECT *
FROM tokens
WHERE token = $1;

-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1;

-- name: CreateToken :one
INSERT INTO tokens (created_at, token, user_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateToken :one
UPDATE tokens
SET updated_at = $1, token = $2, user_id = $3
WHERE id = $4
RETURNING *;
