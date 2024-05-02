-- Active: 1713183527088@@127.0.0.1@5432@mydatabase@public


CREATE OR REPLACE FUNCTION count_books()
RETURNS INT AS
$$
DECLARE total_count INTEGER;2
BEGIN
SELECT count(*) INTO total_count FROM books;
RETURN total_count;
END;
$$
LANGUAGE plpgsql;

SELECT count_books();



CREATE OR REPLACE FUNCTION avg_books()
RETURNS INT AS
$$
DECLARE total_count INTEGER;
BEGIN
SELECT avg(narxi) INTO total_count FROM books;
RETURN total_count;
END;
$$
LANGUAGE plpgsql;

SELECT avg_books();
INSERT INTO upmig (id, name)(SELECT avg_books(),'qwerty');



CREATE OR REPLACE FUNCTION count_books()
RETURNS INT AS
$$
import psycopg2

# Устанавливаем соединение с базой данных
conn = psycopg2.connect("dbname=your_db user=your_user password=your_password")

# Создаем курсор для выполнения SQL-запросов
cur = conn.cursor()

# Выполняем запрос на подсчет книг
cur.execute("SELECT count(*) FROM books")

# Получаем результат запроса
total_count = cur.fetchone()[0]

# Закрываем курсор и соединение
cur.close()
conn.close()

return total_count
$$
LANGUAGE plpython;
