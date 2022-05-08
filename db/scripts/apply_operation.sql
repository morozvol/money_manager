-- PROCEDURE: public.apply_operation(integer, double precision, integer)

DROP FUNCTION IF EXISTS public.apply_operation(integer, double precision, integer, text, timestamp);

CREATE OR REPLACE FUNCTION public.apply_operation(
    p_id_account integer,
    p_sum double precision,
    p_id_category integer,
    p_description text,
    p_time timestamp)
    RETURNS int
    LANGUAGE 'plpgsql'
AS
$BODY$
declare
    bal           float := 0;
    category_type int;
    p_id          int;
    p_id_user     int;
begin

    SELECT type
    into category_type
    from category
    where id = p_id_category;

    SELECT CAST(balance + iif(category_type = 1, p_sum, p_sum * -1) as float)
    into bal
    from account
    where id = p_id_account;

    SELECT id_user
    into p_id_user
    from account
    where id = p_id_account;

    IF bal < 0 THEN
        RAISE 'На счету недостаточно средств для проведения операции' USING ERRCODE = '23505';
    ELSE

        INSERT INTO operation(time, sum, id_category, id_account, description, id_user)
        VALUES (p_time, p_sum, p_id_category, p_id_account, p_description, p_id_user)
        RETURNING id INTO p_id;

        UPDATE account
        SET balance = balance + iif(category_type = 1, p_sum, -p_sum)
        WHERE id = p_id_account;

        RETURN p_id;

    END IF;

end;
$BODY$;

ALTER FUNCTION public.apply_operation(integer, double precision, integer, text, timestamp)
    OWNER TO postgres;