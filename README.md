# ushortener

URL Shortening Service

## Installation

- Clone this repository onto your local machine

```
git clone https://github.com/nemzyxt/ushortener.git
```

- Navigate to the project directory and install required packages

```
go get
```

## Usage

Start the server

```
go run server.go
```

You can now use the service via its endpoints

## Example

Request:

```
POST http://localhost:1234/shorten
Content-Type: application/json

{
    "longURL": "https://github.com/nemzyxt/ushortener"
}
```

Sample response:

```
{
    "shortURL": "http://localhost:1234/AO_2paiIg"
}
```

