# wolfcoin

The goal of project wolfcoin is to make a modern golang
currency, in the style of [Algorand](https://www.algorand.com/), but much simpler
and with the [Cardano](https://cardano.org/) style of staking.

We encourage you to rent a server on 
[do](https://m.do.co/c/560b7001e430) or [vultr](https://www.vultr.com/?ref=8507322)
and sell us your CPU cycles and your good name. If you read my [how to bank in the modern era](https://andrewarrow.substack.com/p/in-order-to-bank-in-the-modern-era) post you'll see I come from the background of a former ADA Cardano stake pool operator.

![image](https://wolfschedule.com/assets/26k.png)

I got a bunch of commits in this [other wolf project](https://github.com/andrewarrow/wolfservers/commits/main) and was using it to run 3 different pools, each with a relay.
I finally realized at those numbers, I was better off just getting 6% from Algorand via coinbase.

But something about waiting for these 5 day epochs with Cardano got me thinking about
[this other other wolf project](https://github.com/andrewarrow/wolfschedule) which you
can see live in a modern browser at: [wolfschedule.com](https://wolfschedule.com/) 

If you get some message about WASM it means your browser isn't modern. Try firefox. 
But with the Cardano system, you get your rewards every 5 days. And there is this
big countdown to each of the 5 day periods. I remember many times being
excited about the epoch ending tomorrow or in a few hours. It becomes like pay day.

This code is an adventure into making a genesis file with 30,000 `ed25519` 
public keys, and another file with those same private keys, and giving each
address 1,000,000 wolfcoin (WC) in a var called `books` that's a map from string
to int64.

Everytime a transaction matches the private key, it's writen to a file `tx.log`.

# Design Goals

Try for 100% golang! No need for C please.

When starting a new node from scratch, catchup mode is default.

Very simple and easy to follow path from main.go to where
juicey details like DNS record used to find 1st peer, or how
double spend is prevented, or where rewards are calculated.

Tiny amount of code. The number of lines of go code for wolfcoin should
be 10% of Algorand, maybe even less, maybe 1%.

New people get to start with very little money and earn rewards on
each moon event. With Cardano you won't earn rewards with a small
amount of ADA.

# Calendar

You can read [this post](https://andrewarrow.substack.com/p/moving-a-cli-to-the-browser-with) on how to understand your new calendar. And if you get the WASM running
try the letter `u` or `d` on keyboard to move up or down dates.

There is an event just as important as sunrise and sunset in our lives, think
of it as moonrise, moonset, and they happen about every 15 days. What could
be better? You're getting a huge windfall of newly minted wolfcoins at these
piviotal moments.

If you are skeptical, all I can say is A-B test it. It can help with many things like:

Having trouble picking a date in the future and commiting to something? Now you have
very solid anchor dates available.

Got the Monday thru Friday blues? This shakes everything up. 
Like a wonderful reset. Soon you'll find yourself doing some 
activity in the middle of the day on a Tuesday and before you'd be stuck in that i-must-be-working-and-productive-zombie-zone.

Want to work hard and really acheive something? Now you've got ~15 days! Go for it, run run run!

Want to chill and take a cycle off? That's encouraged.

# Why

On the coinbase application for listing a new currency it asks that big question: why?

I want a simpler golang codebase than Algorand, I want the Cardano staking model, I want true decentralization by encouring anyone with enough money to rent just 1 server to particpate in the proof of stake CPU network, and I want people to break free from
a calendar that's not helping us.

