xo ${CONNSTR} -N -M -B -T PersonById -o models/ << ENDSQL
SELECT
				name::text AS name,
				email::text AS email
FROM person
WHERE
				id = %%personId uuid.UUID%%
ENDSQL
