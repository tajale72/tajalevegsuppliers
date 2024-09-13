# tajalevegsuppliers
    
# tools
    - go , postgres@16


# things required to start the service
    - brew services start postgresql@16
    - source ~/.bash_profile
    - psql postgres

# docker run 
    docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=tajalevegsuppliers -p 5432:5432 -d postgres

# docker run to use the local volume
   docker run --name postgres-container \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=admin \
  -e POSTGRES_DB=tajalevegsuppliers \
  -p 5432:5432 \
  -v ~/postgres-data:/var/lib/postgresql/data \
  -d postgres

# docker run 
    docker exec -it postgres-container psql -U postgres -d tajalevegsuppliers

# Create a migration file for postgres
 pg_dump -h localhost -U postgres -d tajalevegsuppliers > dbname_backup.sql




