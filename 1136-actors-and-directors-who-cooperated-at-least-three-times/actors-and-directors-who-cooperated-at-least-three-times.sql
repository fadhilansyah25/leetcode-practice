/* Write your T-SQL query statement below */
SELECT
    actor_id, director_id
FROM
(
SELECT 
    actor_id, 
    director_id,
    COUNT(*) AS total_pairing
FROM ActorDirector
GROUP BY actor_id, director_id
HAVING COUNT(*) >=3 
) AS ActorDirectorAgg