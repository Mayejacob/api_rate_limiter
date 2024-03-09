# GoLang Rate Limiting Library

This GoLang library provides rate limiting functionalities with three different options: token bucket, tollbooth, and per client rate limiting.


## Token Bucket: 
A classic rate limiting algorithm that simulates a leaky bucket with a fixed capacity. Requests are allowed as long as there are enough tokens available.
## Tollbooth: 
A popular Go library that offers various rate limiting implementations, including token bucket and fixed window counters. (https://github.com/didip/tollbooth)
## Per-Client Rate Limiting:
 A custom implementation that tracks the rate limit for each individual client.
## Installation

You can install the library using `git clone`:

```sh
git clone https://github.com/Mayejacob/api_rate_limiter.git
```