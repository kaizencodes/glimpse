# Use an official Go runtime as the base image
FROM golang:1.20.7

RUN mkdir glimpse
# Set the working directory inside the container
WORKDIR /glimpse

# Copy the source code into the container's working directory
COPY . .

# Build the Go application
RUN go mod download \ 
        && go build -o glimpse \
        && go install cuelang.org/go/cmd/cue@latest

# Expose the port the application listens on
# EXPOSE 8080

# Command to run the application
CMD ["./glimpse"]
