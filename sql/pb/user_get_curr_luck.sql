CREATE OR REPLACE FUNCTION user_get_curr_luck(_tid BIGINT)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugcl.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT ub.curr_luck
			FROM users_base ub
			WHERE ub.tele_id = _tid
		) ag
	) ugcl
	INTO _response;

	RETURN _response;
END;
$function$