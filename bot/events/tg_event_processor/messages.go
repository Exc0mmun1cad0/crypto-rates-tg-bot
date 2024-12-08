package telegram

const msgHelp = `
This bot is created to output latest rates of 100 popular cryptocurrencies
It supports following commands:
	/start
	/help
	/rates
	/list
	/add
	/delete
	/deleteall
`

const (
	msgHello         = `Hello and welcome to my first adequate tg bot. For more info type /help`
	msgUknownCommand = `Unknown command`
	msgRates         = "Date & Time: %s\nCurrent cryptocurrencies rate in USD:\n"
	msgToAdd         = `enter names of cryptocurrencies to add to your list`
	msgToDel         = `enter names of cryptocurrencies to delete from your list`
	msgDelAll        = `your cryptocurrency list was cleared`
	msgAdd           = `following cryptocurrencies were added to your list:`
	msgDel           = `following cryptocurrencies were deleted from your list:`
	msgNoTokens      = `you have no cryptocurrencies in your list`
	msgListInfo      = "Following cryptocurrencies are available:\n"
	msgList          = `1.  bitcoin
2.  ethereum
3.  tether
4.  binance-coin
5.  usd-coin
6.  xrp
7.  solana
8.  cardano
9.  dogecoin
10. tron
11. multi-collateral-dai
12. polygon
13. polkadot
14. wrapped-bitcoin
15. litecoin
16. bitcoin-cash
17. chainlink
18. shiba-inu
19. unus-sed-leo
20. trueusd
21. avalanche
22. stellar
23. monero
24. okb
25. uniswap
26. ethereum-classic
27. binance-usd
28. cosmos
29. bitcoin-bep2
30. filecoin
31. internet-computer
32. maker
33. lido-dao
34. crypto-com-coin
35. vechain
36. quant
37. near-protocol
38. aave
39. stacks
40. bitcoin-sv
41. the-graph
42. algorand
43. hedera-hashgraph
44. render-token
45. tezos
46. the-sandbox
47. eos
48. injective-protocol
49. theta
50. axie-infinity
51. elrond-egld
52. xinfin-network
53. thorchain
54. decentraland
55. fantom
56. kava
57. ecash
58. paxos-standard
59. neo
60. synthetix-network-token
61. flow
62. trust-wallet-token
63. chiliz
64. kucoin-token
65. zcash
66. fei-protocol
67. frax-share
68. iota
69. klaytn
70. curve-dao-token
71. rocket-pool
72. huobi-token
73. mina
74. conflux-network
75. gatetoken
76. casper
77. dydx
78. ftx-token
79. gala
80. compound
81. nexo
82. wootrade
83. zilliqa
84. dash
85. 1inch
86. oasis-network
87. basic-attention-token
88. arweave
89. gnosis-gno
90. pancakeswap
91. nem
92. qtum
93. aelf
94. gemini-dollar
95. holo
96. just
97. loopring
98. convex-finance
99. celo
100. enjin-coin`
)
