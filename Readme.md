# minesweeper

Minesweeper API RESTful made in Golang

## Description of the solution

- For the structure I chose Clean Architecture and the example was taken from: [Clean Architecture in Go](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1).
- For the deploy I created a Docker image, pushed it to [dockerhub](https://hub.docker.com/repository/docker/kosegor/minesweeper), pulled it from my EC2 instance with Golang and Docker installed and ran it.
- I created some tests, the most important ones. Just for showing my testing skills.
- I created a [Library in Java 8 with API Rest Client](https://github.com/egorkos/minesweeper-restlient-library) for interactions with this API.

## Endpoints and usage

### Ping

- Description: check the server is online
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/ping`
- Rest verb: GET
- Possible responses:

  | Http Status Code | Description    |
  | :--------------- | :------------- |
  | 200              | Returns `pong` |
  | 500              | Server Error   |

### Get

- Description: return a saved Game
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/game/{id}`
- Rest verb: GET
- Possible responses:

  | Http Status Code | Description                   |
  | :--------------- | :---------------------------- |
  | 200              | Returns a saved [Game](#Game) |
  | 400              | Bad Request                   |
  | 404              | Not Found                     |
  | 500              | Server Error                  |

### List Games

- Description: return a list of saved Games
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/games`
- Rest verb: GET
- Possible responses:

  | Http Status Code | Description                           |
  | :--------------- | :------------------------------------ |
  | 200              | Returns a list of saved [Game](#Game) |
  | 500              | Server Error                          |

### Create Game

- Description: create a new Game
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/game`
- Rest verb: POST
- Request Body expected:
  - `rows`: Game rows quantity (min:0, max:50)
  - `cols`: Game cols quantity (min:0, max:50)
  - `mines`: Game mines quantity (min: 0, max:rows\*cols-1)
  - `{"rows":1, "cols":3, "mines":1}`
- Possible responses:

  | Http Status Code | Description                 |
  | :--------------- | :-------------------------- |
  | 200              | Returns a new [Game](#Game) |
  | 400              | Bad Request                 |
  | 500              | Server Error                |

### Reveal Cell

- Description: reveal a Cell
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/games/{id}/reveal`
- Rest verb: POST
- Request Body expected:
  - `row`: Row to reveal
  - `col`: Col to reveal
  - `{"row":0, "col":0}`
- Possible responses:

  | Http Status Code | Description                                               |
  | :--------------- | :-------------------------------------------------------- |
  | 200              | Returns a saved [Game](#Game) with revealed [Cell](#Cell) |
  | 400              | Bad Request                                               |
  | 404              | Not Found                                                 |
  | 500              | Server Error                                              |

### Flag Cell

- Description: flag a Cell
- URI: `ec2-18-191-183-190.us-east-2.compute.amazonaws.com:8080/games/{id}/flag`
- Rest verb: POST
- Request Body expected:
  - `row`: Row to flag
  - `col`: Col to flag
  - `{"row":0, "col":0}`
- Possible responses:

  | Http Status Code | Description                                              |
  | :--------------- | :------------------------------------------------------- |
  | 200              | Returns a saved [Game](#Game) with flagged [Cell](#Cell) |
  | 400              | Bad Request                                              |
  | 404              | Not Found                                                |
  | 500              | Server Error                                             |

### Game

#### Model

- id: game id
- startTime: start date and time
- finishTime: finish date and time
- rows: rows quantity
- cols: cols quantity
- mines: mines quantity
- cellsRevealed: cells revealed quantity
- status: game [Status](#Status)
- grid: game board -> matrix of [Cell](#Cell)

#### Json Example

    {
        "id": 1,
        "start_time": "2020-01-21T18:20:54.18293094Z",
        "finish_time": "0001-01-01T00:00:00Z",
        "rows": 1,
        "cols": 3,
        "mines": 1,
        "cells_revealed": 0,
        "game_status": 2,
        "grid": [
            [
                {
                    "mine": false,
                    "revealed": false,
                    "flagged": false,
                    "mines_around": 0
                },
                {
                    "mine": false,
                    "revealed": false,
                    "flagged": false,
                    "mines_around": 1
                },
                {
                    "mine": true,
                    "revealed": false,
                    "flagged": false,
                    "mines_around": 0
                }
            ]
        ]
    }

### Cell

#### Model

- mine: bool mine indicator
- revealed: bool revealed cell indicator
- flagged: bool flagged cell indicator
- minesAround: quantity of mines around Cell

#### Json Example

    {
        "mine": true,
        "revealed": false,
        "flagged": false,
        "mines_around": 0
    }

### Status

| index | description |
| :---- | :---------- |
| 0     | Win         |
| 1     | Loose       |
| 2     | Running     |
