<h1 align="center">
  Chat RTC
</h1>

<br />

Real time server for chatting with Golang. It is a real time service which
broadcasts the events between clients.

<br />

## HTTP

For ```HTTP``` handler we use golang gin. Handler will get user requests on ```/ws```
and sets the configurations needed in order to connect with ```melody``` websocket handler.

## Web Socket

For web sockets we use [```melody```](https://github.com/olahol/melody). Melody implements
the real time handlers so we can make connection between our clients.
