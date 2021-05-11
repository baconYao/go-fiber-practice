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

After creating `go_admin` DB, start up the server, then `GORM` will do `migration` to create tables like below

```mysql
mysql> show tables;
+--------------------+
| Tables_in_go_admin |
+--------------------+
| permissions        |
| role_permissions   |
| roles              |
| users              |
+--------------------+
4 rows in set (0.00 sec)
```

### Permissions

Pre-build the permission values

```mysql
mysql> INSERT INTO permissions (Name) VALUES ('view_orders');

mysql> SELECT * FROM permissions;
+----+---------------+
| id | name          |
+----+---------------+
|  1 | view_users    |
|  2 | edit_users    |
|  3 | view_roles    |
|  4 | edit_roles    |
|  5 | view_products |
|  6 | edit_products |
|  7 | view_orders   |
|  8 | edit_orders   |
+----+---------------+
8 rows in set (0.00 sec)
```

## Develop

```bash
# Run server support live reload
$ air
```

PS. I am using [`air live reload`](https://github.com/cosmtrek/air)

## API example

* POST `/api/register`
  * Body
    ```json
    {
        "first_name": "peiyao",
        "last_name": "chang",
        "email": "yyy@gamil.com",
        "password": "1234qwer",
        "password_confirm": "1234qwer"
    }
    ```

* POST `/api/login`

  * Body
    ```json
    {
        "email": "yyy@gamil.com",
        "password": "1234qwer"
    }
    ```
  * Response

    `jwt token string`

* GET `/api/user`
  * Response
    ```json
    {
        "id": 1,
        "first_name": "peiyao",
        "last_name": "chang",
        "email": "yyy@gamil.com",
    }
    ```

* POST `/api/logout`
  * Response
    ```json
    {
        "message": "success"
    }
    ```

* GET `/api/users`
  * Response
    ```json
    [
        {
            "id": 1,
            "first_name": "peiyao",
            "last_name": "chang",
            "email": "yyy@gamil.com",
        },
        {
            "id": 2,
            "first_name": "haha",
            "last_name": "QQ",
            "email": "maillll"
        }
    ]
    ```

* POST `/api/users`
  * Body
    ```json
    {
        "first_name": "haha",
        "last_name": "QQ",
        "email": "asdas@gamil.com",
    }
    ```
  