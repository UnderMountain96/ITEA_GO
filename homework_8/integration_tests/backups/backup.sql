CREATE TABLE public.articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title CHARACTER VARYING(512),
    body CHARACTER VARYING(2048),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE public.articles OWNER TO postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
