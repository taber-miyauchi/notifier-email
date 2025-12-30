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

### 1. Go to Definition (cross-repo type)

Jump from a type usage to its definition in another repository.

- In `email.go`, click on `Message` (line 28) → **Go to Definition**
- → Highlights `Message` struct (line 13) in `notifier-core/message.go`

**Benefit:** Navigate directly from your implementation to the shared types you depend on—understand the contract without leaving your editor or manually searching another repo.

### 2. Find References (cross-repo function)

Locate all usages of an exported function across repository boundaries.

- In `email.go`, click on `NewEmailNotifier` function (line 18) → **Find References**
- → Highlights `NewEmailNotifier` (line 13) in `notifier-service/main.go`

**Benefit:** See exactly which services consume your implementation—essential for gauging adoption and planning breaking changes to your API.
