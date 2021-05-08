CREATE USER app;
ALTER USER app WITH encrypted password 'apppassword';
CREATE DATABASE test_migrate;
GRANT ALL PRIVILEGES ON DATABASE test_migrate TO app;

