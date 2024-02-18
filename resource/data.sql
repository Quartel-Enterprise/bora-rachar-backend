-- popular tabelas
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (1,'arrighi','arrigh@gmail.com','398999','2024-01-15 12:00:00',NULL,NULL);
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (2,'fabricio','fabricio@gmail.com','1234','2024-01-15 12:00:00',NULL,NULL);
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (3,'pierre','pierre@gmail.com','23543465','2024-01-15 12:00:00',NULL,NULL);


INSERT INTO BORA_RACHAR.`group`(id, name, avatar, access_code, created_by, created_at, updated_at, deleted_at)
VALUES('1', 'teste', '', '', 1, '2024-01-15 12:00:00', NULL, NULL);


INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(1, 1, 1, 0x30, '2006-01-02 15:04:05', NULL);
INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(2, 2, 1, 0x30, '2006-01-02 15:04:05', NULL);
INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(3, 2, 1, 0x31, '2006-01-02 15:04:05', NULL);