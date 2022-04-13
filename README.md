# Go | Fiber | Nodemon

## Start 🚀

```bash
docker-compose up
```

## Stop ⛔

```bash
docker-compose down
```

## Debug 🐛

The application will start the debugger in development mode. Attach to port `2345` using ["Go Remote"](https://www.jetbrains.com/help/go/go-remote.html) in [GoLand](https://www.jetbrains.com/go/).

**Note:** On disconnect: Leave it running

#### Alternative

Another way to start with the debugger is to let Delve build and run.

```json
// nodemon.json
{
  "exec": "dlv debug --headless --continue --listen=:2345 --api-version=2 --accept-multiclient"  
}
```
