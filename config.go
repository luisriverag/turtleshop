package main

var (
	storeName        = "The TurtleCoin Store"
	storeURL         = "http://127.0.0.1:5000"
	storeAddr        = "TRTLuwXBudSBGfc9LPKYwBLwMWTAkTN1pVDtqebvJcBXB5XrAs6iiREaqt2ycWsh2Y7xs6fctehBUDPyXcV2YHfYXyY1ysjQsaH"
	storeView        = "f0affadb8d4ca354ed18b9cdfc8e5486fdbcfa87c74ad5a395d366c361980f00"
	storeCallBackURL = storeURL + "/pay/"
	storeSecureSalt  = "This could be anything, but its best to be long and crazy."
)
var sc ShoppingCart
var wantsLicense bool
var (
	touchThisAndYourMomDiesInHerSleepTonight = boldbrightgreen + `
TRTL-OSLv0 Draft License` + brightgreen + `

Copyright 2020-2021 The TurtleCoin® Developers` + white + `

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

` + white + `1)` + nc + ` Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.

` + white + `2)` + nc + ` Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation and/or
other materials provided with the distribution.

` + white + `3)` + nc + ` Neither the name of the copyright holder nor the names of its contributors
may be used to endorse or promote products derived from this software without
specific written permission.

` + white + `4)` + nc + ` Any modifications, improvements, enhancements, bugfixes, features, or other
changes (hereafter known as "modifications") made to this software must be
submitted to The TurtleCoin® Developers, in writing via a pull-request, which
may or may not be accepted, to the matching source code repository no more than
seven (7) days after such modifications have been made and any such
modifications must be made available under this same license.

` + white + `5)` + nc + ` This software, including its source code, binaries, and documentation, may
not be used to create, manage, maintain, develop, interact with, or otherwise
operate a cryptocurrency network other than the TurtleCoin® network or an
ephemeral TurtleCoin® Testnet (which may only operate in the support of 
testing proposed changes to the TurtleCoin® network) without the explicit 
written consent of The TurtleCoin® Developers.
` + white + `
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS” AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.` + nc + `
	`
)
