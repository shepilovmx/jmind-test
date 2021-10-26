CREATE TABLE IF NOT EXISTS public.block_totals
(
    id            serial PRIMARY KEY,
    block_number  INT UNIQUE NOT NULL,
    transactions  INT NOT NULL,
    amount        FLOAT NOT NULL
);