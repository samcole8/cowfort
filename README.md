# mootd

![](https://img.shields.io/badge/status-maintained-green) [![](https://img.shields.io/github/v/release/samcole8/mootd)](https://github.com/samcole8/mootd/releases/latest)

mootd is a tiny MOTD server that generates randomised daily messages using Cowsay and Fortune. All it requires is a Linux host with Docker installed.

```
 ________________________________________ 
/ All the existing 2.0.x kernels are too \
| buggy for 2.1.x to be the main goal.   |
|                                        |
\ -- Alan Cox                            /
 ---------------------------------------- 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

## Installation

If you haven't already, [install Docker](https://docs.docker.com/engine/install/). Then, use `docker compose` to build and run the container:

```bash
docker compose up -d
```

## Configuration

Define environment variables in the compose file (or run command) to override the defaults.

| Env | Default | Example | Description |
|---|---|---|---|
| `TZ` | `UTC` | `London/Europe` | Timezone used by the server. |
| `RENEWAL_TIME` | `24:00:00` | `12:30:00` | MOTD renewal time (HH:MM:SS). |
| `CHANCE` | `10` | `100` | Reciprocal chance for atypical cow. The example used gives a 1 in 100 chance of a random non-standard cow; the default is 1 in 10. |