# Go Learning Project

The goal of this project is to learn Go in context of back-end development and Web3 interactions.

I planning to learn Go to explore Web3 ecosystem beyond TypeScript and Rust and to have simple back-end language in stack

## Project Description

It's simple: self-written indexer

Indexers are complex solutions that require a lot of effort to make it fast and efficient

In my understanding indexer includes following:

-   Polling new blocks
-   Parsing transactions included in block
-   (Optional) Handling re-orgs
-   Storing indexed data in DB like PostgresQL
-   Provide access to scraped data via REST or GraphQL API

In case of completion of such project, I will be sure, that my skills with specified language are sufficient to make other production grade solutions

## Project Timelines

I'm doing this project at my free-time, so I'm limited in terms of hours spent weekly to just around 10-15 hours weekly

It would be perfect to complete the project in max total of 2 months, which seems achievable for me

Start: 27.09.2025

## Go Learning Checkpoints

Basic Go

-   Familiar with reading/writing from/to files
-   Able to handle WebSocket connections
-   Make HTTP requests and handle errors
-   Prometheus metrics integration
-   Proper logging, parsable and analyzable in Loki/Grafana
-   Unit & Integration testing

go-ethereum

-   Making subscriptions for new txs/blocks/headers
-   Calling arbitrary RPC methods
-   Calling typed contract functions via use of ABIs
-   Building and sending transactions (not actually part of the project, but part of basic skill set)
-   Retrieving data via Multicall3

Back-End

-   Swagger
-   REST API and (or) GraphQL API
-   OpenAPI specification generation
-   PostgresQL
-   Middleware (auth/logging/metrics)

## Project Architecture (by Claude)

Architecture advised

```
block-parser/
├── cmd/
│   ├── indexer/
│   │   └── main.go                 # Main application entrypoint
│   └── api/
│       └── main.go                 # API server entrypoint
├── internal/
│   ├── config/
│   │   ├── config.go               # Configuration struct and loading
│   │   └── env.go                  # Environment variable handling
│   ├── indexer/
│   │   ├── indexer.go              # Main indexer logic
│   │   ├── block_processor.go      # Block processing logic
│   │   ├── tx_parser.go            # Transaction parsing
│   │   └── reorg_handler.go        # Re-org handling
│   ├── rpc/
│   │   ├── client.go               # RPC client abstraction
│   │   ├── alchemy.go              # Alchemy provider
│   │   ├── infura.go               # Infura provider
│   │   └── multicall.go            # Multicall3 integration
│   ├── storage/
│   │   ├── postgres/
│   │   │   ├── postgres.go         # PostgreSQL connection
│   │   │   ├── migrations/         # DB migrations
│   │   │   └── queries.sql         # SQL queries
│   │   └── repository/
│   │       ├── block_repo.go       # Block data repository
│   │       └── tx_repo.go          # Transaction repository
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── blocks.go           # Block endpoints
│   │   │   └── transactions.go     # Transaction endpoints
│   │   ├── middleware/
│   │   │   ├── auth.go             # Authentication
│   │   │   ├── logging.go          # Request logging
│   │   │   └── metrics.go          # Prometheus metrics
│   │   └── router.go               # Route definitions
│   ├── models/
│   │   ├── block.go                # Block data structures
│   │   ├── transaction.go          # Transaction structures
│   │   └── contract.go             # Contract structures
│   └── metrics/
│       ├── metrics.go              # Prometheus metrics setup
│       └── collectors.go           # Custom metric collectors
├── pkg/
│   ├── logger/
│   │   └── logger.go               # Structured logging setup
│   ├── errors/
│   │   └── errors.go               # Custom error types
│   └── utils/
│       ├── hex.go                  # Hex conversion utilities
│       └── retry.go                # Retry logic utilities
├── test/
│   ├── integration/
│   │   ├── indexer_test.go         # Integration tests
│   │   └── api_test.go             # API integration tests
│   └── fixtures/
│       ├── blocks.json             # Test data
│       └── transactions.json      # Test data
├── scripts/
│   ├── migrate.go                  # Database migration script
│   └── setup.sh                   # Project setup script
├── deployments/
│   ├── docker/
│   │   ├── Dockerfile              # Container definition
│   │   └── docker-compose.yml     # Local development setup
│   └── k8s/                       # Kubernetes manifests
├── docs/
│   ├── api.md                      # API documentation
│   └── architecture.md            # Architecture documentation
├── .env.example                    # Environment variables template
├── .gitignore
├── README.md
├── go.mod
├── go.sum
└── Makefile                        # Build and development commands
```

## Code Size Projections (by Claude)

Based on your indexer project specification and the architecture we outlined, here's my estimation:

### Core Application (~2,500-3,500 lines)

-   **Config & Environment**: ~200-300 lines
-   **RPC Client Layer**: ~400-600 lines
-   **Indexer Core Logic**: ~800-1,200 lines
-   **Database Layer**: ~600-900 lines
-   **API Handlers**: ~300-500 lines
-   **Models/Structs**: ~200-300 lines

### Supporting Code (~1,500-2,000 lines)

-   **Error Handling**: ~150-200 lines
-   **Logging & Metrics**: ~300-400 lines
-   **Utilities & Helpers**: ~200-300 lines
-   **Middleware**: ~200-300 lines
-   **Database Migrations**: ~100-200 lines
-   **Main Applications** (`cmd/`): ~100-150 lines

### Testing (~1,500-2,500 lines)

-   **Unit Tests**: ~1,000-1,500 lines
-   **Integration Tests**: ~500-1,000 lines

### Configuration & Scripts (~500-700 lines)

-   **Docker/K8s configs**: ~200-300 lines
-   **SQL queries**: ~150-250 lines
-   **Scripts & Makefiles**: ~150-200 lines

## **Total Estimation: 6,000-8,700 lines**

### Conservative estimate: **~7,000 lines**

## Factors that could increase size:

-   Complex transaction parsing logic
-   Sophisticated re-org handling
-   Multiple contract ABIs
-   Advanced error recovery
-   Comprehensive monitoring
-   GraphQL API (vs simple REST)

## Factors that could decrease size:

-   Using ORM instead of raw SQL
-   Leveraging more third-party libraries
-   Simpler data models
-   Basic error handling

For a **production-grade indexer** with proper testing, monitoring, and error handling, 7,000 lines is realistic. This excludes generated code (protobuf, OpenAPI, etc.) and vendor dependencies.

Your 2-month timeline with 10-15 hours/week (80-120 total hours) seems achievable for this scope.
