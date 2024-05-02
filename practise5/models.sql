--22222222222222222222222222222222222222222222

create table Employees (
	id INT PRIMARY KEY,
	name VARCHAR(50)
);
insert into Employees (id, name) values (1, 'Reinaldos');
insert into Employees (id, name) values (2, 'Antony');
insert into Employees (id, name) values (3, 'Kylie');
insert into Employees (id, name) values (4, 'Chalmers');
insert into Employees (id, name) values (5, 'Ive');



create table EmployeeUNI (
	id INT PRIMARY KEY,
	unique_id INT
);
insert into EmployeeUNI (id, unique_id) values (1, 4);
insert into EmployeeUNI (id, unique_id) values (2, 3);
insert into EmployeeUNI (id, unique_id) values (3, 1);


SELECT u.unique_id, e.name FROM Employees e LEFT JOIN EmployeeUNI u on e.id = u.id;

--333333333333333333333333333333333333333333333333

