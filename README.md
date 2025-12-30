# notifier-email

## Project Structure

```
├── notifier-core/          ← The foundation
│   ├── go.mod
│   ├── notifier.go         ← Notifier interface
│   ├── message.go          ← Message struct, Priority enum
│   └── README.md
│
├── notifier-email/         ← Implementation (you are here)
│   ├── go.mod              ← depends on notifier-core
│   ├── email.go            ← EmailNotifier implements Notifier
│   └── README.md
│
└── notifier-service/       ← Consumer/API
    ├── go.mod              ← depends on both
    ├── main.go             ← HTTP server using both packages
    └── README.md
```

## Overview

Email implementation of the `Notifier` interface from `notifier-core`.

## Types

- **`EmailNotifier`** - Sends notifications via SMTP

## Dependencies

- `github.com/sourcegraph-ce/notifier-core` - For `Notifier` interface and `Message` type

## Precise Code Navigation Demo

This repo **implements** the `Notifier` interface. Try these:

1. **"Find Implementations"** on `Notifier` in `notifier-core` → lands on `EmailNotifier.Send()`
2. **"Go to Definition"** on `core.Message` → jumps to `notifier-core/message.go`
