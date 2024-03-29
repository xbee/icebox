package coins

//go:generate go-enum -f=coins.go --noprefix --marshal --lower

// Coin is an enumeration of coins that are allowed.
// ENUM(
//	BTC=0
//	TEST=1
//	LTC=2
//	DOGE=3
//	RDD=4
//	DSH=5
//	PPC=6
//	NMC=7
//	FTC=8
//	XCP=9
//	BLK=10
//	NSR=11
//	NBT=12
//	MZC=13
//	VIA=14
//	XCH=15
//	RBY=16
//	GRS=17
//	DGC=18
//	CCN=19
//	DGB=20
//	MONA=22
//	CLAM=23
//	XPM=24
//	NEOS=25
//	JBS=26
//	ZRC=27
//	VTC=28
//	NXT=29
//	BURST=30
//	MUE=31
//	ZOOM=32
//	VPN=33
//	CDN=34
//	SDC=35
//	PKB=36
//	PND=37
//	START=38
//	MOIN=39
//	EXP=40
//	EMC2=41
//	DCR=42
//	XEM=43
//	PART=44
//	ARG=45
//	SHR=48
//	GCR=49
//	NVC=50
//	AC=51
//	BTCD=52
//	DOPE=53
//	TPC=54
//	AIB=55
//	EDRC=56
//	SYS=57
//	SLR=58
//	SMLY=59
//	ETH=60
//	ETC=61
//	PSB=62
//	LDCN=63
//	XBC=65
//	IOP=66
//	NXS=67
//	INSN=68
//	OK=69
//	BRIT=70
//	CMP=71
//	CRW=72
//	BELA=73
//	VASH=74
//	FJC=75
//	MIX=76
//	XVG=77
//	EFL=78
//	CLUB=79
//	RICHX=80
//	POT=81
//	QRK=82
//	TRC=83
//	GRC=84
//	AUR=85
//	IXC=86
//	NLG=87
//	BITB=88
//	BTA=89
//	XMY=90
//	BSD=91
//	UNO=92
//	MTR=93
//	GB=94
//	SHM=95
//	CRX=96
//	BIQ=97
//	EVO=98
//	STO=99
//	BIGUP=100
//	GAME=101
//	DLC=102
//	ZYD=103
//	DBIC=104
//	STRAT=105
//	SH=106
//	MARS=107
//	UBQ=108
//	PTC=109
//	NRO=110
//	ARK=111
//	USC=112
//	THC=113
//	LINX=114
//	ECN=115
//	DNR=116
//	PINK=117
//	PIGGY=118
//	PIVX=119
//	FLASH=120
//	ZEN=121
//	PUT=122
//	ZNY=123
//	UNIFY=124
//	XST=125
//	BRK=126
//	VC=127
//	XMR=128
//	VOX=129
//	NAV=130
//	FCT=131
//	EC=132
//	ZEC=133
//	LSK=134
//	STEEM=135
//	XZC=136
//	SBTC=137
//	RPT=139
//	LBC=140
//	KMD=141
//	BSQ=142
//	RIC=143
//	XRP=144
//	BCH=145
//	NEBL=146
//	ZCL=147
//	XLM=148
//	NLC2=149
//	WHL=150
//	ERC=151
//	DMD=152
//	BTM=153
//	BIO=154
//	XWC=155
//	BTG=156
//	BTC2X=157
//	SSN=158
//	TOA=159
//	BTX=160
//	ACC=161
//	BCO=162
//	ELLA=163
//	PIRL=164
//	XRB=165
//	VIVO=166
//	FRST=167
//	HNC=168
//	BUZZ=169
//	MBRS=170
//	HSR=171
//	HTML=172
//	ODN=173
//	ONX=174
//	RVN=175
//	GBX=176
//	BTCZ=177
//	POA=178
//	NYC=179
//	MXT=180
//	WC=181
//	MNX=182
//	BTCP=183
//	MUSIC=184
//	BCA=185
//	CRAVE=186
//	STAK=187
//	WBTC=188
//	LCH=189
//	EXCL=190
//	LCC=192
//	XFE=193
//	EOS=194
//	TRX=195
//	KOBO=196
//	HUSH=197
//	BANANO=198
//	ETF=199
//	OMNI=200
//	BIFI=201
//	UFO=202
//	BIS=209
//	NEET=210
//	BOPO=211
//	BOXY=215
//	BITG=222
//	ASK=223
//	SMART=224
//	XUEZ=225
//	VAR=233
//	UC=247
//	NANO=256
//	ARA=312
//	RAP=321
//	BLOCK=328
//	MEM=333
//	PHR=444
//	KOTO=510
//	XRD=512
//	BCS=555
//	ACT=666
//	SSC=668
//	BTW=777
//	BEET=800
//	QVT=808
//	CLO=820
//	ADT=886
//	NEO=888
//	LBTC=998
//	BCD=999
//	BTN=1000
//	BBC=1111
//	CDY=1145
//	DFC=1337
//	BCX=1688
//	XTZ=1729
//	ADA=1815
//	XMX=1977
//	EGEM=1987
//	HODL=1989
//	QTUM=2301
//	ETP=2302
//	DEO=3552
//	NAS=2718
//	AXE=4242
//	BPA=6666
//	SAFE=6688
//	BTV=7777
//	BTQ=8339
//	BTP=8999
//	BTF=9888
//	BTCS=33878
//	ORT=88888
//	WICC=99999
//	LAX=1712144
//	WAN=5718350
//	WAVES=5741564
// )
type Coin uint32
