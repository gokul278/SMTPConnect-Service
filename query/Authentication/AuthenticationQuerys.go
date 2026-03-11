package authenticationquery

var VerifyEmailPhoneNoSQL = `
SELECT uc."refUserId", U."refRTId", UA."refUAPass"
FROM userdomain."UsersCommunication" uc
         JOIN auth."UsersAuth" UA on uc."refUserId" = UA."refUserId"
         JOIN public."Users" U ON U."refUserId" = uc."refUserId"
WHERE uc."refUCMail" = $1
  AND U."refUserStatus" = TRUE;
`

var SignUpSQL = `
WITH email_check AS (
    SELECT 1
    FROM userdomain."UsersCommunication"
    WHERE "refUCMail" = $3
),
-- 1. Calculate the ID separately
next_id_calc AS (
    SELECT (10000 + COALESCE(MAX("refUserId"), 1)) AS next_id
    FROM public."Users"
),
ins AS (
    INSERT INTO public."Users" ("refUserCustId", "refRTId", "refUserName")
    -- 2. Select from the calculated ID, and apply the filter here
    SELECT next_id, $1, $2
    FROM next_id_calc
    WHERE NOT EXISTS (SELECT 1 FROM email_check)
    RETURNING "refUserId"
),
insC AS (
    INSERT INTO userdomain."UsersCommunication" ("refUserId", "refUCMail")
    SELECT "refUserId", $3
    FROM ins
),
insA AS (
    INSERT INTO auth."UsersAuth" ("refUserId", "refUAPass")
    SELECT "refUserId", $4
    FROM ins
),
insAu AS (
    INSERT INTO  audit."TransHistory" ("transTypeId", "refTHData", "refTHTime", "refUserId", "refTHActionBy")
    SELECT 2, 'Account Created', $5,"refUserId", "refUserId"
    FROM ins
)

SELECT
    CASE 
        WHEN EXISTS (SELECT 1 FROM ins) THEN true 
        ELSE false 
    END AS status,

    CASE 
        WHEN EXISTS (SELECT 1 FROM ins) THEN 'Account Created'
        ELSE 'Email Already Exists'
    END AS message,

    CASE 
        WHEN EXISTS (SELECT 1 FROM ins) THEN 201
        ELSE 409
    END AS statuscode;
`
