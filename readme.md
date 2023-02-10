<p align="center">
  <a href="https://gateway.careers360.com/" target="blank"><img src="https://production-cnext.s3.ap-south-1.amazonaws.com/vision.jpg" width="500" alt="API Gateway Logo" /></a>
</p>

[circleci-image]: https://img.shields.io/circleci/build/github/nestjs/nest/master?token=abc123def456
[circleci-url]: https://circleci.com/gh/nestjs/nest

  <p align="center">API Gateway for Careers360 built in Golang with <a href="https://www.krakend.io/" target="_blank">KrakenD </a> framework, Codename: <b>Vision</b></p>
    <p align="center">
<img src="https://img.shields.io/cirrus/github/careers-360/cnext-gateway?style=plastic" alt="Discord"/>
<img src="https://img.shields.io/badge/golang-%3E1.01-brightgreen" alt="Golang Version" />
<img src="https://img.shields.io/badge/docker-yes-brightgreen" alt="Requires Docker" />
</p>

## Description

Consumers of REST API content (specially in microservices) often query backend services that weren't coded for the UI implementation. This is of course a good practice, but the UI consumers need to do implementations that suffer a lot of complexity and burden with the sizes of their microservices responses.

Vision is basd on KrakenD API Gateway builder and proxy generator that sits between the client and all the source servers, adding a new layer that removes all the complexity to the clients, providing them only the information that the UI needs. Vision acts as an aggregator of many sources into single endpoints and allows you to group, wrap, transform and shrink responses. Additionally it supports a myriad of middlewares and plugins that allow you to extend the functionality, such as adding Oauth authorization or security layers

## Practical Example

A web frontend developer needs to construct a single front page that requires data from 4 different calls to their backend services, e.g:


    1) api.careers360.com/api/1/header
    2) api.careers360.com/api/1/footer
    3) college-service.careers360.com/api/1/college-details/{college_id}
    4) listing-service.careers360.com/api/1/college-listing/{college_id}

The screen is very simple, and the client only needs to retrieve data from 4 different sources, wait for the round trip and then hand pick only a few fields from the response.

What if the client could call a single endpoint?

    1) gateway.careers360.com/api/1/college-detail-combined/{college_id}

That's something Vision can do this for us. And this is how it would look like:

<img src="https://production-cnext.s3.ap-south-1.amazonaws.com/gateway.drawio.png" width="500" alt="API Gateway Logo" align="center" />

Vision (API Gateway) would merge all the data and return only the fields we need

## Installation

```bash
$ go run main.go
```

## Running the app

```bash
# development
$ go run main.go

# production mode
$ export GOMAXPROCS=4
$ export GIN_MODE=release
$ go build main.go
```