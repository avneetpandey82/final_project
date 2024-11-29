# Go Project with MongoDB Integration and User Registration Endpoint

This project demonstrates how to create a Go application that connects to a MongoDB database and handles user registration requests. The user data is submitted through a web form, and the data is parsed and stored in a MongoDB database. Additionally, the application handles basic authentication and CORS setup.

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/dl/) (1.18+)
- [MongoDB](https://www.mongodb.com/) (or MongoDB Atlas for cloud setup)
- [Docker](https://www.docker.com/) (optional, for containerization)

## Project Overview

1. **Backend:**
   - Go handles the backend logic.
   - A REST API endpoint is created to handle user registration requests.
   - MongoDB is used to store user data (like name, email, and password).
2. **Frontend:**

   - An HTML form built with Bootstrap allows users to submit their details (name, email, password, etc.).
   - JavaScript is used to send the form data to the backend via a POST request.

3. **Basic Authentication:**
   - Basic validation is implemented for the user input to ensure required fields are provided.
   - MongoDB database is used to store the user details.

## Project Structure

### 1. **Backend - Go Setup**

### Database Connection and API Endpoint

The Go server connects to MongoDB and listens for POST requests at `/register`. The request is handled by the `handleUserRegister` function which:

- Parses incoming JSON data.
- Inserts the user data into a MongoDB database.
- Sends a response to the client confirming the signup.

### 2. **Frontend - HTML Form with Bootstrap**

We created a simple HTML form using Bootstrap to collect user registration data. This form will send the user input to the backend API using JavaScript's `fetch` function.

### 3. **Handling Authentication**

In the `handleUserRegister` function, basic validation is performed. However, we can add more checks for authentication or validation if needed. For example, we might want to validate if the email already exists in the database.

### 4. **Running the Project**

1. **Docker Setup:**
   You can build a Docker image for this Go application and run it in a container.

   ```bash
   docker build -t user-registration .
   docker run -p 8080:8080 user-registration
   ```
