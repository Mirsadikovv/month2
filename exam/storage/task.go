package storage

import (
	"context"
	"encoding/json"
	"exam-project/model"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTask(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}

		var task model.TodoList

		err = json.Unmarshal(resBody, &task)
		if err != nil {
			fmt.Println("error while unmarshaling data", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		query := `
		INSERT INTO "todo_list" ("task_id", "task_name", "task_time","task_comment","task_priority","user_id")
		VALUES ($1, $2, $3,$4, $5, $6) RETURNING task_id`

		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println("error when creating uuid", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		fmt.Println(task.UserId)
		res, err := db.Exec(ctx, query, id.String(), task.TaskName, task.TaskTime, task.TaskComment, task.TaskPriority, task.UserId)
		if err != nil {
			fmt.Println("error when creating task", err)

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("result :", res)
		w.WriteHeader(http.StatusOK)
	}
}

func UpdateTask(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}

		var task model.TodoList

		err = json.Unmarshal(resBody, &task)
		if err != nil {
			fmt.Println("error while unmarshaling data", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		query := `
		UPDATE "todo_list" SET "task_name" = $1, "task_time" = $2 WHERE "task_id" = $3
	    RETURNING task_id`

		res, err := db.Exec(ctx, query, task.TaskName, task.TaskTime, task.TaskId)
		if err != nil {
			fmt.Println("error when updating task (update) ", err)

			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Println("result :", res)
		w.WriteHeader(http.StatusOK)
	}
}

func SelectTask(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}

		var task model.TodoList

		err = json.Unmarshal(resBody, &task)
		if err != nil {
			fmt.Println("error while unmarshaling data", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tasks := []TodoList{}
		query := `
		SELECT * FROM "todo_list"
		WHERE "user_id"=$1`

		////////////////

		rows, err := db.Query(ctx, query, task.UserId)
		if err != nil {
			fmt.Println("error when selecting tasks:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var task TodoList
			if err := rows.Scan(
				&task.TaskId,
				&task.TaskName,
				&task.TaskTime,
				&task.TaskComment,
				&task.TaskPriority,
				&task.UserId,
				&task.CreatedAt,
				&task.UpdatedAt); err != nil {
				fmt.Println("error when scanning task:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}
		if err := rows.Err(); err != nil {
			fmt.Println("error after scanning tasks:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//////////////////////
		jsonData, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "Failed to marshal tasks to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("result :", tasks)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

	}
}

func SelectPriority(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		resBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("impossible to read all body of response: %s", err)
		}

		var task model.TodoList

		err = json.Unmarshal(resBody, &task)
		if err != nil {
			fmt.Println("error while unmarshaling data", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tasks := []TodoList{}
		query := `
		SELECT * FROM "todo_list"
		WHERE "task_priority"=$1`

		////////////////

		rows, err := db.Query(ctx, query, task.TaskPriority)
		if err != nil {
			fmt.Println("error when selecting tasks:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var task TodoList
			if err := rows.Scan(
				&task.TaskId,
				&task.TaskName,
				&task.TaskTime,
				&task.TaskComment,
				&task.TaskPriority,
				&task.UserId,
				&task.CreatedAt,
				&task.UpdatedAt); err != nil {
				fmt.Println("error when scanning task:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			tasks = append(tasks, task)
		}
		if err := rows.Err(); err != nil {
			fmt.Println("error after scanning tasks:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//////////////////////
		jsonData, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, "Failed to marshal tasks to JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("result :", tasks)
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

	}
}

type TodoList struct {
	TaskId       string    `json:"task_id"`
	TaskName     string    `json:"task_name"`
	TaskTime     time.Time `json:"task_time"`
	TaskComment  string    `json:"task_comment"`
	TaskPriority string    `json:"task_priority"`
	UserId       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func DeleteTask(ctx context.Context, db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var task TodoList
		if err := json.Unmarshal(body, &task); err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}
		taskId := task.TaskId

		query := `
            DELETE FROM "todo_list"
            WHERE "age" > $1`

		_, err = db.Exec(ctx, query, taskId)
		if err != nil {
			fmt.Println("error when deleting task:", err)
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}

		// Отправка успешного ответа
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Task with ID %s deleted successfully", taskId)
	}
}
