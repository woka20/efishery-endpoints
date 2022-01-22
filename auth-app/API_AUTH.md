openapi: 3.0.0
info:
  version: 1.0.0
  title: Auth-App
  description: A simple collections of API for auth-app 

servers:
  # Added by API Auto Mocking Plugin
  - description: SwaggerHub API Auto Mocking
    url: https://virtserver.swaggerhub.com/woka20/Auth-App/1.0.0
  - url: http://localhost:8081/

components:
  securitySchemes:
    bearerAuth:            
      type: http
      scheme: bearer
      bearerFormat: JWT 

security:
  - bearerAuth: []  

paths:
  /hello:
    get:
      description: Returns string "Hello World!"
      responses:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: string
               

        default:
          description: Unregistered error/response

  /user:
    post:
      description: Insert new user to database and return its password
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                phone:
                  type: string
                  description: User's phone number
                name:
                  type: string
                  description: User's username
                role:
                  type: string
                  description: User's role
      responses:
        '200':
          description: Created!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: string
                    description: Related user password
        
        '500':
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message

        default:
          description: Unregistered error/response

  /token:
    post:
      description: Generate JWT Token for stored user with correct phone number and password
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              properties:
                  phone:
                    type: string
                    description: User's phone number
                  password:
                    type: string
                    description: User's password
      responses:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: string
                    description: Related user JWT Token
        
        '500':
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message

  /claims:
    get:
      description: Generate JWT Token for stored user with correct phone number and password
      security:
        - bearerAuth: []
      responses:
        '200':
          description: OK!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  data:
                    type: object
                    properties:
                      phone:
                        type: string
                        description: User's phone number
                      name:
                        type: string
                        description: User's username
                      role:
                        type: string
                        description: User's role
                      timestamp:
                        type: string
                        description: String formatted timestamp 
        
        '500':
          description: Internal Server Error!
          content:
            application/json:
              schema:
                type: object
                properties:
                  status:
                    type: integer
                    description: Response status code
                  message:
                    type: string
                    description: Error Message


> Written with [StackEdit](https://stackedit.io/).
