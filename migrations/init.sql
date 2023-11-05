CREATE USER spark WITH PASSWORD 'dev123456';
CREATE DATABASE spark;

\c spark
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS public.t_pessoa (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    data_criacao TIMESTAMPTZ NOT NULL DEFAULT now(),
    data_atualizacao TIMESTAMPTZ,
    nome TEXT NOT NULL,
    apelido TEXT NOT NULL
);

GRANT SELECT, UPDATE, INSERT, DELETE ON public.t_pessoa TO spark;

CREATE TABLE IF NOT EXISTS public.t_tecnologia (
    id UUID NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    data_criacao TIMESTAMPTZ NOT NULL DEFAULT now(),
    data_atualizacao TIMESTAMPTZ,
    pessoa_id UUID REFERENCES public.t_pessoa(id),
    nome TEXT
);

GRANT SELECT, UPDATE, INSERT, DELETE ON public.t_tecnologia TO spark;
GRANT ALL PRIVILEGES ON DATABASE spark TO spark;
