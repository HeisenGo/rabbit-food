# Rabbit Food 🥕🍔
# A Go Socket-Based Online Food Ordering System
__Rabbit Food__ is a socket-based online food ordering system developed in Go. It provides a platform for users to order food from various restaurants, manage their profiles, and handle financial transactions securely. Restaurant owners can register and manage their establishments, while administrators have access to an admin panel for system management. The system leverages TCP connections and a custom protocol based on JSON and byte encoding for efficient client-server communication.

## Key Features
- __User Registration and Profile Management:__ Users can create accounts, update their profiles, and manage personal information./n
- __Wallet Functionalities:__ Users can add funds to their wallets, facilitating secure online payments for food orders.
- __Food Ordering:__ Browse menus, select items, and place orders from registered restaurants.
- __Restaurant Creation and Management:__ Restaurant owners can create and manage their restaurant profiles, update menus, and track orders.
- __Admin Panel:__ Administrators have access to an admin panel for system management, user management, and monitoring.
- __CLI Client:__ A command-line interface (**CLI**) client allows users to manage orders, perform transactions, and interact with the system through a terminal interface.

## Architecture
The system is designed with a modular architecture, separating the server and client components for scalability and maintainability. The server component handles client connections, processes requests, interacts with the database, and manages the overall system logic. The client component provides a user-friendly interface (CLI) for users to interact with the system.
Communication between the server and clients is facilitated through TCP connections, ensuring efficient and reliable data transfer. A custom protocol is implemented to standardize the message formats and ensure seamless communication between the components.

## Server Side Project Structure

    ├── api

    │   └── tcp

    ├── cmd

    │   └── server

    ├── config

    ├── internal

    │   ├── errors

    │   │   └── users

    │   ├── models

    │   │   └── user

    │   ├── protocol

    │   └── server

    │       └── handlers

    ├── pkg

    │   ├── adapters

    │   │   └── storage

    │   │       ├── entities

    │   │       └── mappers

    │   ├── jwt

    │   ├── logger

    │   └── utils

    │       └── users

    ├── scripts

    ├── services

    └── test

    │   └──  users


### Explanation:
- **api**: Contains the TCP-related functionality.
- **cmd**: Houses the server command files.
- **config**: Stores configuration files.
- **internal**: Holds internal packages and modules.
  - **errors**: Handles errors, specifically related to users.
  - **models**: Includes user-related models.
  - **protocol**: Contains protocol-related logic.
  - **server/handlers**: Contains handlers for server requests.
- **pkg**: Contains packages used throughout the project.
  - **adapters/storage**: Includes storage-related adapters and mappers.
    - **entities**: Specific entities related to storage.
    - **mappers**: Data mappers for storage.
  - **jwt**: Logic related to JWT.
  - **logger**: Utilities for logging.
  - **utils/users**: Utility functions specifically for users.
- **scripts**: Includes any scripts used in the project.
- **services**: Contains service logic.
- **test/users**: Contains test files specifically for user-related functionality.



## Getting Started
To get started with **Rabbit Food**, follow the instructions in the following. The guide will walk you through the process of setting up the project, configuring the necessary dependencies, and running the server and client components.


The project's database schema is represented by an *Entity Relationship Diagram (ERD)*. You can view the **ERD** **[HERE](https://drive.google.com/file/d/1gNUd9nl6XqafS7znTZN26XCL0raL7jAb/view?usp=sharing)**.

**Note:** For better resolution of the ERD, open it with [draw.io](https://app.diagrams.net/#G176KUF2eQaoieLKbGChO8VnN-5od6StkW) by selecting "Open with" and choosing "draw.io" from the options. You may need to authorize the site to view the diagram.
### Contributing
We welcome contributions from the community! If you'd like to contribute to Rabbit Food, please read our Contributing Guidelines for information on how to get involved.
