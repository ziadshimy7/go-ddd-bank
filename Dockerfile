# syntax=docker/dockerfile:1

# Use a specific version to ensure consistent builds.
FROM golang:1.19 AS build-stage

WORKDIR /app

# Only copy the Go manifest and lock files initially to leverage Docker caching 
# and avoid downloading dependencies if these files haven't changed.
COPY go.mod go.sum ./
RUN go mod download

# Copy only the necessary source files to build the Go application.
# This can help in reducing the image build context.
# Ideally, use a .dockerignore to exclude files/directories like tests, READMEs, etc.
COPY . .

# Build the application in one layer to reduce the number of layers.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /ddd-bank ./

# Use a minimal runtime image.
FROM gcr.io/distroless/base-debian11

# For added security, avoid adding any unnecessary files or environment variables to the final image.
WORKDIR /
COPY --from=build-stage /ddd-bank /ddd-bank

# The .env file might contain secrets or other sensitive data. 
# Consider using Docker Secrets, Kubernetes Secrets, or environment variables for this data.
COPY --from=build-stage /app/.env .env

EXPOSE 8080

# Using a non-root user is good for security.
USER nonroot:nonroot

CMD ["/ddd-bank"]