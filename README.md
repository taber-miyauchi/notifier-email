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

## Testing Precise Code Navigation

Open this repo in Sourcegraph and try the following:

### 1. Go to Definition (cross-repo)

- In `email.go`, click on `core.Message` (line 28) → **Go to Definition**
- → Jumps to `Message` struct in `notifier-core/message.go`

### 2. Find References (cross-repo)

- In `email.go`, click on `EmailNotifier` (line 11) → **Find References**
- → Shows usage in `notifier-service/main.go`
