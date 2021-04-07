# go-fiber-practice

Use [Fiber](https://github.com/gofiber/fiber) web framework to build a tiny web service.

## Prerequisites

* Go v1.16.3+
* Fiber v2
* Mysql 8.0+

### Run a MySQL DB by using Docker
```bash
# user: root
# password: rootroot
docker run --name go-fiber-practice -e MYSQL_ROOT_PASSWORD=rootroot -p 3306:3306 -d mysql:8.0

# enter mysql's interaction mode by (-u) root identity and (-p) without pointing a specific database
docker exec -it go-fiber-practice mysql -u root -p
```

### Create a `go_admin` database
``` mysql
CREATE DATABASE
 IF
   NOT EXISTS go_admin DEFAULT CHARACTER
   SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| go_admin           |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```
