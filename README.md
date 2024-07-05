# Code Compiler Web App

## Overview

This web application allows users to execute code snippets in various programming languages through a user-friendly interface. It leverages the Piston API for executing the code securely and efficiently.

## Features

- **Multiple Language Support**: Execute code in languages like JavaScript, Python, Go, and Java.
- **Real-time Execution**: Submit code snippets and get immediate execution results.
- **User-friendly Interface**: A simple and intuitive web interface for code submission.

## Installation

### Prerequisites

- Git
- Go (for server-side)
- Node.js and npm (for client-side)

### Steps

1. Clone the repository: `git clone <repository-url>`
2. Navigate to the project directory: `cd <project-directory>`
3. Install server dependencies: `cd server && go get`
4. Install client dependencies: `cd ../client npm install`

## Usage

### Starting the Server

Navigate to the `server` directory and run: `go run main.go`

This will start the backend server.

### Running the Client

In the `client` directory, start the development server: `npm run dev`

Open your browser and navigate to `http://localhost:5173` to access the web application.

### Submitting Code for Execution

1. Select the programming language from the dropdown.
2. Enter your code snippet in the textarea provided.
3. Click the "Submit" button to execute the code.
4. The execution result will be displayed in the output area.
