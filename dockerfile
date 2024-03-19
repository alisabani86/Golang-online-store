# Use the official Go image with version 1.20
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod file to the working directory
COPY go.mod .
COPY . .

# Application builder step
RUN go mod tidy && go mod download
RUN go build .

# Make the binary executable
RUN chmod +x /app/bin



# Set the entry point for the container
ENTRYPOINT ["/app/bin"]
