dogo [![Build Status](https://travis-ci.org/domluna/dogo.svg?branch=master)](https://travis-ci.org/domluna/dogo)[![Coverage Status](https://coveralls.io/repos/domluna/dogo/badge.png?branch=master)](https://coveralls.io/r/domluna/dogo?branch=master)[![GoDoc](https://godoc.org/github.com/domluna/dogo?status.png)](https://godoc.org/github.com/domluna/dogo)
====

## Overview

DigitalOcean Go Client V2 API.

Alternatives:
  * [DO Go Official API](https://github.com/digitalocean/godo)
  * [Pearkes DO API](https://github.com/pearkes/digitalocean)

I think this API is a sweetspot between these two. It has more features than Pearkes but more lightweight than the DO Official API.

## Getting Started

### Installing

```sh
$ go get github.com/domluna/dogo
```

### Example

```go
package main

import (
	"github.com/domluna/dogo"
)

func main() {

  // If the token is the empty string("") then it'll attempt
  // to fill the value under env var $DIGITALOCEAN_TOKEN
	client := dogo.NewClient("") // $DIGITALOCEAN_TOKEN

	// get all droplets
	droplets, err := client.ListDroplets()
	if err != nil {
		// deal with error
	}
	// do stuff with droplets
}
```
