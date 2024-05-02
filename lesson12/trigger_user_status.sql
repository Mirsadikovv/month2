CREATE TABLE members (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE
);

CREATE TABLE memberships (
    id SERIAL PRIMARY KEY,
    member_id INT NOT NULL REFERENCES members(id),
    membership_type VARCHAR(50) NOT NULL DEFAULT 'free'
);


CREATE OR REPLACE FUNCTION create_membership()
RETURNS TRIGGER
AS 
$$
BEGIN 
  INSERT INTO memberships(member_id)
  VALUES (NEW.id);
  RETURN NEW;
END;
$$
LANGUAGE plpgsql;


CREATE TRIGGER after_insert_memeber_trigger
AFTER INSERT ON members
FOR EACH ROW 
EXECUTE FUNCTION create_membership();

INSERT INTO members(name, email)
VALUES('John Doe', 'john.doe@gmail.com'),
('John Roe', 'john.roe@gmail.com')
RETURNING *;


SELECT * FROM memberships;


ALTER TRIGGER trigger_name
ON table_name 
RENAME TO new_trigger_name;

