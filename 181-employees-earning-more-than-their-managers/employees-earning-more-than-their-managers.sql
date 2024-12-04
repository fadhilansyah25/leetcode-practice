/* Write your T-SQL query statement below */

SELECT 
    e.name as Employee
    -- , e.salary,
    -- e.managerId,
    -- m.name as managerName,
    -- m.salary as managerSalary
FROM Employee e
LEFT JOIN Employee m ON e.managerId = m.id
WHERE e.salary > m.salary
