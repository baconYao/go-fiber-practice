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
| order_items        |
| orders             |
| permissions        |
| products           |
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

* PUT `/api/users/info`
  * Body
    ```json
    {
        "first_name": "new yao222",
        "last_name": "ppchang",
        "email": "hip@mail.com"
    }
    ```

* PUT `/api/users/password`
  * Body
    ```json
    {
      "password": "123456",
      "password_confirm": "123456"
    }
    ```

* POST `/api/logout`
  * Response
    ```json
    {
        "message": "success"
    }
    ```

* GET `/api/users?page=1`
  * Response
    ```json
    {
        "data": [
            {
                "id": 1,
                "first_name": "peiyao",
                "last_name": "chang",
                "email": "yyy@gamil.com",
                "role_id": 1,
                "role": {
                    "id": 1,
                    "name": "admin",
                    "permissions": null
                }
            },
            {
                "id": 7,
                "first_name": "hi",
                "last_name": "there",
                "email": "yyy22@gamil.com",
                "role_id": 1,
                "role": {
                    "id": 1,
                    "name": "admin",
                    "permissions": null
                }
            }
        ],
        "meta": {
            "last_page": 0,
            "page": 1,
            "total": 2
        }
    }
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

* POST `/api/roles`
  * Body
    ```json
    {
        "name": "test2",
        "permissions": ["1", "3"]
    }
    ```
  * Response
    ```json
    {
        "id": 5,
        "name": "test2",
        "permissions": [
            {
                "id": 1,
                "name": ""
            },
            {
                "id": 3,
                "name": ""
            }
        ]
    }
    ```

* PUT `/api/roles/:id`
  * Body
    ```json
    {
        "name": "test2 to 3",
        "permissions": ["5", "4"]
    }
    ```
  * Response
    ```json
    {
        "id": 5,
        "name": "test2 to 3",
        "permissions": [
            {
                "id": 4,
                "name": ""
            },
            {
                "id": 6,
                "name": ""
            }
        ]
    }
    ```

* POST `/api/upload`
  * Body
    form-data
      key: image
        type: file
  * Response
    ```json
    {
        "url": "http://localhost:3000/api/uploads/???.jpg"
    }
    ```
  * ![Alt text](./readme_images/upload-form-example.png)

* GET `/api/uploads/[image_name].extension`
  * Response
    file
  * ![Alt text](./readme_images/get-uploaded-image.png)