---
title: Testnet Deployment Guidelines
sidebar_title: 2. Deploy to Testnet
description: Important information for developers using Flow Testnet
---

Flow Testnet exists to help developers test their software and smart contracts against a live network. It's also used by the Flow core team a means of releasing and testing new protocol and smart contract features before these are integrated into Flow's main network (Mainnet).

When the Flow protocol is updated or a new version of Cadence is released, those updates will always be made available on the the Flow Emulator _before_ they're integrated into Flow Testnet or Flow Mainnet.

## Risks for Early Adopters

### Breaking Changes

In these early days of Flow, version updates to Cadence, Flow Node software, and the Flow SDKs will often contain important updates as well as breaking changes.

Breaking changes will be a fact of life for early adopters of Flow's development stack, and you may often need to manually re-deploy your application's contracts after updates. But that's OK, we're here to help!

### Re-Deploying Your Application

If you discover your application is broken after an update, use the latest emulator to test changes to your application. Once you are satisfied that you've patched any breaking changes, you'll need to get in touch with the Flow team to redeploy your contracts to Testnet.

Before re-deploying we'll ask you to follow the steps listed in the <br/> [**Testnet Testing Guidelines.**](../testnet-testing)

<!-- ### Version Compatibility

A version compatibility table can be found here. This will help you navigate version compatibility between Emulator, SDK, and Network Node (flow-go) versions. -->

### **Testnet Sporking**

"Sporking" (soft forking) is the process of upgrading Flow network node software and migrating the state from the previous version.

Currently, **historical event data is not migrated between sporks,** so you'll need to design your application with this in mind. We recognize the usefulness of historical event data and plan on adding a means of accessing it in the near future.

### "Real Value" on Testnet

Flow Testnet is explicitly for experimentation and testing and should not be used to exchange "real value" (For example, developing a fiat money on/off-ramp for your testnet application). Accounts using the testnet to exchange real value will be subject to removal.

## Getting Started on Testnet

Before using Flow Testnet, you'll need to get in touch with the Flow team so we can help you bootstrap your application onto the network (we'll cover how to do this below). This is a temporary restriction while Flow is under heavy development that will help us keep the network stable and help us best support Flow's early adopters.

### **Accessing Flow Testnet**

To connect to the testnet from your code, use the following host address:

- `access.devnet.nodes.onflow.org:9000`


Code Example:

```go
import "github.com/onflow/flow-go-sdk/client"

func main() {
  flowAccessHost := "access.devnet.nodes.onflow.org:9000"
  flowClient, _ := client.New(flowAccessHost, grpc.WithInsecure())

  // ...
}
```

### **Creating Accounts & Deploying Contracts**

Account creation and contract deployment on Flow Testnet are currently restricted to authorized users. To create an account, you'll need to generate an asymmetric cryptographic key pair. Flow currently supports the ECDSA signature algorithm and both P256 and secp256k1 curves; you can generate key pairs using the Flow CLI tool's `keys generate` command. See here for how to install the flow cli: [https://github.com/onflow/flow-cli](https://github.com/onflow/flow-cli)

Once you've generated a key pair, visit [https://testnet-faucet.onflow.org/](https://testnet-faucet.onflow.org/) and input the _public key_ into the "Create Account" form at the top of the page.

Flow isn't responsible for securing and storing the private keys for testnet accounts. Your application will need to have access to the private key counterparts to the public keys you submitted with your account creation request, in order to carry out transactions on Flow using the SDKs.

Before deploying your smart contracts, we'll ask you to follow the steps listed here:

**Testnet Testing Guidelines**

### **Acquiring FLOW Tokens for Development**

Once your accounts have been created, you can request testnet FLOW through the same page ([https://testnet-faucet.onflow.org/](https://testnet-faucet.onflow.org/)).

When your token request has been approved you'll receive an email to the address you submitted confirming the token deposit to the requested account!

Likewise, you'll receive an email if there were problems with your request.

## Developing on Testnet

Developing on Flow Testnet in the early days of Flow will often require cooperation between your team and Flow's core developers.

Flow Testnet comes with some useful contracts already deployed. We're calling these the **core contracts.** There are reference pages with import addresses for the core contracts here: [https://docs.onflow.org/protocol/core-contracts/](/protocol/core-contracts/)

Once you're accounts are set up, and you're ready to develop, you can look over some code examples from the `flow-go-sdk` here:

- [https://github.com/onflow/flow-go-sdk/tree/master/examples](https://github.com/onflow/flow-go-sdk/tree/master/examples)