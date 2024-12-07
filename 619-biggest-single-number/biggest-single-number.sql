/* Write your T-SQL query statement below */
SELECT COALESCE(
    (SELECT TOP 1
        num
    FROM MyNumbers
    GROUP BY num
    HAVING COUNT(*) = 1
    ORDER BY num DESC
    ),
    NULL
) AS num