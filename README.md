# coffee-supply-management

This project sets up a Hyperledger Fabric network, deploys chaincode, and starts a backend server and frontend application. Follow the steps below to get everything up and running.

## Prerequisites

Before you begin, ensure you have the following software installed on your machine:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Node.js](https://nodejs.org/) (which includes npm)
- [Go](https://golang.org/)
- [Hyperledger Fabric binaries](https://hyperledger-fabric.readthedocs.io/en/release-2.2/install.html)

## Setup the Network

1. **Clone the repository**:
    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. **Complete the prerequisite setup**:
    - Make sure Docker and Docker Compose are running.
    - Ensure you have all required binaries and configurations.

3. **Bring up the network**:
    ```sh
    ./setup.sh up
    ```

## Install and Deploy Chaincode

1. **Install and deploy the chaincode**:
    ```sh
    ./setup.sh deployCC
    ```

## Start the Backend Server

1. **Start the backend server through the debugger**:
    - Open your backend project in your preferred IDE (e.g., VS Code).
    - Set breakpoints if needed.
    - Start the debugger to run the server.

## Start the Frontend

1. **Navigate to the frontend directory**:
    ```sh
    cd frontend
    ```

2. **Install the dependencies**:
    ```sh
    npm install
    ```

3. **Start the frontend server**:
    ```sh
    npm start
    ```

Your Hyperledger Fabric network, backend server, and frontend application should now be running. Open your browser and navigate to `http://localhost:3000` to see the frontend application.

## Troubleshooting

- If you encounter any issues while bringing up the network, make sure all Docker containers are stopped and try again.
- For backend server issues, check the debugger console for detailed error messages.
- For frontend issues, check the browser console and terminal for error messages.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


