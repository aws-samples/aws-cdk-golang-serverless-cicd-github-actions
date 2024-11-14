FROM golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5816886514199316e4a2fb as build
WORKDIR /lambda
# Copy dependencies list
COPY go.mod go.sum ./

# Build with optional lambda.norpc tag
COPY lambda.go .
RUN go build -tags lambda.norpc -o bootstrap lambda.go
FROM golang:1.21@sha256:4746d26432a9117a5f58e95cb9f954ddf0de128e9d5816886514199316e4a2fb as run

# Create a non-root user
RUN useradd -m appuser

WORKDIR /lambda

# Copy the binary from the build stage
COPY --from=build /lambda/bootstrap ./bootstrap

# Add curl for healthcheck
RUN yum update -y && yum install -y curl && yum clean all

# Add a simple healthcheck script
COPY healthcheck.sh .
RUN chmod +x healthcheck.sh

# Change ownership of the application files to the new user
RUN chown -R appuser:appuser /lambda

# Switch to the non-root user
USER appuser

# Add HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 \
  CMD ["./healthcheck.sh"]
  
ENTRYPOINT [ "./bootstrap" ]