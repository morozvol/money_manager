-- PROCEDURE: public.apply_operation(integer, double precision, integer)

DROP PROCEDURE IF EXISTS public.apply_operation(integer, double precision, integer);

CREATE OR REPLACE PROCEDURE public.apply_operation(
    IN id_account integer,
    IN sum double precision,
    IN id_category integer)
    LANGUAGE 'plpgsql'
AS $BODY$
declare
    bal float := 0;
    category_type int;
begin

    SELECT type
    into category_type
    from category
    where id = id_category;

    SELECT
        CAST(balance + iif(category_type = 1,sum,sum * -1) as float)
    into bal
    from account
    where id = id_account;

    IF bal < 0 THEN
        RAISE 'На счету недостаточно средств для проведения операции' USING ERRCODE = '23505';
    ELSE

        INSERT INTO operation(time, sum, id_category, id_account) VALUES
            (NOW(), sum, id_category, id_account);

        UPDATE account
        set  balance = balance + iif(category_type =1,sum,-sum)
        where id = id_account /*RETURNING id INTO id_operation*/;
        commit;

    END IF;

end;
$BODY$;

ALTER PROCEDURE public.apply_operation(integer, double precision, integer)
    OWNER TO postgres;
