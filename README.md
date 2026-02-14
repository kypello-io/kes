# Kypello KES - Key Encryption Service

**This is the Kypello fork of MinIO KES, part of the Kypello ecosystem.**

## About This Fork

This repository is a community-maintained fork of [MinIO KES](https://github.com/minio/kes) under the AGPLv3 license. It is maintained by the [Kypello project](https://github.com/kypello-io) to provide key encryption services for Kypello Object Storage and other S3-compatible deployments.

**Maintenance Policy**: This fork receives bug fixes, security updates, and dependency updates only. New features are contributed upstream to MinIO KES when possible.

---

**KES is a cloud-native distributed key management and encryption server designed to secure modern applications at scale.**

 - [What is KES?](#what-is-kes)
 - [Installation](#install)
 - [Quick Start](#quick-start)
 - [Documentation](#docs)

[![fossabot enabled](https://img.shields.io/badge/fossabot-enabled-brightgreen.svg)](https://fossa.com/)

## Kypello Ecosystem

Kypello KES is part of the [Kypello project](https://github.com/kypello-io), a community-maintained fork of MinIO that preserves enterprise features like OIDC/SSO and the Admin UI under the AGPLv3 license. KES provides encryption key management for:

- **Kypello Object Storage** - S3-compatible object store (see [kypello-io/kypello](https://github.com/kypello-io/kypello))
- Server-side encryption (SSE-S3, SSE-KMS)
- Client-side encryption
- Other applications requiring secure key management

## What is KES?

KES (Key Encryption Service) is a distributed key management server that scales horizontally. It can either be run as edge server close to the applications
reducing latency to and load on a central key management system (KMS) or as central key management service. KES nodes are self-contained
stateless instances that can be scaled up and down automatically.

<p align="center">
  <img src='.github/arch.png?sanitize=true' width='70%'>
</p>

## Install

The KES server and CLI is available as a single binary, container image or can be build from source.

<details open="true"><summary><b><a name="homebrew">Homebrew</a></b></summary>

> **Note**: A kypello-io Homebrew tap is not yet available. Use binary releases or build from source instead.

For development, you can install the upstream MinIO KES which is API-compatible:
```sh
brew install minio/stable/kes
```

</details>

<details><summary><b><a name="docker">Docker</a></b></summary>

Pull the latest release via:
```
docker pull ghcr.io/kypello-io/kes:latest
```

</details>

<details><summary><b><a name="binary-releases">Binary Releases</a></b></summary>

| OS      | ARCH    | Binary                                                                                       |
|:-------:|:-------:|:--------------------------------------------------------------------------------------------:|
| linux   | amd64   | [linux-amd64](https://github.com/kypello-io/kes/releases/latest/download/kes-linux-amd64)         |
| linux   | arm64   | [linux-arm64](https://github.com/kypello-io/kes/releases/latest/download/kes-linux-arm64)         |
| darwin  | arm64   | [darwin-arm64](https://github.com/kypello-io/kes/releases/latest/download/kes-darwin-arm64)       |
| windows | amd64   | [windows-amd64](https://github.com/kypello-io/kes/releases/latest/download/kes-windows-amd64.exe) |

Download the binary via `curl` but replace `<OS>` and `<ARCH>` with your operating system and CPU architecture.
```
curl -sSL --tlsv1.2 'https://github.com/kypello-io/kes/releases/latest/download/kes-<OS>-<ARCH>' -o ./kes
```
```
chmod +x ./kes
```

You can also verify the binary with [minisign](https://jedisct1.github.io/minisign/) by downloading the corresponding [`.minisig`](https://github.com/kypello-io/kes/releases/latest) signature file.
Run:
```
curl -sSL --tlsv1.2 'https://github.com/kypello-io/kes/releases/latest/download/kes-<OS>-<ARCH>.minisig' -o ./kes.minisig
```
```
minisign -Vm ./kes -P RWTx5Zr1tiHQLwG9keckT0c45M3AGeHD6IvimQHpyRywVWGbP1aVSGav
```

> **Note**: If using minisign verification, the signing key may still be MinIO's key until Kypello establishes its own signing infrastructure.

</details>   
   
<details><summary><b><a name="build-from-source">Build from source</a></b></summary>

Download and install the binary via your Go toolchain:

```sh
go install github.com/kypello-io/kes/cmd/kes@latest
```

</details>
   
## Quick Start

Get started by setting up your own KES server in less than five minutes. This guide uses a local development configuration.

<details open><summary><b>First steps</b></summary>

#### 1. Start a Development KES Server

For testing and development, start a KES server with in-memory storage:

```sh
kes server --dev
```

This starts KES at `https://127.0.0.1:7373` with a self-signed certificate and prints the API key to the console.

#### 2. Configure CLI

In a new terminal, point the KES CLI to your local server:
```sh
export KES_SERVER=https://127.0.0.1:7373
export KES_API_KEY=<copy-from-server-output>
```

#### 3. Create a Key
Create a new root encryption key - e.g. `my-key`:
```
kes key create my-key
```
> Note: Creating a key will fail with `key already exists` if it already exists.

#### 4. Generate a DEK
Derive a new data encryption key (DEK):
```sh
kes key dek my-key
```
The plaintext part of the DEK is used by applications to encrypt data.
The ciphertext part is stored alongside the encrypted data for future decryption.

> **Production Setup**: For production deployments, configure KES with a proper KMS backend (Vault, AWS KMS, etc.) instead of in-memory storage. See the [integration guides](https://github.com/minio/kes/wiki#supported-kms-targets) for details.

</details>   

## Docs

Kypello KES maintains API compatibility with upstream MinIO KES. Most documentation applies directly to this fork.

### Documentation Resources

- [MinIO KES Documentation](https://min.io/docs/kes/) - Comprehensive KES documentation
- [Integration Guides](https://github.com/minio/kes/wiki#supported-kms-targets) - Supported KMS backends
- [Command Line](https://min.io/docs/kes/cli/#available-commands) - CLI reference
- [Server API](https://min.io/docs/kes/concepts/server-api/) - HTTP API documentation
- [Go SDK](https://pkg.go.dev/github.com/minio/kes-go) - Compatible Go client library

> **Note**: This fork maintains compatibility with upstream MinIO KES. The upstream documentation applies directly. For Kypello-specific configurations or integration with Kypello Object Storage, see the examples in this repository.

### Monitoring

KES servers provide an API endpoint `/v1/metrics` that observability tools like [Prometheus](https://prometheus.io/) can scrape.
Refer to the [monitoring documentation](https://min.io/docs/kes/concepts/monitoring/) for setup instructions.

For a graphical Grafana dashboard, refer to the [example](examples/grafana/dashboard.json).

![](.github/grafana-dashboard.png)  

## FAQs

<details><summary><b>I have received an <code>insufficient permissions</code> error</b></summary>
   
This means that you are using a KES identity that is not allowed to perform a specific operation, like creating or listing keys.

The KES [admin identity](https://github.com/kypello-io/kes/blob/master/server-config.yaml#L8)
can perform any general purpose API operation. You should never experience a `not authorized: insufficient permissions`
error when performing general purpose API operations using the admin identity.

In addition to the admin identity, KES supports a [policy-based](https://github.com/kypello-io/kes/blob/master/server-config.yaml#L77) access control model.
You will receive a `not authorized: insufficient permissions` error in the following two cases:
1. **You are using a KES identity that is not assigned to any policy. KES rejects requests issued by unknown identities.**

   This can be fixed by assigning a policy to the identity. Checkout the [examples](https://github.com/kypello-io/kes/blob/master/server-config.yaml#L79-L88).
2. **You are using a KES identity that is assigned to a policy but the policy either not allows or even denies the API call.**
   
   In this case, you have to grant the API permission in the policy assigned to the identity. Checkout the [list of APIs](https://github.com/minio/kes/wiki/Server-API#api-overview).
   For example, when you want to create a key you should allow the `/v1/key/create/<key-name>`. The `<key-name>` can either be a
   specific key name, like `my-key-1` or a pattern allowing arbitrary key names, like `my-key*`.
   
   Also note that deny rules take precedence over allow rules. Hence, you have to make sure that any deny pattern does not
   accidentally matches your API request.

</details>   
   
***

## License

Use of `KES` is governed by the AGPLv3 license that can be found in the [LICENSE](./LICENSE) file.

This project is a fork of [MinIO KES](https://github.com/minio/kes), originally developed by MinIO, Inc. and licensed under AGPLv3. The Kypello fork maintains the same AGPLv3 license terms with no commercial exception.
