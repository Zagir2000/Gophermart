BEGIN TRANSACTION;

ALTER TABLE users RENAME TO '__users';

COMMIT;