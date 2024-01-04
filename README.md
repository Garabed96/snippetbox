**CMD Directory** contains application specific code for executing the application.
The Web Application lives in cmd/web directory.

**internal directory** contains ancillary non-application-specific code used in project.
Reusable code like validation helpers, SQL database models

**ui directory** contains user-interface assets by the web app, 
ui/html will contain HTML templates
ui/static contains static files (css / images)

**RUN the application**
`go run ./cmd/web`