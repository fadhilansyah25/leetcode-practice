/* Write your T-SQL query statement below */
SELECT 
    email
    -- ,
    -- COUNT(*)
FROM Person
GROUP BY email
HAVING COUNT(*) > 1