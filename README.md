# Pismo Technical Challenge - Transaction System

## Description:
This is a transaction system project developed as part of the technical challenge proposed by Pismo.
The main goal is to create an application that manages financial transactions, allowing the creation, listing, and querying of accounts and transactions.



## Main Features:
- Account Registration: Enables the creation of accounts associated with a document (CPF or CNPJ).
- Transaction Registration: Allows the creation of financial transactions associated with an account, including operations such as purchase, payment, withdrawal, among others.
- Account and Transaction Query: Offers the ability to query detailed information about accounts and transactions, facilitating the analysis and monitoring of financial activities.



## Project Structure:
The project follows a modular structure, separating responsibilities into different packages.
The main packages include:

- `account`: Contains business logic related to accounts, including services, database storage, and unit tests.

- `cmd`: This directory houses the main entry point of the application, which can be executed to start the service.

- `config`: Provides the structure for managing configurations and associated tests.

- `docker-compose.yaml`: Docker composition file that facilitates running services in containers.

- `infrastructure`: Responsible for handling the infrastructure layer, such as initializing and managing database connections.

- `internal`: Contains internal packages necessary for the application to function.

- `migrations`: SQL files used for database migrations.

- `pkg`: Modules and packages that are shared across different parts of the project.

- `transaction`: Deals with logic related to transactions, including services, database storage, and unit tests.



## How to Use:
### Installation:
  To install project dependencies, run the following command
  ```bash
  make ensure
  ```
### Local Execution:
  ```bash
  make build
  make run
  ```
### Tests:
  Run unit tests with the following command:
  ```bash 
  make test
  ```
### Linting:
  Perform linting checks with the following command:
  ```bash 
  make lint
  ```
### Docker Execution:
  To build, run, and test the application using Docker, use the following commands:
  ```bash
  make docker/build
  make docker/run
  make docker/migrate/up
  ```

