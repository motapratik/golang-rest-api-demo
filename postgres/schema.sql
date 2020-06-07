
CREATE TABLE users (
   user_id serial PRIMARY KEY,
   name VARCHAR (150) NOT NULL,
   mobile_number VARCHAR (20) NOT NULL,
   password VARCHAR (60) NOT NULL,
   created_at TIMESTAMP NOT NULL
);


CREATE EXTENSION pgcrypto;

--CALL sf_register_user('A','1234','asdf')

-- Procedure to register a new user
CREATE OR REPLACE PROCEDURE sf_register_user(
  pName VARCHAR (150),
  pNumber VARCHAR (20),
  pPass VARCHAR (60)
) LANGUAGE plpgsql
AS $$
BEGIN
  INSERT INTO users (
    name, mobile_number, password, created_at
  ) VALUES (
    pName,
    pNumber,
    crypt(pPass, gen_salt('bf')),
    NOW()
  );

END;
$$;