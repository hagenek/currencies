# Country Info CLI Application

The Country Info CLI application fetches and displays country information, including currencies, for specified regions. It supports fetching data for Europe or the entire world.

The app takes on argument, which is the region of the currencies

## Development

You need go installed:
https://go.dev/doc/install

You can just try the app with go run:

From root of project:

```bsh
    go run main.go world
```

### Prerequisites

- Go (version 1.21 or newer)
- Git

### Building and running as binary

```bsh
go build -o currencies
```

#### Using the program

Example:

```
./currencies world
```

or for europe:

```
./currencies europe
```
