# fabric-sample-invoice

## Development Environment:
#### go version go1.11.5 linux/amd64
#### OS: Xubuntu (Ubuntu 18.04.1 LTS)


### NOTE: GOPATH must be correctly set in your local machine
You can check if go is properly installed by running **go version** in your terminal
<br />
Check also if gopath is correct. Run **echo $GOPATH**

Make sure you have a **basic-network** from fabric-sample if not click here[`click here`](https://github.com/hyperledger/fabric-samples) to get git clone fabric-sample copy the basic-network and paste it in the folder fabric-sample-invoice

Also, you need you need to git clone [`blockchain-training-labs`](https://github.com/bchinc/blockchain-training-labs)


### Setup Network
#### Step1:
Create a folder **invoice** under **fabric-sample-invoice/**
<br />
Then copy all files under blockchain-training-labs/**node/** in this repository then paste it to the directory
<br />
<br />
**NOTE:** Do not include **node_modules/**

##### Output:

<pre>fabric-sample-invoice/
  | ---- invoice/
              | ---- app.js
              | ---- enrollAdmin.js
              | ---- package.json
              | ---- registerUser.js
              | ---- startFabric.sh
</pre>

#### Step2:
Create a folder **invoice** under **fabric-sample-invoice/chaincode/**
<br />
Then copy the **go/** folder in this repository or simply create a new folder named **go/** inside **fabric-sample-invoice/chaincode/invoice/** then paste it to the directory
<br />

##### Output:

<pre>fabric-sample-invoice/
  | ---- chaincode/
              | ---- invoice/
                          | ---- go/
                                      | ---- invoice.go
</pre>

#### Step3:
Open terminal then change directory to **/fabric-sample-invoice/invoice/**
Then run **./startFabric.sh**
<br />
Then run **npm install**
<br />
Then run **node enrollAdmin.js**
<br />
Then run **node registerUser.js**
<br />
Then run **node app.js**
<br />

**./startFabric.sh**
<br />
You should see something like this in your terminal
<br />  
![alt text](https://github.com/jenriellegaon/blockchain-training-labs/blob/master/startFabric.png)
<br />
<br />

**node enrollAdmin.js**
<br />
You should see something like this in your terminal
<br />
![alt text](https://github.com/jenriellegaon/blockchain-training-labs/blob/master/enrollAdmin.png)
<br />
<br />

**node registerUser.js**
<br />
You should see something like this in your terminal
<br />
![alt text](https://github.com/jenriellegaon/blockchain-training-labs/blob/master/registerUser.png)
<br />
<br />

**node app.js**
<br />
You should see something like this in your terminal
<br />
![alt text](https://github.com/jenriellegaon/blockchain-training-labs/blob/master/app.png)
<br />
<br />

#### Step4:
Test the endpoints using **POSTMAN** or **INSOMNIA REST Client**
<br />
<br />

#### Testing Endpoints

#### Display All Invoices
##### http://localhost:3000/
##### Use the GET http request in this function as we are getting data
<br />
<br />
<br />

#### Raise Invoice
##### http://localhost:3000/invoice
##### Use the POST http request in this function as we are pushing data
Select **Form URL Encoded** as a structure
##### Parameters
+ invoiceid
+ invoicenum
+ billedto
+ invoicedate
+ invoiceamount
+ itemdescription
+ gr
+ ispaid
+ paidamount
+ repaid
+ repaymentamount

##### NOTE: gr , ispaid , paidamount , repaid , repaymentamount default values are as follows N , N , 0 , N , 0
**gr = N**
<br />
**ispaid = N**
<br />
**paidamount = 0**
<br />
**repaid = N**
<br />
**repaymentamount = 0**
<br />
<br />
<br />

#### Goods Received
##### http://localhost:3000/invoice
##### Use the PUT http request in this function as we are modifying a data
Select **Form URL Encoded** as a structure

##### Parameters
+ invoiceid
+ gr
<br />
<br />
<br />

#### Bank Payment to Supplier
##### http://localhost:3000/invoice
##### Use the PUT http request in this function as we are modifying a data
Select **Form URL Encoded** as a structure

##### Parameters
+ invoiceid
+ ispaid
<br />
<br />
<br />

#### OEM Repays to Bank
##### http://localhost:3000/invoice
##### Use the PUT http request in this function as we are modifying a data
Select **Form URL Encoded** as a structure

##### Parameters
+ invoiceid
+ repaid
<br />
<br />
<br />


#### Next steps will be published soon# fabric-sample-invoice
