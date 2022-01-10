-- name: CreateItem :one
INSERT INTO items (name, value)
VALUES (@name, @value)
RETURNING id;

-- name: UpdateItem :exec
UPDATE items SET name = @name, value = @value
WHERE id = @id;

-- name: GetItemById :one
SELECT *
FROM items
WHERE id = @id;

-- name: GetItemByName :one
SELECT *
FROM items
WHERE name = @name;

-- name: CheckItemExists :one
SELECT id
FROM items
WHERE name = @name;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = @id;

-- name: GetItems :many
SELECT *
FROM items;