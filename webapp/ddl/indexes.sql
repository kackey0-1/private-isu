ALTER TABLE comments ADD INDEX post_id_idx (post_id);
ALTER TABLE comments ADD INDEX user_id_idx (user_id);
ALTER TABLE posts ADD INDEX created_at_idx (created_at);
ALTER TABLE users ADD INDEX del_flg_idx (del_flg);