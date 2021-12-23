To start a simple test database: `docker run --name mysql-test -p 3306:3306 -e MYSQL_ROOT_PASSWORD=superpass -d mysql:8.0`

To run the existing container: `docker start mysql-test`

To connect just use the following parameters:
- user: root
- password: superpass
- port: 3306
- host: localhost
If using DBeaver it may need to set allowPublicKeyRetrieval to true on driver properties settings.
