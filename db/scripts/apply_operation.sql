-- PROCEDURE: public.apply_operation(integer, double precision, integer)

DROP PROCEDURE IF EXISTS public.apply_operation(integer, double precision, integer);

CREATE OR REPLACE PROCEDURE public.apply_operation(
    IN id_account integer,
    IN sum double precision,
    IN type integer)
    LANGUAGE 'plpgsql'
AS $BODY$
declare
    bal float := 0;
begin

    SELECT
        CAST(balance + iif(type=1,sum,sum * -1) as float)
    into bal
    from account
    where id = id_account;

    IF bal < 0 THEN
        RAISE 'На счету недостаточно средств для проведения операции' USING ERRCODE = '23505';
    ELSE

        INSERT INTO operation(time, sum, type, id_account) VALUES
            (NOW(), sum, type, id_account);

        UPDATE account
        set  balance = balance + iif(type=1,sum,-sum)
        where id = id_account /*RETURNING id INTO id_operation*/;
        commit;

    END IF;

end;
$BODY$;

ALTER PROCEDURE public.apply_operation(integer, double precision, integer)
    OWNER TO postgres;
