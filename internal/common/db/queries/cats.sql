
-- name: GetCat :one
SELECT * FROM cats
WHERE id = $1 LIMIT 1;

-- name: ListCats :many
SELECT * FROM cats
ORDER BY created_at;

-- name: CreateCat :exec
INSERT INTO cats (
  id, created_at, name, years_of_experience, breed, salary
) VALUES (
  $1, $2, $3, $4, $5, $6
);

-- name: UpdateCat :exec
UPDATE cats 
SET name = $2, years_of_experience = $3, breed = $4, salary = $5
WHERE id = $1;

-- name: DeleteCat :exec
DELETE FROM cats
WHERE id = $1;
