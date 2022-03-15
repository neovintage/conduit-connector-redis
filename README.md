# Conduit Redis Connector

This is a proof of concept in building a connector for conduit. This a source connector that will subscribe to a single channel in Redis. Once the payload is in Conduit, you can send the data anywhere you need to in your infrastructure.

## Getting Started

### Prerequsites

- Conduit is install on your machine. See the Conduit instructions for getting started.
- Redis is installed. See Redis instructions. If you're on a mac, it's as easy as doing `brew install redis`.

### Steps

Let's assume we have Conduit installed your home path (e.g. `~/`)

1. Clone this repo
2. Run `make`. This will produce a binary called `conduit-redis-connector`.
3. Start Conduit (example: `./conduit`)
    ```
    $ ~/conduit
    ```

4. Start Redis
    ```
    $ redis-server
    ```

4. Create a pipeline in Conduit. Let's use the API for this:
   ```

   ```

5. For this part, we'll need to set up a pipeline via the API in Conduit. Here's an example payload using curl to post to the API:
   ```
   $ curl -

   ```

6. Let's say we want to save all of the messages that come off the channel into a file. Let's set that up

