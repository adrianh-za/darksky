Apple acquired Dark Sky on 31 March 2020 and this has affected the API which this repository relies on.

>Our API service for existing customers is not changing today, but we will no longer accept new signups. The API will continue to function through the end of 2021.

https://blog.darksky.net/dark-sky-has-a-new-home/


# darksky
Library written in Go to allow querying of weather conditions from Dark Sky

## Usage ##

You will need to create an (free) account at https://darksky.net/dev

Documentation for the Dark Sky API can be found at https://darksky.net/dev/docs

1) go get github.com/adrianh-za/go-darksky
2) browse to $/go/src/github.com/adrianh-za/go-darksky/examples
3) sudo -E go run main.go
4) run until end

## Acknowledgements ##

This project draws inspiration from <b><a href="https://github.com/mlbright" target="_blank">Martin-Louis Bright</a></b> and his Dark Sky Go library.

Please find his library for Dark Sky at https://github.com/mlbright/darksky
