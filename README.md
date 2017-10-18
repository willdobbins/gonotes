# Notes
A bare-bones note taking app which mostly exists as a training wheels app for
looking at Go projects and how they function and can be structured.

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
title VARCHAR(255) NOT NULL,
body TEXT NOT NULL
```

## ToDo
* Fixup: Add defer Close() for db connections.
* Feature: Color your notes? (simple 1 field change
* Feature: Tags - categorization of notes (Adding new model to project)
* Add a caching layer so we don't repeat DB queries (example of composing a
struct in another struct)
* Add an edit template and switch our routes to be less REST server and more app
* Tests & mock service implementations (Need to introduce tests in lesson)
* Switch gin to another auto builder or fix huge build delays. (realize?)
* Document with an eye towards explaining to novice go devs.

## Changelog
* 2017-10-18 Restyles web app to use css.
* 2017-10-17 Adds title field to main branch (too simple for HW anyway).
* 2017-10-13 HTML output instead of JSON, templates introduced.
* 2017-10-12 Initial commit.
