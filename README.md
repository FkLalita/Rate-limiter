# Rate Limiter Middleware for Go

A reusable rate limiter middleware package for Go applications, implementing a token bucket algorithm.

## Overview

This Go package provides a middleware for enforcing rate limits in your HTTP applications. It uses a token bucket algorithm to control the number of requests a user can make in a given time period.

## Features

- Simple integration with existing Go projects.
- Adjustable rate-limiting parameters: capacity and refill amount.
- Token bucket algorithm for efficient rate limiting.

## Getting Started

### Installation

```bash
go get github.com/FkLalita/Rate-limiter


## Usage

```bash
import "github.com/FkLalita/Rate-limiter"
rl := ratelimiter.NewRateLimiter(10, 1) // capacity: 10, refill amount: 1
router.Handle("/protected/resource", rl.Middleware(http.HandlerFunc(protectedResourceHandler)))






