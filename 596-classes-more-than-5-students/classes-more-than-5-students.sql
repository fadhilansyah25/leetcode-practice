/* Write your T-SQL query statement below */
WITH ClassAgg AS (
    SELECT
    class,
    COUNT(*) AS total_appeared
    FROM Courses
    GROUP BY class
)

SELECT class
FROM ClassAgg
WHERE total_appeared >= 5