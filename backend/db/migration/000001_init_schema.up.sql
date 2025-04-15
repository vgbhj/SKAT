CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);


CREATE TABLE materials (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  filename VARCHAR(255) NOT NULL,
  description TEXT,
  user_id INT REFERENCES users(id),
  upload_date DATE
);


INSERT INTO users (username, password)
	VALUES(
		'admin',
		'$2a$10$YgzepzPAE0OZWr9P6mQVu.Ind9xcSN/DGCfOiVT8XClxWjWLWbfpa'
);