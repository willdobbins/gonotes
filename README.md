# Notes
A bare-bones note taking app. 

## Up & Running
```bash
docker-compose -f ./docker/compose/vm/docker-compose.yml up
```

## Desired Endpoints
```
GET    /notes/        List of existing notes
POST   /notes/        Create a new note
GET    /notes/:id     Fetch a specific note
DELETE /notes/:id     Delete an existing note
PATCH  /notes/:id     Alter an existing note
```

## Current Note Schema
```
id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
body TEXT NOT NULL
```

## ToDo
* New field: title (string)
* Feature: Tags - quick categorization system
* templates instead of JSONifying all the things (In Progress: see branch 'templates')
* Add a caching layer so we don't repeat DB queries?