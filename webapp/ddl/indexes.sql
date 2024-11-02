ALTER TABLE comments ADD INDEX post_id_idx (post_id);
ALTER TABLE posts ADD INDEX created_at_idx (created_at);
