# Notes
A bare-bones note taking app. 

## Up & Running

## Desired Endpoints
```
GET    /notes/        List of existing notes
POST   /notes/        Create a new note
GET    /notes/:id   Fetch a specific note
DELETE /notes/:id   Delete an existing note
PATCH  /notes/:id     Alter an existing note
```

## Note Schema
```
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
body TEXT NOT NULL
```

## ToDo
* Add a title to a note
* Track creation date
* Tags - quick categorization system
* Better templates