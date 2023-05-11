# Use the official Go image as the base image
FROM golang:1.18-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go app
RUN go build -o main .

# Use a smaller base image for the final stage
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/main .
 
 

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
