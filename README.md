**CMD Directory** contains application specific code for executing the application.
The Web Application lives in cmd/web directory.

**internal directory** contains ancillary non-application-specific code used in project.
Reusable code like validation helpers, SQL database models

**ui directory** contains user-interface assets by the web app, 
ui/html will contain HTML templates
ui/static contains static files (css / images)

**RUN the application**
`go run ./cmd/web`

### **CLI Flags when starting application**

`go run ./cmd/web -addr=":80"`


### Connect to mysql
`mysql -u root -p`

### MySQL Location:
**Generally they are here** 

For Linux/Unix: /var/lib/mysql
For macOS: /usr/local/mysql/data
For Windows: C:\ProgramData\MySQL\MySQL Server 8.0\Data

** But if you can't find it there **
For example I installed mysql using homebrew, 
go into mysql and run:

`SHOW VARIABLES LIKE "datadir";` 

# creating user in mysql
`CREATE USER 'web'@'localhost';`

`GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';`

`ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';`