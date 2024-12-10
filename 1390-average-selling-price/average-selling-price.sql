/* Write your T-SQL query statement below */
SELECT 
    P.product_id,
    COALESCE(ROUND(SUM(units * price) / SUM(CAST(units AS FLOAT)), 2), 0) AS average_price
FROM
UnitsSold US
RIGHT JOIN Prices P ON P.product_id = US.product_id AND US.purchase_date BETWEEN P.start_date AND P.end_date
GROUP BY P.product_id