# testing

1. nodeA gets valid tx FROM A to B, Amount 100
2. nodeB gets valid tx from A to C, Amount 100

Total amount A has is Amount 100.

3. nodeC gets nodeA msg about tx A2B
4. nodeC gets nodeB msg about tx A2C

nodeC is first to notice the double spend.

it records A2B ok, but rejects A2C.

5. nodeD gets nodeB msg about tx A2C
6. nodeD gets nodeB msg about tx A2B

nodeD is 2nd to notice the double spend.

it records A2C ok, but rejects A2B.




