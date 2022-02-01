# to-do-list_server

## i) Installation

### 1. Install golang on your machine
   
   go to this [page](https://go.dev/dl/) and follow the instructions to install golang on your machine

### 2. Install PostgreSQL on your machine
   
   go to this [page](https://www.postgresql.org/download/) and follow the instructions to install postgresql on your machine
   
### 3. Install Redis on your machine
   
   go this [page](https://redis.io/topics/quickstart) and follow the instructions to install redis on your machine
  
## ii) Database setup

### For linux
   
   step 1: Open terminal. Using your favorite editor, Open the `.bashrc` file located in the /home/<username> directory. I am using vim
   
   ```console
   $ vim ~/.bashrc
   ```
   
   step 2: Declare the environment variable for postgresql
   
   ```bash
   export POSTGRE_DB_USERNAME=<username>
   export POSTGRE_DB_PASSWORD=<password>   
   export POSTGRE_DB_NAME='test_db'
   export POSTGRE_DB_PORT='5432'
   ```
   
   step 3: Declare the environment variable for redis
   ```bash
   export REDIS_HOST="localhost:6379"
   export REDIS_PASSWORD=<password>
   ```
   
   step 4: save the changes and restart the terminal
   
   step 5: login to your postgresql cli using your username and password
   ```console
   $ sudo -u <username> psql
   ```
   once you are logged in, the terminal prompt looks like(in my case the username is postgres)
   ```console
   postgres=#
   ```
   
   step 6: Create a database called test_db
   ```console
   postgres=# CREATE DATABASE test_db;
   ```
   
   step 7: Use the database and create a two tables
   ```console
   postgres=# \c test_db
   test_db=# CREATE TABLE users(user_id INT NOT NULL GENERATED ALWAYS AS IDENTITY, email VARCHAR(255) NOT NULL, PRIMARY KEY(user_id));
   test_db=# CREATE TABLE todo_list(task_id INT GENERATED ALWAYS AS IDENTITY,
   task_name TEXT NOT NULL,
   is_completed BOOLEAN NOT NULL,
   user_id INT NOT NULL,
   PRIMARY KEY(task_id),
   CONSTRAINT fk_users
      FOREIGN KEY(user_id) 
	  REFERENCES users(user_id));
   ```
   
 ## iii) Server setup
   
   step 1: Clone the project
   ```console
   $ git clone https://github.com/Harsha-S2604/to-do-list_server.git
   ```
   
   step 2: change directory to the cloned project
   ```console
   $ cd to-do-list_server
   ```
   
   step 3(optional): download the dependencies
   this step is optional since `go run/go build` automatically runs the `go mod download`
   ```console
   $ go mod download
   ```
   
   step 4: run the server
   ```console
   $ go run main.go
   ```
   
   now the server runs on localhost:8080 or the port you have configured
   
 ## iv) Working
   There are totally 5 APIs AddUser, AddTask, GetTask, UpdateTask, RemoveTask
   step 1: open the terminal
   ### Add User
   ```console
   $ curl --header "Content-Type: application/json" \
		--request POST \
		--data '{"email":"test@123.com"}' \
		'localhost:8080/api/v1/todo/user/add'
   ```
   
   ### Add Task
   ```console
   $ curl --header "Content-Type: application/json" \
		--request POST \
		--data '{"user":{"userId":9}, "taskName":"backend_task", "isCompleted":false}' \
		'localhost:8080/api/v1/todo/task/add'
   ```
   
   ### Get Task
   ```console
   $ curl --request GET 'localhost:8080/api/v1/todo/task/tasks/5?limit=5&offset=1'
   ```
   
   ### Update Task
   ```console
   $ curl --request PUT 'localhost:8080/api/v1/todo/task/update/9?userId=5&offset=1' -F isCompleted=true
   ```
   
   ### Delete Task
   ```console
   $ curl --request DELETE 'localhost:8080/api/v1/todo/task/remove/9?userId=5&offset=1'curl --request DELETE 'localhost:8080/api/v1/todo/task/remove/9?userId=5&offset=1'
   ```
   
