Features

- Implementing the circuit breaker pattern. Prevents cascading failures and improves system resilience. Automatically halts operations temporarily when a threshold of failures is reached.
- Granularity: Each microservice has its own circuit breaker instance for isolated failure management.

- Modular Design: Each microservice is designed as a separate, independent module.
- Inter-Service Communication: Utilizes HTTP requests with context and circuit breaker integration.
- Logging: Standardized logging format across all services.
  Contextual Information: Includes key information like timestamps, service names, correlation IDs, and log levels.
  Logs are designed to be aggregated in a centralized logging system for easier access and analysis.

- GraphiQL GUI at /graphiql implemented for testing and development 
