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

# Response

Your response will look something like this:

```json
{
	"latitude": 40.602847,
	"longitude": -75.47022,
	"city": "Allentown",
	"state": "PA"
}
```

If the zipcode you request isn't in the database, you'll get a `404` status with an empty body.

Add a `?distance=` parameter to your request to see how far away two zipcodes are. The response format will be the same as above with an additional `distance` key that contains info about the second zipcode as well as `miles` and `kilometers` keys. If the second zipcode can't be found, the `distance` key won't be included.

# Data

This project uses public-domain zipcode data from the [Zip Code Database Project](http://zips.sourceforge.net).

