-- name: AddServices :many 
SELECT * FROM services limit $1 offset $2;
