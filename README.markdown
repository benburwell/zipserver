[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/benburwell/zipserver)

# Usage

To run the server on port 8080 (set the `PORT` environment variable to override):

```sh
go build .
./zipserver
```

To get zipcode info using `curl` and [`jq`](https://stedolan.github.io/jq/):

```sh
curl --silent http://localhost:8080/zip/18101 | jq '.'
```

# Data

This project uses public-domain zipcode data from the [Zip Code Database Project](http://zips.sourceforge.net).

