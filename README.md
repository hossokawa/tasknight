# Introduction

Tasknight is a web-based task manager built using the GoTTH stack (Go, Templ, Tailwind CSS, HTMX) and MongoDB. Currently the project has all the task functionalities in place and also user registration, but I still need to fully implement user login/JWT authentication and associate the tasks with the user who created them.

# Installation

1. Create a .env file and set the PORT, MONGO_URI and JWT_SECRET environment variables
2. Run `tailwind -o ./static/css/output.css --watch` to generate the css file and watch for any modifications
3. Run `air` to watch for changes in templ files and start the server
