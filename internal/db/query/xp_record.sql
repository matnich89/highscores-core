-- name: CreateXpRecord :one
INSERT INTO xp_records (
    skill,
    character_id,
    last_xp_count,
    last_level,
    last_position,
    updated_at,
    created_at
) VALUES (
      $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: GetXpRecord :one
SELECT * FROM xp_records
WHERE skill = $1 AND character_id = $2;

-- name: UpdateXpRecord :one
UPDATE xp_records
SET last_xp_count = $3,
    last_level = $4,
    last_position = $5,
    updated_at = $6
WHERE character_id = $1 AND skill = $2
RETURNING *;
