# The Blockchain
* [URL](https://www.coursera.org/learn/uciblockchain?action=enroll&specialization=uci-blockchainhttps%3A%2F%2Fwww.coursera.org%2Flearn%2Fuciblockchain%3Fspecialization%3Duci-blockchain)
* Textbook 
  * [The Internet of Money](https://www.amazon.com/Internet-Money-Andreas-M-Antonopoulos/dp/194791006X/ref=tmm_pap_swatch_0?_encoding=UTF8&qid=1559232641&sr=8-3)
  * [Blockchain Basics](https://www.amazon.com/Blockchain-Basics-Non-Technical-Introduction-Steps-dp-1484226038/dp/1484226038/ref=mt_paperback?_encoding=UTF8&me=&qid=1559232841https://www.amazon.com/Blockchain-Basics-Non-Technical-Introduction-Steps-dp-1484226038/dp/1484226038/ref=mt_paperback?_encoding=UTF8&me=&qid=1559232841)

## Module I

### Foundations of Blockchain and Digital Currencies

#### Assignment 1 
**What is blockchain**: The blockchain is essentially a distributed database (ledger) that keeps a record of transactions, 
specifically in cryptocurrency, across multiple machine linked via P2P network.  

**How does it work**: The blockchain is essentially aa distributed database that keeps a record of everything that happens
on the P2P network. Unlike other ledgers, (meta)data on the blockchain is non-fungible (ie cannot be changed).    
* **Step 1**: A user requests a new transaction
* **Step 2**: data is distributed across P2P network 
* **Step 3**: Validate the data using complex algorithm verification
* **Step 4**: Write the transaction on the blockchain (this could be multiple steps)
* ![img.png](blockchain_process.png)

[Investopedia](https://www.investopedia.com/terms/b/blockchain.asp)

**What kind of Blockchain companies are you currently aware of or interested in**: The idea of blockchain can be utilized 
in just about any industry; from finance, to data management nad security. Below is a subset of topics and corresponding 
companies: 
* Private blockchain -- Hyperledger 
* Finance -- Litconin
* Data storage -- filecoin
* DiPin - - Dimo
* IoT -- AIOT
* Services to store and sell crypto -- Coinbase, Binance,  crypto.com

**Do you have a business idea for a new blockchain company**: I'm part of a project called [AnyLog](https://anloy.co) which
allows to manage (IoT and real-time) data directly at the edge. It is using the blockchain to manage the network's metadata
layyer.

#### Quiz
1. Which one of the following actions involving a digital camera is associated with the **APPLICATION** layer?
   * Selecting the zoom level when framing a photograph. 
   * Storing the image in JPEG format. 
   * When saving an image, inserting a timestamp into each image’s file name. 
   * Applying a file compression algorithm to increase storage capacity.

The answer can be found in Drescher: Step 1.

2. Which one of the following actions involving an airline reservation website is associated with the **IMPLEMENTATION** layer?
Holding a reservation for five minutes while the payment is processed.

3. The center node in a centralized cellular telephone system fails. What happens next (choose the BEST answer)?
Nobody will be able to communicate

4. Why are distributed systems (such as the Internet) considered to be more reliable than non-distributed systems?
If one internet server (node) fails, data can be rerouted through alternate pathways.

5. Which of the following is NOT a distributed system?
An automobile’s electrical system.

6. Which one of the following is a characteristic of a digital currency such as bitcoin?
It is used to communicate value, albeit in a decentralized manner

7. Which type of currency system relies on intermediaries to carry out transactions? 
National currencies.

8. What is the fundamental reason for implementing a blockchain?
Ensure integrity in a distributed system.

9. Suppose your local economy is experiencing double-digit inflation. How will the value of your bitcoin holdings change?
It is not possible to determine precisely how bitcoin value will change because it is not tied to national or 
institutional banks.

10. Why is bitcoin so revolutionary? (Choose the BEST answer.)
It allows monetary transactions to take place in a distributed, decentralized manner without needing to trust a central authority.


## Module II
#### Assignment 1 
This week we're considering blockchain architectural structures. There are many different types of distributed system 
architecture configurations. How does the software architecture system allow the blockchain to manage user interactions 
efficiently and transparently?  

The blockchain is a way to manage P2P networks in a decentralized manner using a shared ledger in order to store data 
across the network. Essentially decentralizing storage of data. 

Whether using a tool in a public or private setting, the ultimate goal is a decentralized network that has the ability to
validate transactions via Proof of Work (PoW)1 or Proof of Stake (PoS)2.

Since there are multiple ways by which to utilize / deploy a blockchain, it ultimately depends on the goal of the 
product/services being offered and who the "customers" (whether miners, stakeholder or users) are. For example a 
blockchain network that's used for finance (example LitCoin) may look similar to one used by IoT devices (example IoTA), 
to the naked eye, when in reality the underlying work is substantially different.

Like with any product, the implemented architecture can cause a company to either make-it or break-it, as it influences 
everything from performance, to ease of use, and even who would use the product. So while, blockchain is a great tool 
to build decentralization on to of, how and which flavor is used is the key, rather than just using it as a _keyword_.  

1. Proof of Work is the mechanism for minors to validate new blocks using complex mathematical algorithems 
2. Proof of Stake is the mechanism for validation of new blocks based on their stake (number of coins they hold)

## Module III
### Assignment I 
This week, let’s explore what constitutes a peer-to-peer system. Consider other prominent online peer-to-peer software 
systems and discuss their similarities and differences to a blockchain system such as the Ethereum peer-to-peer system.

Peer-to-Peer (P2P) is the idea that 2 (or more) machines are able to communicate with one another. The most common form 
of P2P is an internet network. The blockchain is similar to the internet as each account has a unique address  
that's used to identify it. The main difference between the internet and the blockchain is that when using the internet
the address (URL) is used ot access the machine one wants to reach. While with the blockchain an address is used to inform 
others who signed the block to the ledger.

When looking at Ethereum in specific, the signed block is usually associated with someone paying / giving a service, with
an agreement that's non-fungible, thus can be legal binding. 


### Quiz 
1. Which one of the following does NOT describe a peer-to-peer system?

One node serves as a dispatcher ensuring reliable point-to-point communication among the others.

2. What is the relationship between the blockchain and a peer-to-peer system?

Software developers sharing code directly with each other.

3. In order for individuals to trust making purchases through a peer-to-peer system (such as eBay), that system needs to provide assurance that it CANNOT be compromised by (choose ALL that apply).

- Technology failures
- Malevolent participants

4. Ideally, one needs to know the number of nodes and their trustworthiness in order to maintain the integrity of a peer-to-peer system. If this is not possible, then:

The blockchain is probably the best option for ensuring system integrity.

5. When we refer to blockchain as a data structure, we mean:

Data organized into units called blocks.

6. When we refer to blockchain as an algorithm, we mean:

The sequence of computer instructions that implement blockchain functionality.

7. When we refer to blockchain as a suite of technologies, we include the following (choose the BEST answer):

Blockchain-data structure, blockchain-algorithm, cryptography, and security technologies.

## Module IV 
### Assignment
This week let’s engage in a discussion together about the tremendous business value that underpins a well designed and 
fully operationalized peer-to-peer system. Discuss Bitcoin or other cryptocurrencies you are familiar with and like. 

Blockchain was a revolution idea that allows for distributed database (ledger) that keeps a record of transactions. The 
original developers of the blockchain want to create a decentralized currency that wasn't dependent on a centralized 
government - Bitcoin. Since then, others have developed other crypto (decentralized) currencies (ex. Ethereum and Litcoin), 
as well as use the blockchain's decentralization for security and data management. 

One of the core technologies that were revolutionized through the blockchain was machine sharing; this can be in terms 
of hard-disk storage (ex. FileCoin) or CPU services for complex querying (ex. CPUcoin). What this means, is that instead 
of paying a cloud services provider (ex. AWS and Google Cloud) an arm and a leg for running things on large servers, 
companies and individual were able to accomplish the same services for cheaper in a more distributed manner. Each of 
these companies released (ICO-ed) their own coin in order to raise funds, while allowing for a unified form of payment
for the service(s) their product allowed for on the blockchain.  




