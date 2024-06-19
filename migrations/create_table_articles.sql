CREATE TABLE IF NOT EXISTS articles (
    id BIGSERIAL PRIMARY KEY
    , user_id BIGINT NOT NULL
    , category_id BIGINT NOT NULL
    , content TEXT NOT NULL
    , title VARCHAR(50) NOT NULL
    , tags TEXT NULL
    , page TEXT NULL
    , description TEXT NULL
    , created_by BIGINT NOT NULL
    , created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
    , updated_by BIGINT NULL
    , updated_at TIMESTAMP WITH TIME ZONE NULL
    , CONSTRAINT users_id_user_id FOREIGN KEY (user_id) REFERENCES users (id)
    , CONSTRAINT users_id_created_by FOREIGN KEY (created_by) REFERENCES users (id)
    , CONSTRAINT users_id_updated_by FOREIGN KEY (updated_by) REFERENCES users (id)
    , CONSTRAINT categories_id_category_id FOREIGN KEY (category_id) REFERENCES categories (id)
)