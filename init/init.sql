CREATE TABLE user
(
    id         CHAR(26)     NOT NULL COMMENT 'ユーザーID(ulid)',
    name       VARCHAR(20)  NOT NULL COMMENT 'ユーザー名',
    email      VARCHAR(254) NOT NULL COMMENT 'メールアドレス',
    password   CHAR(60)     NOT NULL COMMENT 'パスワード',
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
);
