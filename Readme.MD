# LRU Cache with Go Backend and React Frontend

This project implements a Least Recently Used (LRU) cache with a backend built on Golang and a frontend application developed using React. The backend provides API endpoints for interacting with the cache, while the frontend consumes these APIs to retrieve and set key/value pairs in the cache.

## Features

- **LRU Cache**: Implements a cache that stores Key/Value pairs with expiration. Keys are evicted from the cache after a set expiration time.
- **RESTful API**: Exposes GET and POST endpoints for retrieving and setting key/value pairs in the cache.
- **React Frontend**: Provides a user-friendly interface for interacting with the cache, allowing users to view cached keys and set new key/value pairs.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Node.js installed on your local machine.
- Go installed on your local machine.
- PostgreSQL installed and running locally.

## Installation

To set up the project, follow these steps:

1. Clone the repository: `https://github.com/saisandeepd7/LRU-Cache`
2. Navigate to the backend directory: `cd backend`
3. Install dependencies: `go mod tidy`
4. Set your postgreSQL Credentials
4. Start the backend server: `cd cmd\server\main.go` - `go run main.go`
5. Navigate to the frontend directory: `cd frontend`
6. Install dependencies: `npm install`
7. Start the frontend server: `npm start`

## Usage

Once the project is set up, you can access the frontend application at `http://localhost:3000`. From there, you can view cached keys and set new key/value pairs in the cache.


