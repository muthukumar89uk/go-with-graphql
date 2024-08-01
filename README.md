# Go GraphQL CRUD Example

This repository demonstrates how to set up a Go project to perform CRUD (Create, Read, Update, Delete) operations using GraphQL.

## Features

- GraphQL server implementation in Go
- CRUD operations for a simple entity (e.g., User)
- Basic schema definition and resolvers

## Requirements

- Go 1.15 or higher
- Docker (optional but recommended for setting up a database)

## Getting Started

### Installation

1. **Clone the repository:**

      ```
    git clone https://github.com/muthukumar89uk/go-with-graphql.git
    ```
  Click here to directly [download it](https://github.com/muthukumar89uk/go-with-graphql/zipball/master). 

2. **Install Go dependencies:**

    ```sh
    go mod tidy
    ```

### Setup Database
1. **Locally:**

    Install PostgreSQL from the [official PostgreSQL website](https://www.postgresql.org/download/).

    Create a database named `graphql` and a user with appropriate privileges.

### Generate GraphQL Folder Strucure

  - For the generate graphql folder using the command:

    ```
    go run github.com/99designs/gqlgen init
    ```

### GraphQL Server

1. **Create the GraphQL schema:**

    Create a `schema.graphql` file in your project directory with the following content:

    ```graphql
    type User {
       id: Int!
       username: String!
       email: String!
       phone_number: String!
       password: String!
    }

    type DeleteStatus{
      iserror: Boolean!
      message: String
    }

    input NewUser {
       name: String!
       email: String!
       phoneNumber: String!
       password: String!
    }

    type Query {
       GetAllUsers: [User!]!
       GetOneUserById(id: Int!): User!  
    }

    type Mutation {
       createUser(input: NewUser!): User!
       updateUser(id: Int!, input: NewUser!): User!
       deleteUser(id: Int!): DeleteStatus
    }
    ```
    
2. **Create the resolver implementation:**

    Create a `graph` directory with `resolver.go` and the generated code:

    `resolver.go`:

    Run the following command to generate the necessary GraphQL code:

    ```sh
    go run github.com/99designs/gqlgen generate
    ```

    Implement the CRUD operations in the generated resolver file (`graph/schema.resolvers.go`):


### Run the Application

1. **Run the GraphQL server:**

    ```sh
    go run .
    ```

2. **Access the GraphQL Playground:**

    Open your browser and navigate to `http://localhost:8081/` to access the GraphQL Playground.


