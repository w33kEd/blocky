# Go-Blocky

A **minimal, educational peer‑to‑peer blockchain implementation written in Go**.
This project demonstrates how a basic blockchain network works end‑to‑end, including:

* Block and transaction data structures
* Proof‑of‑Work (PoW) mining
* Transaction signing and verification (ECDSA)
* Peer discovery and neighbour synchronization
* Distributed consensus via longest‑chain rule
* Simple HTTP‑based P2P communication

The codebase is intentionally explicit and low‑level to make blockchain mechanics easy to inspect, debug, and extend.

> **This project is for learning and experimentation. It is not production‑grade and should not be used for real assets.**

---

## Architecture Overview

Each node runs:

* An HTTP server (one port per node)
* A local blockchain instance
* A transaction pool (mempool)
* A miner (Proof‑of‑Work)
* A neighbour discovery loop

Nodes:

* Discover peers by scanning a configurable local port range
* Broadcast transactions to neighbours
* Periodically resolve conflicts using the longest valid chain
* Mine blocks independently

All nodes are **fully equal peers** — there is no master node.

---

## Features

* Deterministic block hashing (SHA‑256)
* ECDSA‑signed transactions
* Coinbase (mining reward) transactions
* Automatic peer discovery
* Automatic peer synchronization
* Chain conflict resolution
* JSON‑based REST API

---

## Requirements

* Go **1.20+** recommended
* Linux / macOS / Windows

Check Go installation:

```bash
go version
```

---

## Installation

Clone the repository:

```bash
git clone https://github.com/w33ked/go-blocky.git
cd go-blocky
```

Install dependencies:

```bash
go mod tidy
```

---

## Running a Single Node

Start one blockchain node on port `5000`:

```bash
go run . --port=5000
```

You should see output similar to:

```text
private_key ...
public_key ...
blockchain_address ...
```

The node will:

* Generate a miner wallet
* Create a genesis block
* Start neighbour discovery

---

## Running Multiple Nodes (Local Network)

### Example: 3‑Node Network

Open **three terminals**.

#### Terminal 1

```bash
go run . --port=5000
```

#### Terminal 2

```bash
go run . --port=5001
```

#### Terminal 3

```bash
go run . --port=5002
```

Each node will:

* Scan ports `5000–5003`
* Automatically discover other running nodes
* Synchronize chains

Neighbour lists will stabilize automatically.

---

## Port Configuration

Defined in `block/blockchain.go`:

```go
BLOCKCHAIN_PORT_RANGE_START = 5000
BLOCKCHAIN_PORT_RANGE_END   = 5003
```

Only ports with running nodes will be added as neighbours.

If a port is unused, connection attempts may log:

```text
connect: connection refused
```

This is expected behavior.

---

## Mining

Mining can be triggered manually or automatically.

### Manual Mining

```bash
curl http://localhost:5000/mine
```

### Automatic Mining Loop

```bash
curl http://localhost:5000/mine/start
```

Mining:

* Locks the blockchain
* Adds a coinbase transaction
* Performs Proof‑of‑Work
* Broadcasts consensus requests to neighbours

---

## Transactions

### Submit a Transaction

```bash
curl -X POST http://localhost:5000/transactions \
  -H "Content-Type: application/json" \
  -d '{
    "sender_blockchain_address": "...",
    "recipient_blockchain_address": "...",
    "sender_public_key": "...",
    "value": 1.5,
    "signature": "..."
  }'
```

Transactions are:

* Verified using ECDSA signatures
* Broadcast to all neighbours
* Added to the mempool

---

## REST API

### Get Blockchain

```
GET /chain
```

### Get Transaction Pool

```
GET /transactions
```

### Submit Transaction

```
POST /transactions
```

### Broadcast Transaction (Peer‑to‑Peer)

```
PUT /transactions
```

### Clear Transaction Pool

```
DELETE /transactions
```

### Mine a Block

```
GET /mine
```

### Start Mining Loop

```
GET /mine/start
```

### Resolve Conflicts

```
PUT /consensus
```

### Get Wallet Balance

```
GET /amount?blockchain_address=ADDRESS
```

---

## Consensus Model

* Longest valid chain wins
* Chains are validated by:

  * Hash linkage
  * Proof‑of‑Work difficulty
* Nodes replace their chain automatically when a longer valid chain is found

---

## Security Notes

* No TLS
* No authentication
* No DoS protection
* No persistence (in‑memory only)

This project is **strictly educational**.

---

## Extending the Project

Common next steps:

* Persistent storage (LevelDB / BoltDB)
* Proper P2P handshake
* Gossip‑based transaction propagation
* Merkle trees
* Adjustable difficulty
* Network bootstrap nodes
* CLI wallet tooling

---

## License

MIT License

---

## Author

Built for learning, experimentation, and protocol exploration.

If you are studying blockchains, distributed systems, or Go — this project is designed for you.
