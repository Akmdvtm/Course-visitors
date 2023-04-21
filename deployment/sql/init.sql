BEGIN;

CREATE TYPE user_role AS ENUM ('user','admin');

CREATE TABLE IF NOT EXISTS users
(
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR                  NOT NULL,
    discord_name VARCHAR UNIQUE                    DEFAULT '',
    telegram_id  VARCHAR UNIQUE                    DEFAULT '',
    role         user_role                NOT NULL DEFAULT 'user',
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

COMMIT;