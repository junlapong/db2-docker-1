# Docker WAS & DB2

IBM WebSphere and DB2 containers with JDBC connection

from: https://github.com/keensoft/was-db2-docker

## Description

Docker composition to ease application testing in IBM WebSphere environment. It contains IBM WebSphere Application Server and IBM DB2 database server connected using server datasource.

Servers base is created using IBM provided containers in [Docker Hub](https://hub.docker.com/u/ibmcom/) .

You can add your application to WebSphere container assets to be automatically deployed. See detailed info in _was_ dir.

## Volumes

A shared volume is created between the containers to use DB2 drivers in WebSphere JDBC provider.

## Usage

```
docker-compose up
```

## Credentials

### DB2

DB2 provided

- Username: **db2inst1** is used to send all commands to database
- Password: **pa55w0rd** is configured via environment value `DB2INST1_PASSWORD` or `db2/.env`
- Hostnname/Port: localhost:50000

### WebSphere

WebSphere provided

- Username: **wsadmin**
- Password: **pa55w0rd**

via `was/assets` file.

## URLs

- **IBM WebSphere Administrative Console:** https://localhost:9043/ibm/console
- **Deployed application:** http://localhost:9080/docker

