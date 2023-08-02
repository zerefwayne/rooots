## How to run database?

### Fetch official postgresql image

### Start a container

```
docker run --name roootsdb -p 5432:5432 -e POSTGRES_USER=roootsadmin -e POSTGRES_PASSWORD=roootspw -e POSTGRES_DB=roootsdb -d postgres
```

### Use psql running inside container

```
docker exec -it roootsdb psql -U roootsadmin roootsdb
```