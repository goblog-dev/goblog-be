CREATE TABLE IF NOT EXISTS categories (
    id BIGSERIAL PRIMARY KEY
    , name VARCHAR(50) NOT NULL UNIQUE
    , created_by BIGINT NOT NULL
    , created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    , updated_by BIGINT NULL
    , updated_at TIMESTAMP WITH TIME ZONE NULL
    , CONSTRAINT users_id_created_by FOREIGN KEY (created_by) REFERENCES users (id)
    , CONSTRAINT users_id_updated_by FOREIGN KEY (updated_by) REFERENCES users (id)
)