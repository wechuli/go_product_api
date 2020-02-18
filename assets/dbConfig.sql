CREATE DATABASE productdb;
CREATE USER goprojectuser WITH PASSWORD 'password';

ALTER ROLE goprojectuser SET client_encoding TO 'utf8';
ALTER ROLE goprojectuser SET default_transaction_isolation TO 'read committed';
ALTER ROLE goprojectuser SET timezone TO 'UTC';


GRANT ALL PRIVILEGES ON DATABASE productdb TO goprojectuser;