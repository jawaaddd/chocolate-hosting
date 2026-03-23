CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE,
    phone_number VARCHAR(20) UNIQUE,
    otp INT,
    otp_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE game_versions (
    id VARCHAR(50) PRIMARY KEY,
    version_type VARCHAR(20),
    java_version INT NOT NULL,
    download_url TEXT NOT NULL
    server_hash VARCHAR(40) NOT NULL
);

CREATE TABLE compute_instances (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ip VARCHAR(45) NOT NULL,
    current_servers INT DEFAULT 0,
    current_players INT DEFAULT 0,
    max_servers INT NOT NULL,
    max_players INT NOT NULL
);

CREATE TABLE servers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    version_id VARCHAR(50) NOT NULL REFERENCES game_versions(id),
    compute_id UUID NOT NULL REFERENCES compute_instances(id),
    compute_port INT NOT NULL,
    server_name VARCHAR(255) UNIQUE NOT NULL,
    max_players INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);