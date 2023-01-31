CREATE TABLE IF NOT EXISTS public.contact(
    id bigserial NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    address text NOT NULL,
    phone_number bigint NOT NULL,
    created_time timestamp with time zone NOT NULL DEFAULT now(),
    updated_time timestamp with time zone NOT NULL DEFAULT now(),
    CONSTRAINT contact_pkey PRIMARY KEY (id),
    CONSTRAINT contact_phone_number_key UNIQUE (phone_number)
);
CREATE INDEX contact_index_first_name on contact (first_name);
CREATE INDEX contact_index_last_name on contact (last_name);
CREATE INDEX contact_index_phone_number on contact (phone_number);