DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_roles WHERE rolname = 'meebin_user') THEN
        CREATE USER meebin_user WITH PASSWORD 'password';
    END IF;
END $$;

DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM pg_database WHERE datname = 'meebin_db') THEN
        CREATE DATABASE meebin_db;
    END IF;
END $$;

GRANT ALL PRIVILEGES ON DATABASE meebin_db TO meebin_user;