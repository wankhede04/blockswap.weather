# Blockswap Weather

## Description
A smart contract that will allow any ETH address to perform a registration function, and they can leave any time so they will either have lifecycle status: unregistered, registered or resigned.

Try running some of the following tasks:
```shell
npx hardhat help
npx hardhat test
REPORT_GAS=true npx hardhat test
npx hardhat node
npx hardhat run scripts/registration.ts --network <NETWORK>
```

# Weather Service

The Weather Service App is a Go application that allows users to report weather information. It includes features such as user authentication, rate limiting, and storing weather reports in a database. The app is built using the Gin web framework and integrates with a weather API.

## Installation
To install and run the Weather Service, follow these steps:

1. Make sure you have PostgreSQL installed and running on your system.
2. Clone the repository: git clone https://github.com/wankhede04/blockswap.weather.git
3. Navigate to the project directory: cd blockswap.weather
4. Set up the environment variables. Create a .env file in the root directory and define the necessary variables.
5. Install the necessary dependencies: go mod tidy
6. Build the application: go build
7. Start the weather service: go run main.go

The weather service should now be running and accessible at http://localhost:8080.

## Endpoints
- POST/report-weather
    - Request Body:
    - JSON object with the following properties:
        - address (string): The name of the location for which weather information is being reported.
        - report (string): The temperature in Celsius.
        - signature (string): The humidity percentage.
    - Response:
        - Status Code: 201 (Created)
        - Body: Weather report submitted

## Architecture and Flow
The weather service is built using the Gin framework and follows a client-server architecture. Here's a high-level overview of the flow:

1. The server starts by connecting to the PostgreSQL database and performing the necessary migrations to create the required tables.
2. The server creates an instance of the Membership, which handles the registration of new memberships.
3. The server creates an instance of the WeatherService, which handles weather reporting and authentication.
4. The server starts a watcher goroutine, which periodically checks for changes in membership status from contracts deployed on blockchain and updates them accordingly.
5. The server sets up the necessary routes using the Gin framework.
6. When a registration event is raised by the contract, the server creates a new Membership record in the database.
7. When a weather report post request is received, the server follows this architecture:
   - First it validates the membership status from database and verify signature of member with data
   - Then rate limit middleware checks for time limitations and window which needs to accept post request
   - Finally server will submit weather report in database with member ship ID and update lastCall in DB
8. The watcher goroutine periodically fetches events from the contracts and updates their status based on simulated events.
The server responds to the client with success or error messages for each request.

## Future Prospect
1. Rate limiting in distributed environments: Address the challenges of rate limiting in a distributed environment where multiple pods are deployed. This can be achieved by implementing a cache system (e.g., Redis) or handling rate limiting at the load balancer level.
2. Security enhancements: Implement additional security measures such as input validation, request throttling, and protection against common web vulnerabilities (e.g., CSRF, XSS).
3. Error handling and logging: Enhance error handling by providing meaningful error messages and implementing a robust logging mechanism to track application events and troubleshoot issues effectively.