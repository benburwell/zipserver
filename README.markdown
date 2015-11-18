# zipserver

A simple json api that returns location info about zipcodes.

## Usage

To run the server on port 8080:

```sh
go build .
./zipserver
```

To get zipcode info using `curl` and [`jq`](https://stedolan.github.io/jq/):

```sh
curl --silent http://localhost:8080/zip/18101 | jq '.'
```

## Data

This project uses public-domain zipcode data from the [Zip Code Database Project](http://zips.sourceforge.net).

