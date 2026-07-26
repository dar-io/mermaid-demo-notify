# Integrating with mermaid-demo-payments

mermaid-demo-payments POSTs `{ type, payload }` to this service's `/ingest` endpoint.

| Upstream event | This service does |
|---|---|
| `charge.succeeded` | Renders a receipt and queues email |
| `charge.failed` | Queues a failure notice on email and push |

The envelope is the whole contract. There is no shared schema package, which is
why a field rename upstream is only caught here at runtime.
