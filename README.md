# Task Queues with Go

This project demonstrates a task queue system implemented in Go, focusing on the example of sending emails using Redis.

## Overview

The goal of this project is to showcase a simple task queue system where emails are sent asynchronously. The architecture involves a Go program that adds email sending tasks to a Redis queue, and a worker that processes these tasks in the background.

## Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- Redis Server

## Getting Started

1. **Run Redis Server:**
   Ensure that the Redis server is running. You can use Docker to start a Redis container:
   ```bash
   docker run -p 6379:6379 --name my-redis-container -d redis
