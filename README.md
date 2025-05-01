# pddns.go: Updates DNS entries on Cloudflare

## Usage:
```shell
pddns --config config.json --dns-entry subdomain.domain.com
```

## Cloudflare Setup

Before using this tool, you need to get the following information from your Cloudflare account:

### 1. **Find your Cloudflare Zone ID**
   - Go to the [Cloudflare Dashboard](https://dash.cloudflare.com).
   - Select your domain.
   - Click on the **"Overview"** tab.
   - Scroll down to the **"API"** section.
   - Your **Zone ID** is listed under **Zone ID**.

   You can also follow the official guide on how to find your Zone ID [here](https://developers.cloudflare.com/fundamentals/setup/find-account-and-zone-ids/).

### 2. **Create an API Token**
   - Visit the [Create API Token page](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/).
   - Click on **"Create Token"**.
   - Use a pre-built template or create a custom token.
     - For custom tokens, select the **"DNS Edit"** permissions to allow this tool to update DNS records.
   - After creating the token, **copy** it, as you will need to input it in your configuration file.

### 3. **Configuring the Tool**

   - Create a json config file (named `config.json` here).
   - The `config.json` file should contain the following fields:

   ```json
   {
     "zone_id": "your-zone-id",
     "email": "your-cloudflare-email@example.com",
     "api_token": "your-api-token"
   }
   ```

## Prebuilt binaries ready to go:

| OS         | Arch   | Download |
|------------|--------|----------|
| FreeBSD    | amd64  | [pddns-freebsd-amd64](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-freebsd-amd64) |
| macOS (Intel)      | amd64  | [pddns-darwin-amd64](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-darwin-amd64) |
| macOS (M chips)     | arm64  | [pddns-darwin-arm64](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-darwin-arm64) |
| Linux      | amd64  | [pddns-linux-amd64](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-linux-amd64) |
| Raspberry Pi (3B/4/5)      | arm64  | [pddns-linux-arm64](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-linux-arm64) |
| Windows    | amd64  | [pddns-windows-amd64.exe](https://github.com/YOUR_USERNAME/pddns/releases/download/vX.Y.Z/pddns-windows-amd64.exe) |
