# image-compressor

# Project Setup Guide

This repository contains a Go project with RabbitMQ integration, utilizing Gorilla Mux for routing, Gorm for database operations, and Logrus for logging,nfnt/resize for image compressing

## Prerequisites

Before running the project, ensure that you have the following installed:

- [Go](https://golang.org/doc/install)
- [Docker](https://www.docker.com/get-started)
- [Erlang](https://www.erlang.org/downloads)
- [RabbitMQ](https://www.rabbitmq.com/download.html)

## Installation Steps

1. Install Gorilla Mux:
    ```bash
    go get -u github.com/gorilla/mux
    ```

2. Install Gorm and SQLite Driver:
    ```bash
    go get -u github.com/jinzhu/gorm
    go get -u gorm.io/gorm
    go get -u gorm.io/driver/sqlite
    ```

3. Install Logrus:
    ```bash
    go get github.com/sirupsen/logrus
    ```

4. Install RabbitMQ:
    - Install Erlang: [Download Erlang](https://www.erlang.org/downloads)
    - Install RabbitMQ: [Download RabbitMQ](https://www.rabbitmq.com/download.html)

    After installing RabbitMQ, run the following commands:
    ```bash
    # Enable RabbitMQ Management Plugin
    rabbitmq-plugins enable rabbitmq_management
    ```

    Access RabbitMQ Management Console: [http://localhost:15672/](http://localhost:15672/)
    - Username: guest
    - Password: guest

5. Install RabbitMQ Go Library:
    ```bash
    go get github.com/rabbitmq/amqp091-go
    ```

6. Install nfnt package for image compression:
    ```bash
    go get github.com/nfnt/resize
    ```

## Running the Project

1. Run RabbitMQ:
    - Ensure that RabbitMQ is running.

2. Execute the SQL Script:
    - Use a MySQL client or MySQL Workbench to execute the SQL script in the `database.sql` file.
    - username: root@localhost
    - password: root
    - Database: Zocket

### Producer
1. Open a terminal.
2. Navigate to the producer/cmd directory.
    ```bash
    cd producer/cmd
    ```
3. Run the producer application.
    ```bash
    go run main.go
    ```

### Consumer
1. Open another terminal.
2. Navigate to the consumer/cmd directory.
    ```bash
    cd consumer/cmd
    ```
3. Run the consumer application.
    ```bash
    go run main.go
    ```

## Interacting with the System

### Using Postman

1. Open Postman.
2. Set the endpoint to `localhost:8080/product`.
3. Choose one of the following request payloads and enter it into the request body:

#### Request Payloads:

```json
{
  "user_id": 5,
  "product_name": "Apple Watch",
  "product_description": "A wearable smartwatch that allows users to accomplish a variety of tasks, including making phone calls, sending text messages, and reading emails.",
  "product_images": "https://raw.githubusercontent.com/anushka2911/images/main/uploads/Iwatch.jpg",
  "product_price": 12000.2
}

{
  "user_id": 7,
  "product_name": "Iphone 13",
  "product_description": "6.1-inch (155 mm) display with Super Retina XDR OLED technology at a resolution of 2532×1170 pixels and a pixel density of about 460 PPI with a refresh rate of 60 Hz",
  "product_images": "https://raw.githubusercontent.com/anushka2911/images/main/uploads/phone.jpg",
  "product_price": 80000
}

{
  "user_id": 12,
  "product_name": "Watch",
  "product_description": "6.1-inch (155 mm) display with Super Retina XDR OLED technology at a resolution of 2532×1170 pixels and a pixel density of about 460 PPI with a refresh rate of 60 Hz",
  "product_images": "https://raw.githubusercontent.com/anushka2911/images/main/uploads/watch.jpg",
  "product_price": 25000
}

{
  "user_id": 15,
  "product_name": "Headphone",
  "product_description": "Infinity - JBL Tranz 710, 72 Hrs Playtime with Quick Charge, Wireless On Ear Headphone with Mic, Deep Bass, Dual Equalizer, Bluetooth 5.0 with Voice Assistant",
  "product_images": "https://raw.githubusercontent.com/anushka2911/images/main/uploads/headphone.jpg",
  "product_price": 7000
}

{
  "user_id": 333,
  "product_name": "Headphone",
  "product_description": "Infinity - JBL Tranz 710, 72 Hrs Playtime with Quick Charge, Wireless On Ear Headphone with Mic, Deep Bass, Dual Equalizer, Bluetooth 5.0 with Voice Assistant",
  "product_images": "https://raw.githubusercontent.com/anushka2911/images/main/uploads/headphone.jpg,https://raw.githubusercontent.com/anushka2911/images/main/uploads/Iwatch.jpg,https://raw.githubusercontent.com/anushka2911/images/main/uploads/phone.jpg",
  "product_price": 7000
}
```

## Hosting Images

Images are hosted using [https://gaac.vercel.app/](https://gaac.vercel.app/). 

