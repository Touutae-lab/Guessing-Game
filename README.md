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
          - [ ] Very simple yes / no with password combination
          - [ ] Return "token"

        - `/guess`

            - [ ] Access to this endpoint needed to be authenticated via token returned from login

            - [ ] Guess the hidden number - if correct, return HTTP 201 and regenerate the number
       - `bonus`
            - [ ] Use of middleware for authentication
           
            - [ ] If we wanted to hide the guess data by not using GET, can we use other method to do so
