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

# Create a migration file
migrate create -ext sql -dir migrations -seq create_tables

# Edit the created migration file in the "migrations" folder with your SQL statements.

# Run migrations
migrate -path migrations -database "postgres://your-username:your-password@your-postgres-host:5432/your-database-name?sslmode=disable" -verbose up
