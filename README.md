<h1>ToDo</h1>
<h3>ToDo is a web application that allows its users to store their tasks. </h3>
<h3>Users can give their tasks titles, categorize them, write detailed text about the task, change its state in progress and set the deadlines.</h3>
<br>
<h1>How to run</h1>
<h3>Notice: To run it on your computer you should have GO programing language downloaded.</h3>
<li>Either clone the repo or download archive</li>
<li>Go to the project folder</li>
<li>In terminal type <b><em>go run .</em></b> or <b><em>go run main.go</em></b></li>
<li>Or build docker image <b><em>docker build -t [name]:[tag] .</em></b> and then run it <b><em>docker run -p [port]:[port] [name]:[tag]</em></b></li>
<li>And proceed to the app by clicking <a href="http://localhost:3000">here</a></li>
<br>
<h1>Technologies</h1>
<h3>The server side of the website was written in GO programing language, its frameworks and packages (Gin, Gorilla etc.).</h3>
<h3>As DBMS MongoDB was used. The project uses MongoDB Atlas to store data in the cloud.</h3>
<h3>The client side of the website was written in HTML, CSS (BootStrap).</h3>
<br>

# ✅ API Documentation
## 🔐 Authentication & Authorization

| Method | Path        | Handler           | Description                           | Access Level  |
|--------|-------------|-------------------|---------------------------------------|---------------|
| GET    | /login      | LoginPage         | Serve login page                      | Public        |
| POST   | /login      | Login             | Authenticate user, return JWT/session| Public        |
| GET    | /logout     | Logout            | Log out user                          | Authenticated |
| GET    | /register   | RegistrationPage  | Serve registration page               | Public        |
| POST   | /register   | Register          | Register a new user                   | Public        |

## 🏠 Home Page

| Method | Path | Handler  | Description      | Access Level |
|--------|------|----------|------------------|--------------|
| GET    | /    | HomePage | Render homepage  | Public       |

## 📋 ToDo Operations

| Method | Path         | Handler      | Description                       | Access Level  |
|--------|--------------|--------------|-----------------------------------|---------------|
| GET    | /todo        | ToDoPage     | Show all user's tasks             | Authenticated |
| GET    | /add         | AddToDoPage  | Render task creation form         | Authenticated |
| POST   | /add         | AddToDo      | Submit new task                   | Authenticated |
| GET    | /read/:id    | ReadPage     | View details of a specific task   | Authenticated |
| GET    | /edit/:id    | EditPage     | Render edit form for a task       | Authenticated |
| POST   | /edit/:id    | EditToDo     | Submit edited task                | Authenticated |
| GET    | /delete/:id  | DeleteToDo   | Delete a specific task            | Authenticated |

## 🔎 Search, Sort, Group (User View)

| Method | Path     | Handler   | Description                        | Access Level  |
|--------|----------|-----------|------------------------------------|---------------|
| GET    | /search  | ToDoPage  | Show search results for tasks      | Authenticated |
| GET    | /sort    | ToDoPage  | Show sorted task list (by date)    | Authenticated |

## 🛠️ Admin Operations

| Method | Path                 | Handler     | Description                            | Access Level |
|--------|----------------------|-------------|----------------------------------------|--------------|
| GET    | /admin               | AdminPage   | Admin dashboard                        | Admin        |
| GET    | /admin/search        | AdminPage   | Search across users' tasks             | Admin        |
| GET    | /admin/sort          | AdminPage   | View sorted tasks for all users        | Admin        |
| GET    | /admin/delete/:user  | DeleteUser  | Delete a user and their data           | Admin        |

## 📁 Static Files

| Method | Path       | Description                     |
|--------|------------|---------------------------------|
| GET    | /static/*  | Serve static assets (CSS/JS/img)|

