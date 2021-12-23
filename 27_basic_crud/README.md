To start a simple test database: `docker run --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=superpass -d mysql:8.0`

To run the existing container: `docker start mysql-test`

To run the app: `go run main.go`
