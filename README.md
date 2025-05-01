# pddns.go: Updates DNS entries on Cloudflare

## Usage:
```shell
pddns --config config.json --dns-entry subdomain.domain.com --ip 123.456.789.012
```
If no IP is specified, the new IP is set to current external IP.

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
     "api_token": "your-api-token"
   }
   ```

## Prebuilt binaries ready to go:

After running the `make` command, you will find the following compiled binaries in the `build` directory:

| Platform              | Binary Path                         | Notes                                      |
|-----------------------|-------------------------------------|--------------------------------------------|
| **macOS (Intel)**      | [build/pddns-macos-amd64](./build/pddns-macos-amd64) | Make executable with `chmod +x`            |
| **macOS (M chips)**      | [build/pddns-macos-arm64](./build/pddns-macos-arm64) | Make executable with `chmod +x`            |
| **Linux**      | [build/pddns-linux-amd64](./build/pddns-linux-amd64) | Make executable with `chmod +x`            |
| **Raspberry Pi (3B/4/5)**      | [build/pddns-linux-arm64](./build/pddns-linux-arm64) | Make executable with `chmod +x`            |
| **FreeBSD**    | [build/pddns-freebsd-amd64](./build/pddns-freebsd-amd64) | Make executable with `chmod +x`            |
| **Windows**    | [build/pddns-windows-amd64.exe](./build/pddns-windows-amd64.exe) | Run as `.exe` file directly                |
