# pddns.go: Updates DNS entries on Cloudflare

Usage:
```shell
pddns --config config.json --dns-entry subdomain.domain.com
```

`config.json` example:
```config.json
{
  "zone_id": "abcd1234567890",
  "api_token": "abcd1234567890",
  "email": "name@domain.com"
}
```
