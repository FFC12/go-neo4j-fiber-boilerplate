## Fiber with Neo4j Boilerplate

This boilerplate provides a robust foundation for building web applications using the Fiber web framework, Neo4j database, JWT-based authentication, Swagger for API documentation, and Docker for containerization.

### Features

- **Fiber Web Framework**: Built on the fast and efficient Fiber framework for handling HTTP requests.
- **Neo4j Database**: Utilizes Neo4j, a powerful and flexible graph database.
- **JWT Authentication**: Implements basic JWT-based authentication for secure user access.
- **Swagger Documentation**: Includes Swagger for automatic API documentation generation.
- **Docker Support**: Containerize your application for easy deployment and scaling.


### Run
To run the application locally without Docker, follow these steps:

1) Run `go mod download`.
2) Start the app using `air`.

If you prefer to build and run it with Docker, make sure Docker is installed and execute the following command in your terminal:
```bash
docker-compose up -d
```

Additionally, ensure that your `.env` file is correctly configured. You can use the Neo4j Sandbox credentials for this purpose.

## License

MIT