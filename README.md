# HTTP protocol in go

Implementing HTTP protocol in go without using any external library

Uses the net/tcp package to create the server

## What does it do?

The net/tcp creates a tcp sever as we server at a port and then accept the connection, we keep the connection in a loop and spawn go routines so that multiple connections can be done

Http methods have been split and body is served when we hit localhost in the browser

## What can it do currently?

* Can server multiple endpoints like /, /hello and so on 
* Implemented query parameters
