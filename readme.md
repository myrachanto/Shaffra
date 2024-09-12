# User Microservice

This repository provides a User Microservice featuring CRUD operations, concurrency management, and comprehensive testing. It also includes API documentation and a Dockerfile for containerization.

## Features

- **CRUD Operations:** Create, Read, Update, and Delete user data.
- **Concurrency Management:** Handles concurrent requests efficiently.
- **Testing:** Includes unit and integration tests.
- **API Documentation:** Detailed documentation for the service endpoints.
- **Docker Support:** Dockerfile included for creating a Docker image.

## Getting Started

To get started with the User Microservice, follow these steps:

1. **Clone the Repository:**
```bash
 git clone https://github.com/myrachanto/Shaffra
```
Navigate to the User Directory:

cd Shaffra/user

Run the Service:

  ```bash
make run
```
Build Docker Image: The repository includes a Dockerfile. To build the Docker image, use:

 ```bash
make dockerize
```

## Assignment 2: Buggy Project
The buggy_project.go file contains code with several issues, primarily related to unhandled errors, including database-related errors. You can find a solved version of the project in the repository, which addresses and resolves these errors.

## Assignment 3: E-Commerce Platform Design
This assignment outlines the design of an e-commerce platform consisting of three microservices:

### User Service: Manages user-related operations.
Product Service: Handles product information and operations.
Order Service: Manages orders and related functionalities.
Key Design Considerations:

### Authentication: Strategies for secure user authentication.
Scaling: Techniques to ensure the system can handle increased load.
Database Choice: MongoDB is selected due to its ability to handle unstructured data effectively, making it ideal for an e-commerce platform.

Documentation

API Documentation: Detailed API endpoints and usage are documented within the service.
Design Documentation: Available in the design directory, outlining the architecture and design decisions for the e-commerce platform.