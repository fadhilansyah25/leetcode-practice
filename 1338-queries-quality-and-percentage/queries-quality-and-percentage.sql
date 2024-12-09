/* Write your T-SQL query statement below */
SELECT
    query_name,
    ROUND(AVG(CAST(rating AS FLOAT)/CAST(position AS FLOAT)), 2) AS quality,
    ROUND(100.0 * SUM(IIF(rating<3, 1, 0)) / COUNT(*), 2) AS poor_query_percentage
FROM Queries
GROUP BY query_name
ORDER BY query_name