-- name: CreateTransform :one
INSERT INTO transforms (
  owner,
  long_url,
  short_url
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransform :one
SELECT * FROM transforms
WHERE id = $1 LIMIT 1;

-- name: GetTransformForUpdate :one
SELECT * FROM transforms
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransforms :many
SELECT * FROM transforms
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTransform :one
UPDATE transforms
SET short_url = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransform :exec
DELETE FROM transforms
WHERE id = $1;