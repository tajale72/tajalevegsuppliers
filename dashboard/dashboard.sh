
# Build the Docker image
docker build -t dashboard .

# Run the Docker container
docker run -p 3000:3000 dashboard
