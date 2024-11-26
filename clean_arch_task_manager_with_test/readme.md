## Prerequisites

Before you begin, make sure you have met the following requirements:
- Go 1.x installed on your machine (where x is the latest version)
- MongoDB installed and running locally or accessible via network

## Getting Started

Follow these steps to set up your development environment:

1. **Clone the repository**

```sh
git clone https://yourrepositoryurl/clean_arch_task_manager.git
cd clean_arch_task_manager
```

2. **Set up environment variables**

Copy the `.env.example` file to a new file named `.env` and update the variables to match your local setup. The `.env` file should include the following variables:

- DB_NAME: The name of your MongoDB database
- MONGO_URI: uri to your mongodb server
- JWT_SECTRE: secret key for jwt


```sh
cp delivery/.env.example delivery/.env
```

3. **Install dependencies**

```sh
go mod tidy
```

4. **Run the application**

```sh
go run delivery/main.go
```

### Project Structure
The project is structured as follows to adhere to Clean Architecture principles:

- db/: Database connection setup
- delivery/: Entry point of the application and HTTP server setup
- domain/: Core business logic and models
- repository/: Data access layer
- usecase/: Application-specific business rules
- docs/: API documentation and other documentation resources