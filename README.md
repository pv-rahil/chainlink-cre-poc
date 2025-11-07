# Asset NAV Workflow

This Chainlink CRE (Compute Runtime Environment) workflow fetches real-time Ethereum price data from Price API and submits it on-chain to a smart contract. The workflow demonstrates how to integrate off-chain data sources with on-chain storage using Chainlink's decentralized oracle network.

## Overview

The workflow performs the following operations:
1. **Fetches off-chain data**: Retrieves current Ethereum USD price from Price API
2. **Reads on-chain data**: Queries the current NAV (Net Asset Value) stored in the smart contract
3. **Processes and scales data**: Converts the price to a scaled integer format (multiplied by 100)
4. **Generates CRE report**: Creates a cryptographically signed report with the new price data
5. **Submits on-chain**: Writes the report to the blockchain via the AssetNavContract

## Project Structure

```
chainlink-cre-poc/
├── asset-nav/                          # Workflow directory
│   ├── main.go                         # Main workflow logic
│   ├── config.json                     # Workflow configuration
│   └── README.md                       # This file
├── contracts/
│   └── evm/
│       └── src/
│           ├── AssetNavContract.sol    # Smart contract for storing NAV data
│           ├── abi/                    # Contract ABI files
│           └── generated/              # Go bindings for contract interaction
├── go.mod                              # Go module dependencies
├── go.sum                              # Go module checksums
├── project.yaml                        # CRE project configuration
└── .env                                # Environment variables (private keys)
```

## Smart Contract

The `AssetNavContract` is deployed on Ethereum Sepolia testnet at:
- **Address**: `0xd68558CEB4F472143DcEC7f9e8b0279fF0f7784a`
- **Chain Selector**: `16015286601757825753` (Ethereum Testnet Sepolia)

### Contract Features
- Stores current asset NAV as an `int256` (scaled by 100)
- Maintains historical NAV records with timestamps
- Implements `IReceiverTemplate` for CRE report validation
- Emits `AssetNavUpdated` events on each update

### Data Format
- **Off-chain**: Ethereum price in USD (e.g., $3496.36)
- **On-chain**: Scaled integer (e.g., 349636 represents $3496.36)
- **Conversion**: `onchain_value / 100 = USD_price`

## Configuration

### config.json
```json
{
  "schedule": "0 */1 * * * *",           // Cron schedule (every minute)
  "apiUrl": "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd",
  "evms": [
    {
      "consumerAddress": "0xd68558CEB4F472143DcEC7f9e8b0279fF0f7784a",
      "chainSelector": 16015286601757825753,
      "gasLimit": 1000000
    }
  ]
}
```

## Prerequisites

1. **CRE SDK (v0.6.0-alpha)**: Install the Chainlink CRE CLI
2. **Go 1.25.3+**: Required for building the workflow
3. **Funded Wallet**: Private key with Sepolia ETH for gas fees

## Setup Instructions

### 1. Clone the Repository
```bash
git clone <repository-url>
cd chainlink-cre-poc
```

### 2. Configure Environment Variables
Copy the example environment file and add your private key:
```bash
cp .env.example .env
```

Edit `.env` and add your private key:
```
CRE_ETH_PRIVATE_KEY=<your_private_key_here>
```

**Note**: The private key must correspond to a wallet funded with Sepolia ETH for transaction fees.

### 3. Install Dependencies
```bash
go mod download
```

## Running the Workflow

### Local Simulation
Navigate to the workflow directory and run the simulation:
```bash
cd asset-nav
./cre_v0.6.0-alpha workflow simulate --target local-simulation --config config.json main.go
```

### Expected Output
```
Workflow compiled
{"level":"info","msg":"Simulator Initialized"}
[USER LOG] level=INFO msg="Successfully fetched offchain value" json="{\"ethereum\":{\"usd\":3496.36}}"
[USER LOG] level=INFO msg="Successfully read onchain value" raw=349602 json="{\"ethereum\":{\"usd\":3496.02}}"
[USER LOG] level=INFO msg="Final calculated result to submit onchain" raw=349636 json="{\"ethereum\":{\"usd\":\"3496.36\"}}"
[USER LOG] level=INFO msg="CRE report generated successfully"
[USER LOG] level=INFO msg="Submitting CRE report onchain..."
[USER LOG] level=INFO msg="Report submitted successfully" txHash=0x...
[USER LOG] level=INFO msg="Workflow finished successfully"

Workflow Simulation Result:
{
  "OffchainValue": "3496.36",
  "PreviousOnchainValue": "3496.02",
  "UpdatedValue": "3496.36",
  "TxHash": "0x..."
}
```

## Workflow Logic

### Key Functions

1. **`onCronTrigger`**: Main workflow handler triggered by cron schedule
   - Fetches Ethereum price from CoinGecko
   - Reads current on-chain NAV value
   - Scales and processes the data
   - Generates and submits CRE report

2. **`fetchResult`**: HTTP request handler for API calls
   - Makes GET request to Price API
   - Parses JSON response
   - Converts float to big.Int (scaled by 1e18)

3. **`generateAndSendReport`**: Report generation and submission
   - ABI-encodes the payload (int256)
   - Generates cryptographic CRE report
   - Submits transaction to blockchain
   - Validates transaction status

### Data Flow
```
Price API → Fetch Price → Generate Report → Submit On-chain → Verify Transaction
```

### Testing
Run the workflow simulation to test changes:
```bash
cd asset-nav
./cre_v0.6.0-alpha workflow simulate --target local-simulation --config config.json main.go --broadcast
```
