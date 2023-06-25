# Build the Docker image
docker-build:
	docker build -t fiber-caching-api .

# Run the Docker container
docker-run: docker-build
	docker run -d -p 3000:3000 fiber-caching-api
