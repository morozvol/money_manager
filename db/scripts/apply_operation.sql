-- PROCEDURE: public.apply_operation(integer, double precision, integer)

DROP PROCEDURE IF EXISTS public.apply_operation(integer, double precision, integer, text);

CREATE OR REPLACE PROCEDURE public.apply_operation(
    IN p_id_account integer,
    IN p_sum double precision,
    IN p_id_category integer,
    In p_description text)
    LANGUAGE 'plpgsql'
AS $BODY$
declare
    bal float := 0;
    category_type int;
begin

    SELECT type
    into category_type
    from category
    where id = p_id_category;

    SELECT
        CAST(balance + iif(category_type = 1,p_sum,p_sum * -1) as float)
    into bal
    from account
    where id = p_id_account;

    IF bal < 0 THEN
        RAISE 'На счету недостаточно средств для проведения операции' USING ERRCODE = '23505';
    ELSE

        INSERT INTO operation(time, sum, id_category, id_account, description) VALUES
            (NOW(), p_sum, p_id_category, p_id_account, p_description);

        UPDATE account
        set  balance = balance + iif(category_type =1,p_sum,-p_sum)
        where id = p_id_account /*RETURNING id INTO id_operation*/;
       -- commit;

    END IF;

end;
$BODY$;

ALTER PROCEDURE public.apply_operation(integer, double precision, integer, text)
    OWNER TO postgres;
