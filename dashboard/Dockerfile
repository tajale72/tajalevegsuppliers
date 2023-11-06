# Use an official Go runtime as a parent image
FROM  golang:1.19

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Compile the Go app

RUN go mod download

RUN go build -o main .

EXPOSE 3000

# Set the command to run the executable
CMD ["./main"]