-- +migrate Up

-- +migrate StatementBegin
CREATE SEQUENCE IF NOT EXISTS users_pkey_seq;
CREATE TABLE "users"
(
    id         BIGINT NOT NULL DEFAULT nextval('users_pkey_seq'::regclass),
    email      VARCHAR(30),
    username   VARCHAR(10),
    password   VARCHAR(256),
    name       VARCHAR(50),
    status     CHAR            DEFAULT '1',
    created_by BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_by BIGINT,
    deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted    BOOLEAN         DEFAULT FALSE,
    CONSTRAINT pk_users_id PRIMARY KEY (id)
);

INSERT INTO users
    (name, email, username, password)
VALUES ('user satu', 'user1@mail.com', 'user1', 'VGVzdGluZzEyMyE=');

CREATE SEQUENCE IF NOT EXISTS books_pkey_seq;
CREATE TABLE "books"
(
    id         BIGINT NOT NULL DEFAULT nextval('books_pkey_seq'::regclass),
    name       VARCHAR(50),
    quantity   INT,
    created_by BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_by BIGINT,
    deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted    BOOLEAN         DEFAULT FALSE,
    CONSTRAINT pk_books_id PRIMARY KEY (id)
);

CREATE SEQUENCE IF NOT EXISTS loans_pkey_seq;
CREATE TABLE "loans"
(
    id         BIGINT NOT NULL DEFAULT nextval('loans_pkey_seq'::regclass),
    user_id    BIGINT,
    book_id    BIGINT,
    quantity   INT,
    date_start DATE,
    date_end   DATE,
    created_by BIGINT,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_by BIGINT,
    updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_by BIGINT,
    deleted_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted    BOOLEAN         DEFAULT FALSE,
    CONSTRAINT pk_loans_id PRIMARY KEY (id),
    CONSTRAINT fk_user_loan FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_book_loan FOREIGN KEY (book_id) REFERENCES books (id)
);
-- +migrate StatementEnd