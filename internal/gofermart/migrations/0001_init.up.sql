BEGIN TRANSACTION;

DO $$
BEGIN
   
    CREATE TABLE IF NOT EXISTS users (
            id INT GENERATED ALWAYS AS IDENTITY,
            ulogin TEXT NOT NULL,
            upass TEXT NOT NULL,
            PRIMARY KEY(id),
            UNIQUE(ulogin, upass)
    );

    CREATE INDEX IF NOT EXISTS ulogin_id ON users USING hash(ulogin);
END $$;
--
--
COMMIT TRANSACTION;