CREATE OR REPLACE FUNCTION user_update_trading_asset(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ut users_trading;
BEGIN
	SELECT *
	FROM users_trading
	WHERE tele_id = (js->>'tele_id')::BIGINT
	INTO _ut;

	IF _ut.id ISNULL THEN
		RAISE EXCEPTION 'пользователь не найден';
	END IF;

	UPDATE users_trading SET choose_asset = js->>'choose_asset', choose_price = (js->>'choose_price')::NUMERIC WHERE tele_id = (js->>'tele_id')::BIGINT;
END;
$function$