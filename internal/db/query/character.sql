-- name: CreateCharacter :one
INSERT INTO characters (
    name,
    last_check,
    created_at
) VALUES (
   $1, $2, $3
)
RETURNING *;

-- name: GetCharacterByName :one
SELECT * FROM characters
WHERE name = $1;

-- name: GetCharacter :one
SELECT * FROM characters
WHERE id = $1;

-- name: UpdateCharacter :exec
UPDATE characters
SET last_check = $2
WHERE  id = $1;
