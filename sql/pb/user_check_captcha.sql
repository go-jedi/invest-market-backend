CREATE OR REPLACE FUNCTION user_check_captcha(_vur BIGINT)
	RETURNS BOOLEAN
	LANGUAGE plpgsql
AS $function$
DECLARE
   _c captcha;
BEGIN
	SELECT *
	FROM captcha
	WHERE value = _vur
	INTO _c;

	IF _c.id ISNULL THEN
		RETURN FALSE;
	ELSE
		RETURN TRUE;
	END IF;
END;
$function$