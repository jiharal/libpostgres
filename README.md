# Overview
 libpostgres is a collection of postgresql connections from various sources such as making connections with amp, elastic, etc.

# How to use

## Connect to database
```go
  dbOptions := libpostgres.Options{
		Host:     "db_host",
		Port:     "db_port",
		Username: "db_username",
		Password: "db_password",
		DBName:   "db_name",
		SSLMode:  "db_sslmode",
	}
	db, err := libpostgres.Connect(dbOptions)
	if err != nil {
		fmt.Println("Error conn to DB", err)
		panic(err)
	}
```

## Migrate database
 - Create directory `db/migrations` on yout project.
 - Create your file `sql` with format name 
  - 000001_create_users_table.down.sql
  - 000001_create_users_table.up.sql
In the `.up.sql` file let's create the table:
  ```sql
  CREATE TABLE IF NOT EXISTS users(
    user_id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (300) UNIQUE NOT NULL
  );
  ```
And in the .down.sql let's delete it:
  ```sql
  DROP TABLE IF EXISTS users;
  ```

Read more [here](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

 ```go
	err := Migrate(db, "file://./migrations")
	if err != nil {
		fmt.Println("Error migrate DB", err)
		panic(err)
	}
 ```

# Constribute
 Please welcome .