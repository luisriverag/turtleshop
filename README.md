# TurtleShop
A simple web shop that accepts TurtleCoin
---
![image](https://user-images.githubusercontent.com/34389545/116630530-f0407f00-a918-11eb-8e00-e929699ee25e.png)


### Note: This software is not finished. 

# Donate to help development of TurtleShop
**TRTLuxEnfjdF46cBoHhyDtPN32weD9fvL43KX5cx2Ck9iSP4BLNPrJY3xtuFpXtLxiA6LDYojhF7n4SwPNyj9M64iTwJ738vnJk**


# How does it work?
When a user clicks BUY on an item, a request is sent to TurtlePay.io to help create the invoice and the store presents that back to the user. TurtlePay then watches the blockchain using the view key you provide to check for transactions bound to the payment address provided. 


# How to use

1. Open `inventory.go` and edit the items to create your inventory
2. Open `config.go` and edit the items to suit your store
3. Compile the software using `go build`
4. Run the TurtleShop binary in the same folder containing /templates/
5. Open `127.0.0.1:5000/` in a web browser to see your store

# How to view the license
Run `TurtleShop` with the `--license` flag appended to the end, like this:

```
./TurtleShop --license
```

# Folders
The first time you run TurtleShop, it will create a few folders that it uses to manage the invoices

- **invoices/** # invoices awaiting payment go here
- **sales/** # invoices with completed payments get moved here
- **templates/** # this is where the web views for your store are located

# Contributing
Other than sending a direct donation in TRTL, if you have some dev expertise, you can help.
**TRTLuxEnfjdF46cBoHhyDtPN32weD9fvL43KX5cx2Ck9iSP4BLNPrJY3xtuFpXtLxiA6LDYojhF7n4SwPNyj9M64iTwJ738vnJk**


TurtleShop has a few things on the todo list:

- Create a page the user can look at to see the progress of their payment by pressing refresh
- Create an `invoice/{{invoiceid}}` endpoint and handler so that past invoices can persist publicly
- Finish shopping cart functionality

TurtleShop is written in Go and uses HTML with Bootstrap CSS framework for the views. I prefer that any new features or contributions be free of javascript and cookies, and that any of the Go contributions consist only of STD library, no new imports.

# Thanks!
![image](https://user-images.githubusercontent.com/34389545/116630607-1a923c80-a919-11eb-9d17-c1d0afb1e698.png)
