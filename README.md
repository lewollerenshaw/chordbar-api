# chordbar-api
Chordbar is a music theory tool designed to generate modal chord progressions based on user input. You are able to choose key, mode and the length of the progression you want to generate.

This is the api for chordbar which is written in go.

## Installation 
To install and run this solution, you will need to download the listed prerequisites below.

### Prerequisites
- You will need to install [golang](https://golang.org/dl/) on your machine to run this program.

- You will also need to installl [Gorilla Mux](https://github.com/gorilla/mux) by running this command.

```
go get -u github.com/gorilla/mux
```

### Running chordbar
Once you have installed the listed dependencies, clone the repository and run chordbar with the following commands

```
git clone https://github.com/lewollerenshaw/chordbar-api.git

cd chordbar-api

go build && ./chordbar-api
```

## Resources
[Gorilla Mux](https://github.com/gorilla/mux) implements a request router and dispatcher for matching incoming requests to their respective handler.


