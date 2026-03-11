package userquery

var GetUserProfileQuery = `
SELECT u."refUserId",
       u."refUserName",
       u."refRTId",
       (RT."refRTPrefix" || u."refUserCustId") AS "refUserCustId"
FROM public."Users" u
         JOIN public."RoleType" RT on u."refRTId" = RT."refRTId"
WHERE u."refUserId" = $1
`