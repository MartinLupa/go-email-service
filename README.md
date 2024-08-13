# Go Email Service

## Overview

This is a simple email service built with Go, designed to send emails using different providers such as Mailgun and SparkPost. It exposes a RESTful API that allows clients to send emails by making HTTP POST requests.

## Features

- Supports multiple email providers (Mailgun, SparkPost).
- Simple RESTful API for sending emails.
- Environment-based configuration using `.env` file.
- Docker support for easy deployment.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.
- [Docker](https://www.docker.com/) installed (for Docker-based deployment).
- Environment variables set in a `.env` file.

### Installation

1. **Clone the repository**:

    ```bash
    git clone https://github.com/MartinLupa/go-email-service.git
    cd go-email-service
    ```

2. **Set up environment variables**:

    Create a `.env` file in the root of the project with the following content:

    ```env
    PORT=8080
    MAILGUN_API_KEY=your-mailgun-api-key
    MAILGUN_DOMAIN=your-mailgun-domain
    SPARKPOST_API_KEY=your-sparkpost-api-key
    ```

3. **Build and run the service**:

    You can either run the service directly using Go or use Docker.

    - **Using Go**:

      ```bash
      go run main.go
      ```

    - **Using Docker Compose**:

      Build the Docker image:

      ```bash
      docker-compose up
      ```

## Usage

### API Endpoint

- **POST** `/send-email`

### Request Payload

To send an email, make a POST request to `http://localhost:8080/send-email` with the following JSON payload:

```json
{
    "subject": "Some important email",
    "body": "Some important information in some important email",
    "to": "recipient@example.com"
}
