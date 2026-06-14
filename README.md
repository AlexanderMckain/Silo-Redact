# Silo-Redact

Scrub sensitive data before it hits the LLM.

SiloRedact is a **local-first Windows desktop app** that redacts secrets, credentials, and PII from text before you paste into AI chat, tickets, or support portals. All redaction runs on your machine — nothing is sent to a cloud scrubbing service.

## Why SiloRedact

Paste is permanent. SiloRedact intercepts sensitive data at the clipboard so you can share logs and diagnostics without oversharing. No account, license key, or API keys required.

## Features

- **22 forensic presets** — AWS keys, Stripe secrets, JWTs, IPs, and common credential patterns
- **Ghost Mode** — background clipboard monitoring with real-time redaction
- **Custom regex rules** — build, import, and manage your own patterns
- **Satellite Mode** — compact always-on-top UI for AI workflows
- **Auto-Copy & timed clipboard wipe** — optional 30-second clipboard clear
- **Debug Mode** — opt-in view of which rules matched (use with care around live secrets)

## Get SiloRedact

- **Microsoft Store:** [Install SiloRedact](https://apps.microsoft.com/detail/9mzbmn46q9sh)
- **Website:** [siloredact.com](https://siloredact.com)

## Build from source

**Requirements:** Go 1.24+, [Wails CLI v2](https://wails.io/docs/gettingstarted/installation), Node.js 18+

```bash
cd frontend && npm install && cd ..
wails dev
```

Production build:

```bash
cd frontend && npm run build && cd ..
wails build
```

Windows MSIX packaging assets are under `build/windows/`.

## License

MIT — see [LICENSE](LICENSE).
