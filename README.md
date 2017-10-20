# Notes
## tl;dr
Notes is a simplistic note taking app, mostly intended as a demo of how Golang
based projects are/can be structured.  It is aimed at novices (like its
developer), and mostly intended as an example of how to structure a Golang
project that needs subpackages, based on the structure advice provided by
[Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
and (more sauce)
## Requirements
* docker (either native or machine)
## Up & Running
```bash
$ docker-compose -f ./docker/compose/vm/docker-compose.yml up
```
Once the build is finished, go to your_vm_ip:3000/notes/ (or localhost, if running native docker)

## Libraries Used
* [upper.io/db.v3/mysql](https://upper.io/db.v3/mysql) - A database connection library with wrapped driver for mysql.
* [gin-gonic/gin](https://github.com/gin-gonic/gin) - A very fast router.

## ToDo
* Feature: Tags - categorization of notes (Adding new model to project)
* Add a caching layer so we don't repeat DB queries (example of composing a
struct in another struct)
* Improve our Read template to be a read and update template
* Tests & mock service implementations (Need to introduce tests in lesson)
* Switch gin to another auto builder or fix huge build delays. (realize?)

## Changelog
* 2017-10-19 Big rewrite and documentation effort.
* 2017-10-18 Restyles web app to use css.
* 2017-10-17 Adds title field to main branch (too simple for HW anyway).
* 2017-10-13 HTML output instead of JSON, templates introduced.
* 2017-10-12 Initial commit.
