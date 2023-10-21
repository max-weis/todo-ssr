# todo-ssr

This project is a small To-Do CRUD app written in Go. It is designed to be self-contained within a single binary with no external
dependencies, making it extremely easy to deploy and run. The user interface is server-side rendered (SSR) to keep it simple and efficient. The backend is powered by the Echo
framework and SQLite for data persistence. Adhering to Clean Code and Boundary-Control-Entity.

## Installation Guide

Follow these steps to get the ToDo app up and running on your machine:

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/max-weis/todo-ssr.git
   cd todo-app
   ```

2. **Build the Application**:
   Make sure you have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).
   ```bash
   go build ./cmd/todo/
   ```

3. **Run the Application**:
   ```bash
   ./todo
   ```

   The application will start and listen on port `8080` by default. You can now access it at [http://localhost:8080](http://localhost:8080).
