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
- Docker
## How to run:
To run this program, you can easily use **DockerCompose**, all required dependencies such as Postgres, Redis, etc. are listed in the <br> **docker-compose** file.<br>
First of all, install Docker and Docker Compose on your machine and after run the `docer-compose up -d` command, the program is executed, the default database It is made and also all the seeds are done.

 
