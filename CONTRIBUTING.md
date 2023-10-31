# CONTRIBUTING

## Dev

For run the dev environement, it is needed to run the command:

```
docker compose -f docker-compose.dev.yaml up -d
```

In case of change in one service, it can be rebuilt and redeployed. For example for dashboard service:

```
docker compose -f docker-compose.dev.yaml up -d --build dashboard
```

## Logging

For logging, [Logrus](https://github.com/sirupsen/logrus) is being used. There is an init file for setting up all the configuration, and one middleware for logging the HTTP requests. Then it is needed to

```
log "github.com/sirupsen/logrus"

log.Trace("Something very low level.")
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
// Calls os.Exit(1) after logging
log.Fatal("Bye.")
// Calls panic() after logging
log.Panic("I'm bailing.")

```
