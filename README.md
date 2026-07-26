# mermaid-demo-notify

Notification **client SDK** and delivery service. Consumes webhooks from
[`mermaid-demo-payments`](https://github.com/dar-io/mermaid-demo-payments) and delivers email and push.

> Synthetic fixture for the Mermaid AI GitHub connector. Not a real notification system.

## Layout

| Path | What lives there |
|---|---|
| `internal/ingest/` | Webhook receipt from mermaid-demo-payments |
| `internal/api/` | HTTP handlers for the client SDK |
| `internal/render/` | Template rendering |
| `internal/deliver/` | Email and push transports |
| `internal/queue/` | Consumer draining the delivery queue |
| `db/` | Data models |
