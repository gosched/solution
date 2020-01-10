SELECT
    IFNULL (
        (SELECT DISTINCT Salary FROM Employee
        ORDER BY Salary DESC LIMIT 1, 1),
        NULL
    ) AS SecondHighestSalary;

SELECT
    IFNULL(
        (SELECT DISTINCT Salary
        FROM Employee
        ORDER BY Salary DESC
        LIMIT 1 OFFSET 1),
        NULL) AS SecondHighestSalary;

SELECT MAX(Salary) as SecondHighestSalary FROM Employee
WHERE Salary < (SELECT MAX(Salary) FROM Employee);