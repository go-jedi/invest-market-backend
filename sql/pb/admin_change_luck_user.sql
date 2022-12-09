CREATE OR REPLACE FUNCTION admin_change_luck_user(_tid BIGINT, _nlk character varying)
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

	UPDATE users_base SET curr_luck = _nlk WHERE tele_id = _tid;
END;
$function$