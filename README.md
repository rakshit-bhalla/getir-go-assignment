# getir-go-assignment

<p align="center">
  <a href="https://go.dev/" target="blank"><img src="https://miro.medium.com/max/1400/1*Ifpd_HtDiK9u6h68SZgNuA.png" width="320" alt="Go Logo" /></a>
</p>

  <p align="center">REST API to display in-memory and no-sql db setup in <a href="https://go.dev/" target="_blank">Go</a> language for building efficient server-side applications.</p>

## Description

- REST API with MVC standards using [Go](https://go.dev/) language.
- Documentation has been implemented with [Swagger Open Api](https://swagger.io/specification/).
- [Mongo database](https://www.mongodb.com/) (No SQL) has been used.
- Map for In-memory has been used.

## To Do's
- Add Documentation using Swagger
- Add health checks
- Improve logging and setup log file generation
- Add more testcases testing

## Running the app

- Clone the repository
  ```bash
  $ git clone git@github.com:rakshit-bhalla/getir-go-assignment.git
  ``` 
- Enter the folder
  ```bash
  $ cd getir-go-assignment
  ``` 
- Add the configurations and run it using the following command
  ```bash
  $ go run ./src/main.go 
  ``` 
- Visit this [http://localhost:3000/records](http://localhost:3000/records)
- Base Documentation can be checked here
  ```
  openapi: 3.0.0
  info:
    title: Sample API in go
    version: 0.1.9

  servers:
    - url: https://getir-go-assignment.herokuapp.com/in-memory
      description: URL to interact with the app

  paths:
    /records:
      post:
        summary: Returns a list of records.
        description: Get records created between given startDate and endDate and having totalCount between given minCount, maxCount.
        requestBody:
          description: Record-Query
          required: true
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/RecordQuery'
        responses:
          '200': # status code
            description: A JSON array of records
            content:
              application/json:
                schema: 
                  $ref: '#/components/schemas/RecordsResponse'
          '500':    # status code
            description: Error message
            content:
              application/json:
                schema:
                  type: string
                  value: internal server error

    /in-memory:
      get:
        summary: Returns the key-value pair if exists in the in-memory db or else error message.
        parameters:
          - in: query
            name: key
            schema:
              type: string
            description: The key for which value needs to be searched
        responses:
          '200':    # status code
            description: The key-value pair
            content:
              application/json:
                schema: 
                  $ref: '#/components/schemas/KeyValuePair'
          '404':    # status code
            description: Error message
            content:
              application/json:
                schema:
                  type: string
                  value: not found
          '500':    # status code
            description: Error message
            content:
              application/json:
                schema:
                  type: string
                  value: internal server error
          '400':    # status code
            description: Error message
            content:
              application/json:
                schema:
                  type: string
                  example: invalid param startDate
      post:
        summary: Saves the key-value pair in the in-memory db and return it or else give error message.
        requestBody:
          description: key-value pair
          required: true
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/KeyValuePair'
        responses:
          '200':    # status code
            description: The key-value pair
            content:
              application/json:
                schema: 
                  $ref: '#/components/schemas/KeyValuePair'
          '500':    # status code
            description: Error message
            content:
              application/json:
                schema:
                  type: string
                  value: internal server error

  components:
    schemas:
      RecordQuery:
        type: object
        properties:
          startDate:
            type: string
            format: date
            example: 2016-01-26
          endDate:
            type: string
            format: date
            example: 2018-02-02
          minCount:
            type: integer
            example: 2700
          maxCount:
            type: integer
            example: 3000

      RecordsResponse:
        type: object
        properties:
          code:
            type: integer
            example: 0
          msg:
            type: string
            example: Success
          records:
            type: array
            items:
              $ref: '#/components/schemas/Record'

      RecordsResponse:
        type: object
        properties:
          key:
            type: string
            example: TAKwGc6Jr4i8Z487
          createdAt:
            type: string
            format: date-time
          totalCount:
            type: integer
            example: 2800

      KeyValuePair:
        type: object
        properties:
          key:
            type: string
            example: active-tabs
          value:
            type: string
            example: getir
        # Both properties are required
        required:  
          - key
          - value
  ```