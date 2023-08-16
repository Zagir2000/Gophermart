BEGIN TRANSACTION;

DO $$
BEGIN
       
    CREATE TABLE IF NOT EXISTS orders (
            id INT GENERATED ALWAYS AS IDENTITY,
            ordernumber TEXT NOT NULL,
            userlogin TEXT NOT NULL,
            orderdate DATE NOT NULL,  
            PRIMARY KEY(id),
            UNIQUE(ordernumber, userlogin)
    );

    CREATE INDEX IF NOT EXISTS userlogin_id ON orders USING hash(userlogin);
    CREATE INDEX IF NOT EXISTS orderdate_id ON orders USING hash(orderdate);
END $$;

--
--
COMMIT TRANSACTION;