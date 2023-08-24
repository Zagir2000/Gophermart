BEGIN TRANSACTION;

DO $$
BEGIN
       
    CREATE TABLE IF NOT EXISTS goods (
            id INT GENERATED ALWAYS AS IDENTITY,
            descriptionorder TEXT,
            price BIGINT,
            reward BIGINT,
            reward_type TEXT,
            PRIMARY KEY(id),
            UNIQUE(descriptionorder)
    );

    CREATE INDEX IF NOT EXISTS descriptionorder_id ON goods USING hash(descriptionorder);
END $$;

--
--
COMMIT TRANSACTION;