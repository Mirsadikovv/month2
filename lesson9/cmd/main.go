// package main

// import (
// 	"database/sql"
// 	"log"
// 	"psql-connection/config"
// 	"psql-connection/postgres"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	cfg := config.Load()
// 	postgresDB, err := postgres.NewPostgres(cfg)
// 	if err != nil {
// 		log.Panic("", err)
// 	}
// 	defer postgresDB.Close()

// }

// func createEmployee(db *sql.DB) {
// 	query := `INSERT INTO employee (id,first_name,last_name, department_name)
// 	VALUES($1,$2,$3,$4)`
// 	res, err := db.Exec(query)

// }

// ///////////////

package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"psql-connection/config"
// 	"psql-connection/model"
// 	"psql-connection/postgres"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	cfg := config.Load()

// 	postgresDB, err := postgres.NewPostgres(cfg)
// 	if err != nil {
// 		log.Panic("Failed to connect to postgres", err)
// 	}

// 	defer postgresDB.Close()

// newID, err := uuid.NewV4()
// if err != nil {
// 	log.Fatal("Failed to create new uuid", err)
// }

// employee := model.Employee{
// 	Id:             newID.String(),
// 	FirstName:      "Do' 9876",
// 	LastName:       "Last Name 1",
// 	DepartmentName: "economic 1",
// }
// ch <- employee

// // CREATE
// go func(ch chan model.Employee) {
// 	employee := ch
// 	employeeID := createEmployee(postgresDB, employee)
// 	ch <- employeeID
// }()
// //SELECT
// go func(ch chan string) {
// 	employeeID := <-ch
// 	employeeData := selectEmplyee(postgresDB, employeeID)

// 	fmt.Println("data >>>>>>>>>>>", employeeData.FirstName)
// }()
//UPDATE

// go func() {
// 	employeeData.FirstName = "Do' 6789"
// 	err = updateEmployee(postgresDB, employeeData)
// 	if err != nil {
// 		log.Fatal("updateEmployee failed: ", err)
// 	}
// }()

// employeeData = selectEmplyee(postgresDB, employeeID)

// fmt.Println("employeeData.FirstName >>>>>>>>>>>", employeeData.FirstName)

// SELECT ALL DATA
// 	employeesData := getAllEmployees(postgresDB)
// 	for _, employee := range employeesData {
// 		fmt.Println("\n", employee.FirstName)
// 	}

// }

// func createEmployee(db *sql.DB, emp model.Employee) string {
// 	query := `
// 		INSERT INTO employee (id, first_name, last_name, department_name)
// 		VALUES ($1, $2, $3, $4) RETURNING id`

// 	var id string

// 	res := db.QueryRow(query, emp.Id, emp.FirstName, emp.LastName, emp.DepartmentName).Scan(&id)
// 	// if err != nil {
// 	// 	log.Fatal("error creating employee :", err)
// 	// }

// 	fmt.Println("result :", res)

// 	return id
// }

// func selectEmplyee(db *sql.DB, id string) model.Employee {
// 	var emplyee model.Employee

// 	query := `SELECT * FROM employee WHERE id = $1`

// 	row := db.QueryRow(query, id)

// 	row.Scan(&emplyee.Id, &emplyee.FirstName,
// 		&emplyee.LastName,
// 		&emplyee.DepartmentName,
// 		&emplyee.CreatedAt)

// 	return emplyee
// }

// func getAllEmployees(db *sql.DB) []model.Employee {
// 	rows, err := db.Query(`SELECT * FROM employee`)

// 	if err != nil {
// 		log.Fatal("Error getting all employees", err)
// 	}

// 	defer rows.Close()

// 	var employees []model.Employee

// 	for rows.Next() {
// 		var employee model.Employee

// 		err := rows.Scan
// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"psql-connection/config"
// 	"psql-connection/model"
// 	"psql-connection/postgres"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	cfg := config.Load()

// 	postgresDB, err := postgres.NewPostgres(cfg)
// 	if err != nil {
// 		log.Panic("Failed to connect to postgres", err)
// 	}

// 	defer postgresDB.Close()(&employee.Id, &employee.FirstName,
// 			&employee.LastName,
// 			&employee.DepartmentName,
// 			&employee.CreatedAt)

// 		if err != nil {
// 			log.Fatal("Error getting employee", err)
// 		}

// 		employees = append(employees, employee)

// 	}

// 	return employees
// }

// func updateEmployee(db *sql.DB, employee model.Employee) error {
// 	query := `
// 		UPDATE employee
// 		SET first_name = $1,
// 			last_name = $2,
// 			department_name = $3
// 		WHERE id = $4
// 	`

// 	res, err := db.Exec(query,
// 		employee.FirstName,
// 		employee.LastName,
// 		employee.DepartmentName,
// 		employee.Id,
// 	)
// 	if err != nil {
// 		log.Fatal("Error updating employee", err)
// 		return err
// 	}

// 	count, err := res.RowsAffected()
// 	if err != nil {
// 		log.Fatal("Error  finding rows affected", err)
// 		return err
// 	}

// 	fmt.Println("count >>>>", count)

// 	return nil
// }
