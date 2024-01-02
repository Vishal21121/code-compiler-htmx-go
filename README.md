# Code Compiler Web App

## Description

This project uses the Piston API to execute code in various programming languages. The language and the code to be executed are provided in the request body.

## Installation

To install the project, follow these steps:

1. Clone the repository: `git clone <repository-url>`
2. Navigate into the project directory: `cd <project-directory>`
3. Install the dependencies: `go get`

## Usage

To use the project, you need to send a POST request with a body containing the language and the code to be executed. The language should be a string and the code should be an array of objects with a `Content` property.

Example:

```json
{
  "lang": "python",
  "code": [
    {
      "Content": "print('Hello, World!')"
    }
  ]
}
```
