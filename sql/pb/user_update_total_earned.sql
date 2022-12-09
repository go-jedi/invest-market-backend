CREATE OR REPLACE FUNCTION user_update_total_earned(_tid BIGINT, _em NUMERIC)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ub users_base;
BEGIN
	SELECT *
	FROM users_base
	WHERE tele_id = _tid
	INTO _ub;

	IF _ub.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users_base SET total_earned = _em WHERE tele_id = _tid;
END;
$function$