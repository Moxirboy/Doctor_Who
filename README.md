# Curify App

## Introduction

Curify is a powerful application built with Golang that helps you manage and organize your data seamlessly. This readme provides instructions on how to set up and run the Curify app on your local environment.

## Prerequisites

Before getting started, ensure that you have the following installed on your machine:

- Golang
- Docker (optional but recommended)

## Installation

1. **Clone the Curify repository to your local machine:**

    ```bash
    git clone https://github.com/your-username/curify.git
    ```

2. **Navigate to the project directory:**

    ```bash
    cd curify
    ```

3. **Create a copy of the `.env.e` file and name it `.env`:**

    ```bash
    cp .env.e .env
    ```

    Edit the `.env` file to set the required environment variables for your setup.

4. **Run the following command to perform database migrations after make start:**

    ```bash
    make migration-up
    ```

    This command will apply the necessary database migrations.

## Usage

Once you have completed the installation steps, you can start the Curify app with the following command:

```bash
make start
