# tajalevegsuppliers
    
# tools
    - go , postgres@16


# things required to start the service
    - brew services start postgresql@16
    - source ~/.bash_profile
    - psql postgres
    

# Create a migration file
migrate create -ext sql -dir migrations -seq create_tables

# Edit the created migration file in the "migrations" folder with your SQL statements.

# Run migrations
migrate -path migrations -database "postgres://your-username:your-password@your-postgres-host:5432/your-database-name?sslmode=disable" -verbose up


