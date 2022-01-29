-- Table: public.PhoneNumbers

-- DROP TABLE public."PhoneNumbers";
CREATE SEQUENCE IF NOT EXISTS "PhoneNumbers_id_seq";
CREATE TABLE IF NOT EXISTS public."PhoneNumbers"
(
    "number" character varying(50) COLLATE pg_catalog."default" NOT NULL,
    id integer NOT NULL DEFAULT nextval('"PhoneNumbers_id_seq"'::regclass),
    CONSTRAINT "PhoneNumbers_pkey" PRIMARY KEY (id)
);

-- TABLESPACE pg_default;

ALTER TABLE public."PhoneNumbers"
    OWNER to postgres;