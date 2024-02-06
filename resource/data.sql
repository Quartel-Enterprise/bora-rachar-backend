-- popular tabelas
INSERT INTO BORA_RACHAR.`user` (id,user_id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    ('1','user01','arrighi','arrigh@gmail.com','398999','2024-01-15 12:00:00',NULL,NULL);


INSERT INTO BORA_RACHAR.`group`(id, name, avatar, access_code, created_by, created_at, updated_at, deleted_at)
VALUES('1', 'teste', '', '', 'user01', '2024-01-15 12:00:00', NULL, NULL);
