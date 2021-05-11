# wolfcoin

The goal of project wolfcoin is to make a modern golang
currency, in the style of [Algorand](https://www.algorand.com/), but much simpler
and with the [Cardano](https://cardano.org/) style of staking.

We encourage you to rent a server on 
[do](https://m.do.co/c/560b7001e430) or [vultr](https://www.vultr.com/?ref=8507322)
and sell us your CPU cycles and your good name. If you read my [how to bank in the modern era](https://andrewarrow.substack.com/p/in-order-to-bank-in-the-modern-era) post you'll see I come from the background of a former ADA Cardano stake pool operator.

![image](https://wolfschedule.com/assets/26k.png)

I got a bunch of commits in this [other wolf project](https://github.com/andrewarrow/wolfservers/commits/main) and was using it to run 3 different pools, each with a relay.
I finally realized at those numbers, I was better off just get 6% from Algorand via coinbase.

But something about waiting for these 5 day epochs with Cardano got me thinking about
[this other other wolf project](https://github.com/andrewarrow/wolfschedule) which you
can see live in a modern browser at: [wolfschedule.com](https://wolfschedule.com/) 

If you get some message about WASM it means your browser isn't modern. Try firefox. 
There is an event just as important as sunrise and sunset in our lives, think
of it as moonrise, moonset, and they happen about every 15 days. 

This code is an adventure into making a genesis file with 30,000 `ed25519` 
public keys, and another file with those same private keys, and giving each
address 1,000,000 wolfcoin (WC) in a var called `books` that's a map from string
to int64.

Everytime a transaction matches the private key, it's writen to a file `tx.log`


