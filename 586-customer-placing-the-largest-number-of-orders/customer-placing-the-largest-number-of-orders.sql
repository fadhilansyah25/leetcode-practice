/* Write your T-SQL query statement below */

WITH customerTotalOrder AS (
    SELECT 
        customer_number,
        COUNT(*) AS total_order
    FROM Orders
    GROUP BY customer_number
)

SELECT 
    customer_number
FROM customerTotalOrder
WHERE total_order = (SELECT MAX(total_order) FROM customerTotalOrder)