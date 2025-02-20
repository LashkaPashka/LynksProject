-- Table for Postgres
CREATE TABLE links (
    id SERIAL PRIMARY KEY, -- первичный ключ
    url VARCHAR(20) NOT NULL DEFAULT '',
    hash VARCHAR(20) NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
	
-- Table for mySQL --
CREATE TABLE stats (
    id SERIAL PRIMARY KEY, -- первичный ключ
    url VARCHAR(20) NOT NULL DEFAULT '',
    click INTEGER NOT NULL DEFAULT '',
    average_length INTEGER DEFAULT CURRENT_TIMESTAMP
)
