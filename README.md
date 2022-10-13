**DOCKER CONTAINER**
***SETUP***
- docker pull postgres:12-alpine
- docker images
```
docker run --name <container_name> -e <environment_variable> -p <host_ports:container_ports> -d <image>:<tag>
```
- docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
- docker ps

<span style="color:red">Run command in container</span>
```
docker exec -it <container_name_or_id> <command> [args]
```
<span style="color:red">View container logs</span>.
```
docker logs <container_name_or_id>
```

**TABLEPLUS**

**MIGRATE** <br />
***Install Mac***
`brew install golang-migrate` <br />
***Install Window***
`scoop install migrate`

- `migrate -help`
- `migrate create -ext sql -dir db/migration -seq init_schema`
- `docker start postgres12`
- `docker exec -it postgres12 /bin/sh`

**WHAT IS CRUD ?** <br />
<span style="color:cyan;font-weight:bold">CREATE:</span>
`Insert new records to the database` <br />
<span style="color:cyan;font-weight:bold">READ:</span>
`Select or search for records in the database` <br />
<span style="color:cyan;font-weight:bold">UPDATE:</span>
`Change some fields of the record in the database` <br />
<span style="color:cyan;font-weight:bold">DELETE:</span>
`Remove records from the database` <br />

**DATABASE/SQL**
- Very fast & straightforward
- Manual mapping SQL fields to variables
- Easy to make mistakes, not caught until

**GORM**
- CRUD functions already implemented very short production code
- Must learn to write queries using gorm's function
- Run slowly on high load

**SQLX**
- Quite fast & easy to use
- Fields mapping via query text & struct tags
- Failure won't occur until runtime

**SQLC**
- Very fast & easy to use
- Automatic code generation
- Catch SQL Query errors before generating codes
