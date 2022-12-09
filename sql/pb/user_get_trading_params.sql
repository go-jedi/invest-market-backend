CREATE OR REPLACE FUNCTION user_get_trading_params(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugtp.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT ut.tele_id, ut.choose_asset, ut.choose_price, ut.investment_price, ut.movement_price, ut.waiting_time_sec
			FROM users_trading ut
			WHERE ut.tele_id = _tid
		) ag
	) ugtp
	INTO _response;

	RETURN _response;
END;
$function$