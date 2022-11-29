# Guessing-Game
Golang simple Guessing game

Usage
```
go run main.go
```

listened at http://localhost:8000/


- Backend Task

    - API endpoints

        - `/login`
          - [x] Very simple yes / no with password combination

          - [x] Return "token"

        - `/guess`

            - [x] Access to this endpoint needed to be authenticated via token returned from login

            - [x] Guess the hidden number - if correct, return HTTP 201 and regenerate the number
            
       - `bonus`
            - [x] Use of middleware for authentication
           
            - [x] If we wanted to hide the guess data by not using GET, can we use other method to do so
