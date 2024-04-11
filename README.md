
# simplesurance golang coding challenge

Using only the standard library, create a Go HTTP server that on each request responds with a
counter of the total number of requests that it has received during the previous 60 seconds
(moving window). The server should continue to return the correct numbers after restarting it, by
persisting data to a file.

# implementation notes

I tried to KISS while demonstrating my implementation of the moving window alg. This code balances a few assumptions to keep it simple while providing the functionality (as for example, we save the file on each tick)

## API description

#### Get amount of requests in the last configurable amount of seconds

``
  GET /counter
``
## Env vars
   can be edited on docker-compose.yaml

   API_PORT: 8080

   INTERVAL_IN_SECONDS: 60

   NUMBER_OF_TICKS: 60   //*this specifies the number of ticks on a moving window, essentially we have a tick will last INTERVAL_IN_SECONDS / NUMBER_OF_TICKS*

   FILE_LOCATION: ../sliding_window_counter.gob

# How to run

```bash
  docker-compose -up -d
```
    
