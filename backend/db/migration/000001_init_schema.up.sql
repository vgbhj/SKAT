CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL
);

INSERT INTO users (username, password)
	VALUES(
		'admin',
		'$2a$10$YgzepzPAE0OZWr9P6mQVu.Ind9xcSN/DGCfOiVT8XClxWjWLWbfpa'
);