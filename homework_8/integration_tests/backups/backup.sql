CREATE TABLE public.articles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title CHARACTER VARYING(512),
    body CHARACTER VARYING(2048),
    created_at TIMESTAMP DEFAULT current_timestamp,
    updated_at TIMESTAMP DEFAULT current_timestamp
);

ALTER TABLE public.articles OWNER TO postgres;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

COPY public.articles (id, title, body, created_at, updated_at) FROM stdin;
4ddc8d46-f08f-43da-b227-3afd79c69d16	for update	some body	2023-12-05 14:30:00	2023-12-05 14:30:00
6912354f-43b4-4106-8744-d84471adf59b	for delete	some body	2023-12-05 14:30:00	2023-12-05 14:30:00
\.