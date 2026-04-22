# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Run

```bash
go build -o wol.exe
./wol.exe
```

## Architecture

```
main.go → RangeConfig() → for each config section → registerServer()
                                               → WolTopic{TopicName, Parameter}
                                               → topic.ConnectMqtt() → util.ConnectMqtt()
```

**Flow:** `main.go` loads config, iterates over `config.ini` sections (skipping `DEFAULT`). Each section with `struct=wol` creates a `WolTopic` that:
- Verifies required `mac` and `broadcast` keys
- Connects to Bemfa MQTT broker with that topic name
- Handles `on` → sends Magic Packet via `Server.Wol()`
- Handles `off` (Linux only) → runs `net rpc shutdown` via `util.RunCommand()`

**Topic interface** (`model/topic.go`): `ConnectMqtt()`, `MessageHandler(client, msg)`, `Verify() error` — `WolTopic` implements this.

**Config parsing** (`global/config.go`): Loads `config.ini` via `go-ini/ini`. DEFAULT section holds `bemfa_broker`, `bemfa_port`, `bemfa_client_id`. Per-topic sections hold `mac`, `broadcast`, `ip`, `user`, `password`.

**Logging** (`initialize/zapLog.go`): Dual-output zap logger — file `stdout.log` + console. SugarLogger exposed via `global.SugarLogger`.

## Key Files

- `config.ini` — runtime config (not committed; `config.ini.example` per README)
- `rangeConfig.go` — config section iteration and topic registration
- `util/mqtt.go` — MQTT client connect/subscribe loop
- `Server/wol.go` — Magic Packet construction and UDP send
