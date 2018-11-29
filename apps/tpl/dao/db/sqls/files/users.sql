-- name: user-insert

INSERT INTO users (
 user_name
,user_token
,user_secret
,user_expiry
,user_email
,user_avatar
,user_active
,user_synced
,user_admin
,user_hash
) VALUES (?,?,?,?,?,?,?,?,?,?)

-- name: user-find

SELECT
 user_id
,user_name
,user_token
,user_secret
,user_expiry
,user_email
,user_avatar
,user_active
,user_synced
,user_admin
,user_hash
FROM users
ORDER BY user_name ASC

-- name: user-find-id

SELECT
 user_id
,user_name
,user_token
,user_secret
,user_expiry
,user_email
,user_avatar
,user_active
,user_synced
,user_admin
,user_hash
FROM users
WHERE user_id = ?
LIMIT 1

-- name: user-find-login

SELECT
 user_id
,user_name
,user_token
,user_secret
,user_expiry
,user_email
,user_avatar
,user_active
,user_synced
,user_admin
,user_hash
FROM users
WHERE user_name = ?
LIMIT 1

-- name: user-update

UPDATE users
SET
,user_token  = ?
,user_secret = ?
,user_expiry = ?
,user_email  = ?
,user_avatar = ?
,user_active = ?
,user_synced = ?
,user_admin  = ?
,user_hash   = ?
WHERE user_id = ?

-- name: user-delete

DELETE FROM users WHERE user_id = ?
