# API Monitor

A simple CLI tool to check whether APIs are online and responding correctly.

## Motivation

I work with several public APIs across different projects. 
Instead of opening each URL manually, I built this tool to check all of them at once.

## How it works

Reads a list of APIs from a local `apis.json` file, makes a GET request to each endpoint and prints the result in terminal.

## Setup

Clone the repository:

```bash
git clone https://github.com/utf8-porfavor/api-monitor.git
cd api-monitor
```

Create your local configuration file:

```bash
cp apis.example.json apis.json
```

Edit `apis.json` with your APIs.

## Usage

Just run the project:

```bash
go run main.go
```

## Example Output

```bash
[OK] Lotofácil -> status 200
[FALHA] Minha API -> status 500
[ERRO] API Local -> connection timeout
```