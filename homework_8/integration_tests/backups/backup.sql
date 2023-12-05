CREATE TABLE public.articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title CHARACTER VARYING(512),
    body CHARACTER VARYING(2048),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE public.articles OWNER TO postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO public.articles (id, title, body) VALUES ('6912354f-43b4-4106-8744-d84471adf59b', 'for delete', 'some body');
INSERT INTO public.articles (id, title, body) VALUES ('a462db9b-b7ae-434c-87af-943d080d5c00', 'for update', 'some body');