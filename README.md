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
          - :white_check_mark: Very simple yes / no with password combination

          - :white_check_mark: Return "token"

        - `/guess`

            - :white_check_mark: Access to this endpoint needed to be authenticated via token returned from login

            - :white_check_mark: Guess the hidden number - if correct, return HTTP 201 and regenerate the number
            
       - `bonus`
            - :white_check_mark: Use of middleware for authentication
           
            - :white_check_mark: If we wanted to hide the guess data by not using GET, can we use other method to do so (POST)
