-- popular tabelas
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (1,'arrighi','arrigh@gmail.com','398999','2024-01-15 12:00:00',NULL,NULL);
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (2,'fabricio','fabricio@gmail.com','1234','2024-01-15 12:00:00',NULL,NULL);
INSERT INTO BORA_RACHAR.`user` (id,name,email,pix_key,created_at,updated_at,deleted_at) VALUES
    (3,'pierre','pierre@gmail.com','23543465','2024-01-15 12:00:00',NULL,NULL);


INSERT INTO BORA_RACHAR.`group`(id, name, avatar, access_code, created_by, created_at, updated_at, deleted_at)
VALUES('1', 'teste', '', '', 1, '2024-01-15 12:00:00', NULL, NULL);
INSERT INTO BORA_RACHAR.`group`(id, name, avatar, access_code, created_by, created_at, updated_at, deleted_at)
VALUES(2, 'Escalada', 'https://imagem', '', 1, '2006-01-02 15:04:05', NULL, NULL);


INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(1, 1, 1, 0x30, '2006-01-02 15:04:05', NULL);
INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(2, 2, 1, 0x30, '2006-01-02 15:04:05', NULL);
INSERT INTO BORA_RACHAR.group_participant(id, user_id, group_id, is_admin, admission_date, removed_at)
VALUES(3, 2, 1, 0x31, '2006-01-02 15:04:05', NULL);


INSERT INTO BORA_RACHAR.expense(id, group_id, title, description, category, avatar, value, expense_date, created_by, created_at, payer, updated_at, deleted_at)
VALUES(1, NULL, 'Role pizzaria', '', 'FOOD', 'link photo', 20.0000, '2024-02-21 01:54:43', '1', '2006-01-02 15:04:05', 2, NULL, NULL);
INSERT INTO BORA_RACHAR.expense_payment_split(id, user_id, expense_id, value, created_at, is_debt_settled)
VALUES(1, 1, 1, 15.5000, '2006-01-02 15:04:05', 0);
INSERT INTO BORA_RACHAR.expense_payment_split(id, user_id, expense_id, value, created_at, is_debt_settled)
VALUES(2, 2, 1, 4.5000, '2006-01-02 15:04:05', 0);

INSERT INTO BORA_RACHAR.expense(id, group_id, title, description, category, avatar, value, expense_date, created_by, created_at, payer, updated_at, deleted_at)
VALUES(2, 2, 'Uber volta Escalada', 'descricao detalhada da despesa', 'TRANSPORTATION', 'link photo', 15.0000, '2024-02-21 01:59:37', '2', '2006-01-02 15:04:05', 1, NULL, NULL);
INSERT INTO BORA_RACHAR.expense_payment_split(id, user_id, expense_id, value, created_at, is_debt_settled)
VALUES(3, 3, 2, 10.0000, '2006-01-02 15:04:05', 0);