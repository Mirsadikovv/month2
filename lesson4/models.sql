-- Active: 1712133551411@@127.0.0.1@5432@postgres@public
CREATE TABLE employees(  
    employee_id int NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    department_id int
    );
CREATE TABLE departments(
    department_id int PRIMARY KEY,
    department_name VARCHAR(25)
    );

CREATE TABLE salaries(
    employee_id int, 
    salary_amount int,
    effective_date DATE
    );

INSERT INTO departments VALUES(1,'administrative'),(2,'marketing'),(3,'sales'),(4,'development');

INSERT INTO employees VALUES(1,'Brandon',4),(2,'Tom',3),(3,'Anna',2),(4,'Lana',1);

INSERT INTO salaries VALUES(1,2000,'2023-01-03'),(1,2000,'2023-02-03'),(1,2000,'2023-03-03');

INSERT INTO salaries VALUES(2,1200,'2023-01-03'),(2,1400,'2023-02-03'),(2,1400,'2023-03-03');

INSERT INTO salaries VALUES(3,900,'2023-01-03'),(3,1100,'2023-02-03'),(3,1300,'2023-03-03');

INSERT INTO salaries VALUES(4,500,'2023-01-03'),(4,1000,'2023-01-03'),(4,1100,'2023-02-03'),(4,1500,'2023-03-03');


SELECT e.name,department_name,s.salary_amount 
FROM employees e 
LEFT JOIN departments d ON e.department_id = d.department_id 
LEFT JOIN salaries s ON e.employee_id = s.employee_id WHERE EXTRACT(MONTH FROM s.effective_date) = EXTRACT(MONTH FROM CURRENT_DATE)-1;


SELECT d.department_name, avg(s.salary_amount) AS avg_salary 
FROM employees e 
JOIN departments d ON e.department_id = d.department_id
JOIN salaries s ON e.employee_id = s.employee_id 
GROUP BY d.department_id
ORDER BY avg_salary DESC LIMIT 1;


SELECT e.employee_id, e.name
FROM employees e
JOIN salaries s ON e.employee_id = s.employee_id
GROUP BY e.employee_id, e.name
HAVING COUNT(DISTINCT s.salary_amount) = 1;


SELECT e.employee_id, e.name, SUM(s.salary_amount) AS total_salary
FROM employees e
JOIN salaries s ON e.employee_id = s.employee_id
GROUP BY e.employee_id, e.name
ORDER BY total_salary DESC
LIMIT 1;
