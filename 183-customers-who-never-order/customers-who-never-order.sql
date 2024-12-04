/* Write your T-SQL query statement below */
SELECT 
    name AS Customers
FROM Customers
WHERE id NOT IN(
    SELECT 
        c.id
    FROM Customers c
    JOIN Orders o ON o.customerId = c.id
)