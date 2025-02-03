# Wallet CLI

**Wallet CLI** is a simple command-line tool for managing and tracking your colored papers(money) records.

## Installation

Clone the repository and build:

```sh
git clone <repository-url>
cd wallet-cli
make build
```

## Usage

### Help

```sh
wallet command -h
```

### Show Records

```sh
wallet show [--category <name>] [--limit <number>] [--month <number>]
```

### Add a Record

```sh
wallet add --amount <value> --category <name>
```

### Modify a Record

```sh
wallet modify --id <record_id> [--amount <value>] [--category <name>]
```

## Commands

- `wallet add` → Add a record
- `wallet show` → View records
- `wallet modify` → Edit a record
- `wallet help` → Command help
