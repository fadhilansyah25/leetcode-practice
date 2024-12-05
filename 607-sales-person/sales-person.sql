/* Write your T-SQL query statement below */
SELECT 
    SP.name
FROM SalesPerson SP
WHERE SP.sales_id NOT IN (
    SELECT 
    S.sales_id
    FROM SalesPerson S
    JOIN Orders O ON O.sales_id = SP.sales_id 
    JOIN Company C ON O.com_id = C.com_id
    WHERE C.name = 'RED'
) 