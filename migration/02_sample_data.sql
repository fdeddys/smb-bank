-- +migrate Up

INSERT INTO public.account
(user_name, "password", balance, last_update_by, last_update, account_name)
VALUES( 'deddy', '0C91A43F8E1EC5FCBA28F8A5A34532679305CA131302AD2A06218B47F30CED88', 1000, 'system', NOW(), 'Deddy');

-- +migrate Down