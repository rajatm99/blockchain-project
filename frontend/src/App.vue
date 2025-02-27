<script>
import { ethers } from "ethers"

export default {
  name: 'App',
  data() {
    return {
      walletConnected: false,
      projectName: 'TOKEN PROJECT',
      ethereum: null,
      account: null,
      abi: null,
      address: null,
      balance: null,
      transferAddress: '',
      transferAmount: 0,
      contract: null,
      balanceLoading: false,
    };
  },
  mounted() {
    this.onPageLoad();
  },
  methods: {
    async onPageLoad() {
      // todo import ABI from file
      this.address = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
      this.abi = [
        {
          "inputs": [],
          "stateMutability": "nonpayable",
          "type": "constructor"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "spender",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "allowance",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "needed",
              "type": "uint256"
            }
          ],
          "name": "ERC20InsufficientAllowance",
          "type": "error"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "sender",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "balance",
              "type": "uint256"
            },
            {
              "internalType": "uint256",
              "name": "needed",
              "type": "uint256"
            }
          ],
          "name": "ERC20InsufficientBalance",
          "type": "error"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "approver",
              "type": "address"
            }
          ],
          "name": "ERC20InvalidApprover",
          "type": "error"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "receiver",
              "type": "address"
            }
          ],
          "name": "ERC20InvalidReceiver",
          "type": "error"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "sender",
              "type": "address"
            }
          ],
          "name": "ERC20InvalidSender",
          "type": "error"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "spender",
              "type": "address"
            }
          ],
          "name": "ERC20InvalidSpender",
          "type": "error"
        },
        {
          "anonymous": false,
          "inputs": [
            {
              "indexed": true,
              "internalType": "address",
              "name": "owner",
              "type": "address"
            },
            {
              "indexed": true,
              "internalType": "address",
              "name": "spender",
              "type": "address"
            },
            {
              "indexed": false,
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            }
          ],
          "name": "Approval",
          "type": "event"
        },
        {
          "anonymous": false,
          "inputs": [
            {
              "indexed": true,
              "internalType": "address",
              "name": "from",
              "type": "address"
            },
            {
              "indexed": true,
              "internalType": "address",
              "name": "to",
              "type": "address"
            },
            {
              "indexed": false,
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            }
          ],
          "name": "Transfer",
          "type": "event"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "owner",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "spender",
              "type": "address"
            }
          ],
          "name": "allowance",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "spender",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            }
          ],
          "name": "approve",
          "outputs": [
            {
              "internalType": "bool",
              "name": "",
              "type": "bool"
            }
          ],
          "stateMutability": "nonpayable",
          "type": "function"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "account",
              "type": "address"
            }
          ],
          "name": "balanceOf",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "decimals",
          "outputs": [
            {
              "internalType": "uint8",
              "name": "",
              "type": "uint8"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [
            {
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            },
            {
              "internalType": "address",
              "name": "to",
              "type": "address"
            }
          ],
          "name": "mint",
          "outputs": [],
          "stateMutability": "nonpayable",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "name",
          "outputs": [
            {
              "internalType": "string",
              "name": "",
              "type": "string"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "owner",
          "outputs": [
            {
              "internalType": "address",
              "name": "",
              "type": "address"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "symbol",
          "outputs": [
            {
              "internalType": "string",
              "name": "",
              "type": "string"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [],
          "name": "totalSupply",
          "outputs": [
            {
              "internalType": "uint256",
              "name": "",
              "type": "uint256"
            }
          ],
          "stateMutability": "view",
          "type": "function"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "to",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            }
          ],
          "name": "transfer",
          "outputs": [
            {
              "internalType": "bool",
              "name": "",
              "type": "bool"
            }
          ],
          "stateMutability": "nonpayable",
          "type": "function"
        },
        {
          "inputs": [
            {
              "internalType": "address",
              "name": "from",
              "type": "address"
            },
            {
              "internalType": "address",
              "name": "to",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "value",
              "type": "uint256"
            }
          ],
          "name": "transferFrom",
          "outputs": [
            {
              "internalType": "bool",
              "name": "",
              "type": "bool"
            }
          ],
          "stateMutability": "nonpayable",
          "type": "function"
        }
      ]

      await this.connectWallet()
      await this.getBalance()
    },
    async connectWallet() {
      this.ethereum = window.ethereum
      if (!this.ethereum) {
        console.error("NO WALLET FOUND")
        alert("install metamask")
        return
      }
      const accounts = await this.ethereum.request({ method: "eth_requestAccounts" });

      if (accounts.length == 0) {
        alert("wallet account not found")
        return
      }

      this.account = accounts[0]

      const provider = new ethers.BrowserProvider(window.ethereum);
      await provider.send("eth_requestAccounts", []);
      const signer = await provider.getSigner();

      this.contract = new ethers.Contract(this.address, this.abi, signer);

      this.walletConnected = true

    },
    async getBalance() {
      this.balanceLoading = true;
      try {
        this.balance = await this.contract.balanceOf(this.account)
      } catch (e) {
        console.log("error", e)
      } finally {
        this.balanceLoading = false;
      }
    },
    async transfer() {

      try {
        const tx = await this.contract.transfer(this.transferAddress, this.transferAmount)
        tx.wait()

      } catch (e) {
        console.log("error - ", e)
      }
    },

  }
}
</script>

<template>
  <div class="dark-theme">
    <header class="header">
      <h1>{{ projectName }}</h1>
      <button v-if="!this.walletConnected" @click="connectWallet" class="connect-button">Connect Wallet</button>
      <div v-else>
        Connected : {{ this.account }}
      </div>
    </header>
    <main class="main-content">
      <div class="balance">
        <h2 v-if="!balanceLoading">Balance : {{ balance }} RTN</h2>
        <h2 v-else> Loading Balance ... </h2>

        <div class="refresh-button" @click="getBalance">
          <img src="@/assets/refresh-svgrepo-com.svg" alt="Description of SVG" />
        </div>
      </div>
      <div class="transfer-form">
        <h2>Transfer Tokens</h2>
        <input v-model="transferAddress" type="text" placeholder="Recipient Address" />
        <input v-model="transferAmount" type="number" placeholder="Amount" />
        <button @click="transfer" class="transfer-button">Transfer</button>
      </div>
    </main>
  </div>
</template>

<style scoped>
.dark-theme {
  background-color: #121212;
  color: #ffffff;
  font-family: Arial, sans-serif;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: flex-start;
}

.header {
  width: 100%;
  padding: 20px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #1e1e1e;
}

.connect-button {
  background-color: #007bff;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 5px;
  cursor: pointer;
}

.main-content {
  padding: 20px;
  width: 100%;
  max-width: 600px;
  background-color: #1e1e1e;
  border-radius: 10px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.5);
  margin-top: 20px;
}

.balance {
  margin-bottom: 20px;
  display: flex;
  justify-content: center;
}

.transfer-form {
  display: flex;
  flex-direction: column;
}

.transfer-form input {
  margin-bottom: 10px;
  padding: 10px;
  border: none;
  border-radius: 5px;
}

.transfer-button {
  background-color: #28a745;
  color: white;
  border: none;
  padding: 10px;
  border-radius: 5px;
  cursor: pointer;
}

.refresh-button {
  height: 20px;
  width: 20px;
  margin-left: 20px;
  cursor: pointer;
}

.refresh-button img {
  height: 100%;
  width: 100%;
}

</style>
