FROM golang

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# download all the required dependency for our application
RUN go mod download

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN make build

# This container exposes port 8080 to the host os
EXPOSE 8080

# Run the executable
CMD ["./out/tv-tracker"]

