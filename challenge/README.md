# Meli Reader File

## Description

version  **_v0.0.1:_** 

## Installation

`go get github.com/ingleandrotorres/file-reader-challenge`

By default, `go get` will bring in the latest tagged release version of the api.

## Getting started

This api takes a local file in order to build a id with it data, next It search this ids  against a meli api and build a new entity to save in a repository 

### Library configuration

1. this api does not support full authentication  **yet** and its is only build for testing purpose
2. It is clearly necessary to implement security issues such as credentials to require the mercadolibre API token
3. Some concrete implementation only work for test purposes. It will be replaced for interfaces
4. Some variables may be hard coded. It will be replaced for primitives. 
5. this can get better and better, only we have to get more time and work
