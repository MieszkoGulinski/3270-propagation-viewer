# TN3270 viewer for [HF propagation data](https://www.hamqsl.com/solar.html) API

Simple tool for downloading HF propagation data from [hamqsl.com](https://www.hamqsl.com/solar.html) in XML format, and displaying the data in a [TN3270 terminal emulator](https://en.wikipedia.org/wiki/3270_emulator). It caches the current data, expiration time is 1 hour added to the time of generation as indicated in the XML file.

The viewer is written in Go and uses [go3270](https://github.com/racingmars/go3270) library.

To run the viewer, use `go run .` command. Then, you need to connect to it using a TN3270 client, e.g. `x3270` or `c3270`.