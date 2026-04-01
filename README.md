# TN3270 client for [HF propagation data](https://www.hamqsl.com/solar.html) API

Simple tool for downloading HF propagation data from [hamqsl.com](https://www.hamqsl.com/solar.html) in XML format, and displaying the data in a [TN3270 terminal emulator](https://en.wikipedia.org/wiki/3270_emulator). It caches the current data, expiration time is 3 hours added to the time of generation as indicated in the XML file.

The client is written in Go and uses [go3270](https://github.com/racingmars/go3270) library.

To run the client, use `go run .` command.