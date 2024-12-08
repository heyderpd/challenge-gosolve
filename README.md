# Project gosolve task

This project contains backend and frontend applications to solve the gosolve challenge.
The backend folder has a rest api to resolve a fast search in a big file to find the index from desired value.
The frontend folder has a react app to easyle caller the backend, no authorization needed.

# Solution details

The solution uses the native function `sort.Search` because this does a binary search.
Is better than simply finder because the time complexity of binary search is O(log n).
For more details can found in https://pkg.go.dev/sort#Search

## Requirements

### Backend

- Move to backend folder
- Use specified golang version found in file `go.mod`
- Install golang dependencies

### Frontend

- Move to frontend folder
- Install NVM to use the specified node version OR manually use this node version
- Run `npm install` to install node dependencies

## Tests

- On backend folder, run `make tests` to run locally

## Usage

- On frontend folder, run `npm start` to run locally
- On backend folder, run `make start` to run locally
