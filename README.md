# Multi Tenant reservation-api

**Description**:   A simple RESTful API for booking and reservation hotel rooms. <br>
A Golang web-based restful api hotel reservation system, that acts a broker between hotels and customers.
# 

**Application architecture**:<br>
The architecture used in this program is **multitenancy**, for each new **hotel(tenant)**, a new separate database is created, all of its tables are separate, and the information of each **hotel(tenant)** is stored separately in its own database. <br>
This pattern uses a **multi-tenant** application with many **databases**, all being single-tenant databases. A new database is provisioned for each new **tenant**.
<br>
<img src="https://github.com/RezaEskandarii/repository-images/blob/master/saas-multi-tenant-app-database-per-tenant-13.png"> <br>
# used technologies
- PostgreSql
- Rabbitmq
- Redis
- Minio 
- Docker
- DockerCompose
## How to run(with docker):
To run this program, you can easily use **DockerCompose**, all required dependencies such as Postgres, Redis, etc. are listed in the <br> **docker-compose** file.<br>
First of all, install Docker and Docker Compose on your machine and after run the `docer-compose up -d` command, the program is executed, the default database It is made and also all the seeds are done.
## How to run(without docker):
If for any reason you don't want to run the program with Docker and just want to run and compile the code, follow the steps below:
- install **postgresql** on your machine
- install **rabbitmq** on your machine
- install **redis** on your machine
- install **minio** on your machine
- go to **resorecess** directory and override the **config.yml** with following steps:
  1. add your postgres port and ip address
  2. add your redis client port and ip address
  3. add your rabbitmq address
  4. add your minio conenction
