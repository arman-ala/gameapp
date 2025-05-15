# Game Application

# use-cases

## Game use-cases
### each game has a given number of questions
### the difficulty of the questions is increased by the number of questions answered correctly
### the difficulty level of the questions are "easy", "medium", "hard"
### game winner is determined by the number of questions answered correctly
### each game belongs to a specific category: "sports", "geography", etc.

## User use-cases
### Login
- user can login by phone number and password

### Register
- user can register by phone number

# entity
## User
- id
- phone_number
- password
- avatar
- name

## Game
- id
- category
- Question List
- Players
- Winner

## Question
- id
- question
- answer list
- correct answer
- difficulty
- category
