CREATE TABLE IF NOT EXISTS products(
    id SERIAL PRIMARY KEY,
    product_name VARCHAR,
    product_price INT,
    product_quantity INT
);

CREATE TABLE IF NOT EXISTS categories(
    categorie_id SERIAL,
    categorie_name VARCHAR,
    product_id INT REFERENCES products(id)
);

CREATE OR REPLACE FUNCTION add_product(
    product_name VARCHAR,
    product_price INT,
    product_quantity INT,
    category_name VARCHAR
)
RETURNS INT AS $$
DECLARE
    new_product_id INT;
BEGIN
    INSERT INTO products (product_name, product_price, product_quantity)
    VALUES (product_name, product_price, product_quantity)
    RETURNING id INTO new_product_id;

    INSERT INTO categories (categorie_name, product_id)
    VALUES (category_name, new_product_id);
    RETURN new_product_id;
END;
$$ LANGUAGE plpgsql;







CREATE OR REPLACE FUNCTION update_quantity(
    product_id INT,
    new_quantity INT
)
RETURNS VOID AS $$
BEGIN
    UPDATE products
    SET product_quantity = new_quantity
    WHERE id = update_quantity.product_id;
END;
$$ LANGUAGE plpgsql;






CREATE OR REPLACE FUNCTION get_product_info(
    given_product_id INT
)
RETURNS TABLE (
    product_id INT,
    product_name VARCHAR,
    product_price INT,
    product_quantity INT
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        id AS product_id,
        product_name,
        product_price,
        product_quantity
    FROM 
        products
    WHERE 
        id = get_product_info.given_product_id;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION list_by_category(
    category_name  VARCHAR
)
RETURNS TABLE(
    product_id INT,
    product_name VARCHAR,
    product_price INT,
    product_quantity INT
) AS $$
BEGIN
    RETURN QUERY
    SELECT 
        p.id AS product_id,
        p.product_name,
        p.product_price,
        p.product_quantity
    FROM 
        products p
    JOIN 
        categories c ON p.id = c.product_id
    WHERE 
        c.categorie_name = list_by_category.category_name;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION calculate_total_inventory_value()
RETURNS NUMERIC AS $$
DECLARE
    total_value NUMERIC;
BEGIN
    SELECT SUM(product_price * product_quantity) INTO total_value
    FROM products;

    RETURN total_value;
END;
$$ LANGUAGE plpgsql;
