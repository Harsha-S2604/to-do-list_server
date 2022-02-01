# to-do-list_server

## Installation

### 1. Install golang on your machine
   
   go to this [page](https://go.dev/dl/) and follow the instructions to install golang on your machine

### 2. Install PostgreSQL on your machine
   
   go to this [page](https://www.postgresql.org/download/) and follow the instructions to install postgresql on your machine
   
### 3. Install Redis on your machine
   
   go this [page](https://redis.io/topics/quickstart) and follow the instructions to install redis on your machine
  
## Database setup

### For linux
   
   step 1: Open terminal. Using your favorite editor, Open the `.bashrc` file located in the /home/<username> directory. I am using vim
   
   ```console
   $ vim ~/.bashrc
   ```
   
   step 2: Declare the environment variable
   
   ```bash
   export POSTGRE_DB_USERNAME=<username>
   export POSTGRE_DB_PASSWORD=<password>   
   export POSTGRE_DB_NAME='test_db'
   export POSTGRE_DB_PORT='5432'
   ```
