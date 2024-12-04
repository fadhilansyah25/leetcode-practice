/* Write your T-SQL query statement below */
DELETE Person WHERE id IN (
    SELECT id
    FROM (
        SELECT 
            id,
            email,
            ROW_NUMBER() OVER (PARTITION BY email ORDER BY id ASC) AS row
        FROM Person
    ) AS p
    WHERE p.row > 1
)