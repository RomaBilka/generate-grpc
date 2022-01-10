BEGIN;

CREATE TABLE IF NOT EXISTS public.items
(
    id serial NOT NULL,
    name text NOT NULL,
    value text NOT NULL,
    PRIMARY KEY (id)
    );

ALTER TABLE IF EXISTS public.items
    OWNER to postgres;

COMMIT;