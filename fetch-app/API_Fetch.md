
openapi: 3.0.0
info:
  version: 1.0.0-oas3
  title: Fetch-App
  description: Collections of API for fetch-app
servers:
  - url: http://localhost:5000/
security:
  - bearerAuth: []
paths:
  /commodityList:
    get:
      description: Get list of commodities data with additional field price in USD currency
      response:
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
                    type: array
                    description: >-
                      List of all commodities data with additional field of
                      converted price in USD currency
                    properties:
                      uuid:
                        type: string
                        description: Commodity data ID
                      komoditas:
                        type: string
                        description: Commodity name/type
                      area_provinsi:
                        type: string
                        description: Commodity sell area, province level
                      area_kota:
                        type: string
                        description: Commodity sell area, city level
                      size:
                        type: string
                        description: Commodity's size/amount
                      price:
                        type: string
                        description: Commodity's each item's price in rupiah
                      tgl_parsed:
                        type: string
                        description: Commodity's date time
                      timestamp:
                        type: string
                        description: Timestamp
                      usd_price:
                        type: string
                        description: Commodity's each item's price in USD
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
      responses:
        default:
          description: Default response
  /aggregateData:
    get:
      description: >-
        Get list of aggregated and clean commodities data based on area_province
        and weekly data with max, min, avg, median of data
      response:
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
                    type: array
                    description: >-
                      List of all commodities data with additional USD currency
                      price field
              
                    item:
                      area_provinsi:
                        type: string
                        description: Commodity sell area, province level
                      profit:
                        type: map[string, map[string, int]]
                        description: >-
                          Contains all profits (price * size) at Nth week of the
                          year for it's province (map[tahun, map[minggu,
                          price*size]])
                      max_profit:
                        type: float
                        description: highest weekly profit of the province
                      min_profit:
                        type: float
                        description: lowest weekly profit of the province
                      avg_profit:
                        type: float
                        description: average weekly profit of the province
                      median_profit:
                        type: float
                        description: median of weekly profit of the province
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
      responses:
        default:
          description: Default response
  /token:
    get:
      description: Generate Token for correct phone and password
      security:
        - bearerAuth: []
      response:
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
                      data:
                        type: string
                        description: valid generated token
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
      responses:
        default:
          description: Default response
  /claims:
    get:
      description: >-
        Give Information in including name,phone,role, and timestamp of JWT
        Token for stored user with correct phone number and password
      security:
        - bearerAuth: []
      response:
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
      responses:
        default:
          description: Default response
