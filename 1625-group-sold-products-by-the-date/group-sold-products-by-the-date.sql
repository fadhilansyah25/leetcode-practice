/* Write your T-SQL query statement below */
WITH grouped AS (
    SELECT
        sell_date,
        product
    FROM Activities
    GROUP BY sell_date, product
)
SELECT 
    sell_date,
    COUNT(product) AS num_sold,
    STRING_AGG(product, ',') AS products
FROM grouped
GROUP BY sell_date;