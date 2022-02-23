

-- +migrate Up


CREATE TABLE public.account (
	id bigserial NOT NULL,
	user_name varchar(255) NULL,
	account_name varchar(255) NULL,
	password  varchar(255) NULL,
	balance bigint default 0,
	last_update_by varchar(255) NULL,
	last_update timestamptz NULL,
	PRIMARY KEY (id)
);


-- +migrate Down