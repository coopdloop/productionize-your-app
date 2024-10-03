# E-commerce API

This is a simple e-commerce API built with Go using the standard net/http package.

## Running the application

1. Build the application:
   ```
   make build
   ```

2. Run the application:
   ```
   make run
   ```

3. Run tests:
   ```
   make test
   ```

## API Endpoints

- GET /products: Get all products
- GET /products/{id}: Get a specific product
- POST /products: Create a new product
- PUT /products/{id}: Update a product
- DELETE /products/{id}: Delete a product

## Docker

To build and run the Docker image:

```
docker build -t ecommerce-api .
docker run -p 8080:8080 ecommerce-api
```
