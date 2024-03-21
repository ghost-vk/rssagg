-- name: CreatePost :one
INSERT INTO posts(
  id, 
  created_at, 
  updated_at, 
  title, 
  description,
  published_at,
  url, 
  feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAllPosts :many
SELECT * FROM posts 
ORDER BY created_at DESC 
LIMIT $1 
OFFSET $2;

-- name: CheckIfPostWithUrlExists :one
SELECT EXISTS(SELECT 1 FROM posts WHERE url=$1);

-- name: GetPostsForUser :many
SELECT p.* FROM posts p
  LEFT JOIN feed_follows ff ON ff.feed_id=p.feed_id  
  LEFT JOIN users u ON ff.user_id=u.id
WHERE u.api_key = $1
ORDER BY p.published_at DESC
LIMIT $2
OFFSET $3;
