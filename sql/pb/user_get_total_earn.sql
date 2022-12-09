CREATE OR REPLACE FUNCTION user_get_total_earn(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugte.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT ub.total_earned
			FROM users_base ub
			WHERE ub.tele_id = _tid
		) ag
	) ugte
	INTO _response;

	RETURN _response;
END;
$function$