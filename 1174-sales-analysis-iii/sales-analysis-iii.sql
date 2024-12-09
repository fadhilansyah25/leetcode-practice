/* Write your T-SQL query statement below */
WITH SalesProduct AS (
    SELECT
    s.*,
    p.product_name
    FROM Sales s
    JOIN Product p ON s.product_id =  p.product_id
)

SELECT 
product_id,
product_name
FROM SalesProduct
GROUP BY product_id, product_name
HAVING MIN(sale_date) >= '2019-01-01' AND MAX(sale_date) <= '2019-03-31'