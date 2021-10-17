# Start from golang alpine base image
FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /gomodules

# Copy the gomodules source code to the Working Directory inside the container
COPY gomodules/ .

# Set the Current Working Directory inside the container
WORKDIR /user-ms

# Copy the source code to the Working Directory inside the container
COPY user-ms/ .

# Download all dependencies.
# Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the Go app
RUN go build -o userservice userservice.go

# Set the Current Working Directory inside the container for builded application
WORKDIR /app

# Copy binary from user-ms and configuration to main folder
RUN cp /user-ms/userservice .
RUN cp /user-ms/userservice.toml .

# Delete source code and gomodules
RUN rm -r /gomodules
RUN rm -r /user-ms

# Export necessary ports
EXPOSE 8080 8081

# Command to run when starting the container
CMD ["/app/userservice"]
