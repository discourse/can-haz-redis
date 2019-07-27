# Can Haz Redis?

This is a simple application-level probe for Redis. It is intended to translate
HTTP health checks from a load balancer to PING checks on the Redis server.

Export the [connection string](https://www.iana.org/assignments/uri-schemes/prov/redis)
to the server in the `REDIS_URL` environment variable and point the health check at
`localhost:6378/ping`.
