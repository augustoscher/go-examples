# MySQL

## Setup

Running MySQL on docker

```bash
docker run \
 --name mysql \
 -e MYSQL_ROOT_PASSWORD=root \
 -p 33060:3306 \
 -d \
 mysql:latest
```

Connect on database inside of container

```bash
docker exec -it mysql bash
mysql -u root -p
```

Running Adminer on docker

```bash
docker run \
  --name adminer \
  -p 8080:8080 \
  --link mysql:mysql \
  -d \
  adminer
```

Access [adminer](http://localhost:8080/?server=mysql&username=root).
Password is: root
