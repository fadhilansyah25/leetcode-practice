/* Write your T-SQL query statement below */
SELECT 
    player_id,
    event_date AS first_login
FROM
    (
        SELECT 
            *,
            ROW_NUMBER() OVER (PARTITION BY player_id ORDER BY event_date) AS row
        FROM Activity
    ) AS act
WHERE act.row = 1;