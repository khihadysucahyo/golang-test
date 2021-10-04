SELECT u.id, u.username, u2.username as parentUsername
FROM users u 
LEFT JOIN users u2 ON u.parent = u2.id 
ORDER BY u.id;