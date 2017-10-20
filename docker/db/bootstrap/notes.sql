#Notes db is already created by docker up.
CREATE TABLE notes.notes ( id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, title VARCHAR(255) NOT NULL, body TEXT NOT NULL);
