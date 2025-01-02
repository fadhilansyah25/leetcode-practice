/* Write your T-SQL query statement below */
SELECT 
        P.product_name, 
        Agg_O.unit
FROM Products P
INNER JOIN (
    SELECT product_id, SUM(Unit) AS unit
    FROM Orders
    WHERE YEAR(order_date) = 2020 AND DATEPART(MONTH, order_date) = 02
    GROUP BY product_id
    HAVING SUM(unit) >= 100
) AS Agg_O ON P.product_id = Agg_O.product_id;