package xsorm

import "encoding/base64"

// 由于xstpl无法在没有本地文件的情况下解析到build中的方法，需要在本项目中导出
// 直接使用方法，将那个文件直接导出, 为了方便解决`符号的输入问题，将文件直接base64编码
// openssl base64 -in build.go
func ExportBuildFile() string{
	b := `
cGFja2FnZSB4c29ybQoKaW1wb3J0ICgKCSJkYXRhYmFzZS9zcWwiCgkiZ2l0aHVi
LmNvbS9qaW56aHUvZ29ybSIKCSJnaXRodWIuY29tL3h4aWlhYXNzL2l1dGlscyIK
CSJyZWZsZWN0IgoJInNvcnQiCgkic3RyY29udiIKCSJzdHJpbmdzIgopCgp0eXBl
IHdoZXJlcyBzdHJ1Y3QgewoJdHAgICAgICAgc3RyaW5nCgljb2x1bW4gICBzdHJp
bmcKCW9wZXJhdG9yIHN0cmluZwoJdmFsdWUgICAgaW50ZXJmYWNle30KCXZhbHVl
cyAgIFtdaW50ZXJmYWNle30KCWJvb2xlYW4gIHN0cmluZwoJc3FsICAgICAgc3Ry
aW5nCn0KCnR5cGUgam9pbnMgc3RydWN0IHsKCXRwICAgICBzdHJpbmcKCXRhYmxl
ICBzdHJpbmcKCW9uZSAgICBzdHJpbmcKCW9wICAgICBzdHJpbmcKCXR3byAgICBz
dHJpbmcKCXdoZXJlcyBbXXdoZXJlcwp9Cgp0eXBlIG9yZGVyIHN0cnVjdCB7Cglj
b2x1bW4gICAgc3RyaW5nCglkaXJlY3Rpb24gc3RyaW5nCn0KCnR5cGUgQnVpbGQg
c3RydWN0IHsKCW1vZGVsICAgICAgICAgaW50ZXJmYWNle30KCXRhYmxlICAgICAg
ICAgc3RyaW5nCgl3aGVyZXMgICAgICAgIFtdd2hlcmVzCgliaW5kaW5ncyAgICAg
IFtdaW50ZXJmYWNle30KCWpvaW5zICAgICAgICAgW11qb2lucwoJc2VsZWN0cyAg
ICAgICBbXXN0cmluZwoJZ3JvdXAgICAgICAgICBzdHJpbmcKCXNldHMgICAgICAg
ICAgaXV0aWxzLkgKCW9yZGVycyAgICAgICAgW11vcmRlcgoJb2Zmc2V0ICAgICAg
ICBpbnQKCWxpbWl0ICAgICAgICAgaW50CgllcnIgICAgICAgICAgIGVycm9yCglu
b3RQYW5pYyAgICAgIGJvb2wgICAvLyDlvZNzcWzmiafooYzkuI3miJDlip/ml7bv
vIzmmK/lkKbmipvlh7rplJnor68KCWVmZmVjdFJvdyAgICAgaW50NjQgIC8vIOaJ
p+ihjHNxbOW9seWTjeeahOihjOaVsAoJaXNTZXRTcGxpdFZhbCBib29sICAgLy8g
5piv5ZCm6K6+572u6L+H5YiG6KGo5YC877yM5Zug5Li6c3BsaXRWYWzmnInlj6/o
g73orr7kuLrliJ3lp4vlgLwwCglzcGxpdFZhbCAgICAgIGludDY0ICAvLyDliIbo
oajlgLwKCWNvbiAgICAgICAgICAgc3RyaW5nIC8vIOaMh+WumuS9v+eUqOeahOi/
nuaOpQoJaXNPbldyaXRlICAgICBib29sICAgLy8g5piv5ZCm5by65Yi26K+75Li7
5bqTCgl1bnNjb3BlZCAgICAgIGJvb2wKCW9wZXJhdGUgICAgICAgKm9wZXJhdAp9
CnR5cGUgcmF3IHN0cnVjdCB7CglzcWwgc3RyaW5nCn0KCmZ1bmMgUmF3KHNxbCBz
dHJpbmcpICpyYXcgewoJcmV0dXJuICZyYXd7c3FsfQp9CgovLyDmnoTlu7p3aGVy
ZeivreWPpeeahOWbnuiwgwp0eXBlIFdoZXJlQ2IgZnVuYyhidWlsZCAqQnVpbGQp
Cgp0eXBlIHRhYmxlciBpbnRlcmZhY2UgewoJVGFibGVOYW1lKCkgc3RyaW5nCn0K
CmZ1bmMgTmV3QnVpbGQobW9kZWwgaW50ZXJmYWNle30pICpCdWlsZCB7CglidWls
ZCA6PSBuZXcoQnVpbGQpCglidWlsZC5zZWxlY3RzID0gW11zdHJpbmd7fQoKCWlm
IHN0ciwgb2sgOj0gbW9kZWwuKHN0cmluZyk7IG9rIHsKCQkvLyDmjIflrprooajl
kI0KCQlidWlsZC5UYWJsZU5hbWUoc3RyKQoJfSBlbHNlIHsKCQlidWlsZC5tb2Rl
bCA9IG1vZGVsCgl9CglidWlsZC53aGVyZXMgPSBbXXdoZXJlc3t9CglidWlsZC5z
ZXRzID0gbWFrZShpdXRpbHMuSCkKCWJ1aWxkLmJpbmRpbmdzID0gW11pbnRlcmZh
Y2V7fXt9CglidWlsZC5vcGVyYXRlID0gbmV3T3BlcmF0KCkKCXJldHVybiBidWls
ZAp9CgovLyDlpI3liLblvZPliY3nmoTlhoXlrrnlh7rkuIDkuKrmlrDnmoRidWls
ZOWvueixoQpmdW5jIChidWlsZCAqQnVpbGQpIENsb25lKCkgKkJ1aWxkIHsKCXJl
dHVybiAmQnVpbGR7CgkJbW9kZWw6ICAgICAgICAgYnVpbGQubW9kZWwsCgkJdGFi
bGU6ICAgICAgICAgYnVpbGQudGFibGUsCgkJd2hlcmVzOiAgICAgICAgYnVpbGQu
d2hlcmVzLAoJCWJpbmRpbmdzOiAgICAgIGJ1aWxkLmJpbmRpbmdzLAoJCWpvaW5z
OiAgICAgICAgIGJ1aWxkLmpvaW5zLAoJCXNlbGVjdHM6ICAgICAgIGJ1aWxkLnNl
bGVjdHMsCgkJZ3JvdXA6ICAgICAgICAgYnVpbGQuZ3JvdXAsCgkJc2V0czogICAg
ICAgICAgYnVpbGQuc2V0cywKCQlvcmRlcnM6ICAgICAgICBidWlsZC5vcmRlcnMs
CgkJb2Zmc2V0OiAgICAgICAgYnVpbGQub2Zmc2V0LAoJCWxpbWl0OiAgICAgICAg
IGJ1aWxkLmxpbWl0LAoJCWVycjogICAgICAgICAgIGJ1aWxkLmVyciwKCQlub3RQ
YW5pYzogICAgICBidWlsZC5ub3RQYW5pYywKCQllZmZlY3RSb3c6ICAgICBidWls
ZC5lZmZlY3RSb3csCgkJaXNTZXRTcGxpdFZhbDogYnVpbGQuaXNTZXRTcGxpdFZh
bCwKCQlzcGxpdFZhbDogICAgICBidWlsZC5zcGxpdFZhbCwKCQljb246ICAgICAg
ICAgICBidWlsZC5jb24sCgkJaXNPbldyaXRlOiAgICAgYnVpbGQuaXNPbldyaXRl
LAoJCXVuc2NvcGVkOiAgICAgIGJ1aWxkLnVuc2NvcGVkLAoJfQp9CgpmdW5jIChi
dWlsZCAqQnVpbGQpIFVuc2NvcGVkKCkgKkJ1aWxkIHsKCWJ1aWxkLnVuc2NvcGVk
ID0gdHJ1ZQoJcmV0dXJuIGJ1aWxkCn0KCmZ1bmMgKGJ1aWxkICpCdWlsZCkgU2Vs
ZWN0KGNvbHVtbnMgLi4uaW50ZXJmYWNle30pICpCdWlsZCB7Cglmb3IgXywgc3Ry
IDo9IHJhbmdlIGNvbHVtbnMgewoJCXN3aXRjaCBzdHIuKHR5cGUpIHsKCQljYXNl
IHN0cmluZzoKCQkJcyA6PSBzdHIuKHN0cmluZykKCQkJaWYgc1swOjFdID09ICJg
IiB7CgkJCQlidWlsZC5zZWxlY3RzID0gYXBwZW5kKGJ1aWxkLnNlbGVjdHMsIHMp
CgkJCX0gZWxzZSB7CgkJCQlidWlsZC5zZWxlY3RzID0gYXBwZW5kKGJ1aWxkLnNl
bGVjdHMsICJgIitzKyJgIikKCQkJfQoJCWRlZmF1bHQ6CgkJCWlmIHIsIG9rIDo9
IHN0ci4oKnJhdyk7IG9rIHsKCQkJCWJ1aWxkLnNlbGVjdHMgPSBhcHBlbmQoYnVp
bGQuc2VsZWN0cywgci5zcWwpCgkJCX0KCQl9Cgl9CglyZXR1cm4gYnVpbGQKfQoK
ZnVuYyAoYnVpbGQgKkJ1aWxkKSBXaGVyZShjb2x1bW4gc3RyaW5nLCBhcmdzIC4u
LmludGVyZmFjZXt9KSAqQnVpbGQgewoJdmFyIG9wZXJhdG9yIHN0cmluZwoJdmFy
IHZhbHVlIGludGVyZmFjZXt9CglpZiBsZW4oYXJncykgPT0gMSB7CgkJb3BlcmF0
b3IgPSAiPSIKCQl2YWx1ZSA9IGFyZ3NbMF0KCX0gZWxzZSBpZiBsZW4oYXJncykg
PiAxIHsKCQlvcGVyYXRvciA9IGFyZ3NbMF0uKHN0cmluZykKCQl2YWx1ZSA9IGFy
Z3NbMV0KCX0KCWJ1aWxkLmFkZFdoZXJlcyh3aGVyZXN7b3BlcmF0b3I6IG9wZXJh
dG9yLCB0cDogIkJhc2ljIiwgdmFsdWU6IHZhbHVlLCBjb2x1bW46IGNvbHVtbiwg
Ym9vbGVhbjogImFuZCJ9KQoJcmV0dXJuIGJ1aWxkCn0KCmZ1bmMgKGJ1aWxkICpC
dWlsZCkgV2hlcmVNYXAoZGF0YXMgbWFwW3N0cmluZ11pbnRlcmZhY2V7fSkgKkJ1
aWxkIHsKCWZvciBrZXksIHZhbCA6PSByYW5nZSBkYXRhcyB7CgkJYnVpbGQuV2hl
cmUoa2V5LCB2YWwpCgl9CglyZXR1cm4gYnVpbGQKfQoKZnVuYyAoYnVpbGQgKkJ1
aWxkKSBPcldoZXJlKGNvbHVtbiBzdHJpbmcsIGFyZ3MgLi4uaW50ZXJmYWNle30p
ICpCdWlsZCB7Cgl2YXIgb3BlcmF0b3Igc3RyaW5nCgl2YXIgdmFsdWUgaW50ZXJm
YWNle30KCWlmIGxlbihhcmdzKSA9PSAxIHsKCQlvcGVyYXRvciA9ICI9IgoJCXZh
bHVlID0gYXJnc1swXQoJfSBlbHNlIHsKCQlvcGVyYXRvciA9IGFyZ3NbMF0uKHN0
cmluZykKCQl2YWx1ZSA9IGFyZ3NbMV0KCX0KCWJ1aWxkLmFkZFdoZXJlcyh3aGVy
ZXN7b3BlcmF0b3I6IG9wZXJhdG9yLCB0cDogIkJhc2ljIiwgdmFsdWU6IHZhbHVl
LCBjb2x1bW46IGNvbHVtbiwgYm9vbGVhbjogIm9yIn0pCglyZXR1cm4gYnVpbGQK
fQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBPcldoZXJlQ2IoY2IgV2hlcmVDYikgKkJ1
aWxkIHsKCWJ1aWxkLmFkZFdoZXJlcyh3aGVyZXN7dHA6ICJjYiIsIGJvb2xlYW46
ICJvciIsIHZhbHVlOiBjYn0pCglyZXR1cm4gYnVpbGQKfQoKZnVuYyAoYnVpbGQg
KkJ1aWxkKSBXaGVyZUNiKGNiIFdoZXJlQ2IpICpCdWlsZCB7CglidWlsZC5hZGRX
aGVyZXMod2hlcmVze3RwOiAiY2IiLCBib29sZWFuOiAiYW5kIiwgdmFsdWU6IGNi
fSkKCXJldHVybiBidWlsZAp9CgpmdW5jIChidWlsZCAqQnVpbGQpIFdoZXJlUmF3
KHNxbCBzdHJpbmcsIGJpbmRpbmdzIC4uLmludGVyZmFjZXt9KSAqQnVpbGQgewoJ
aWYgc3FsID09ICIiIHsKCQlyZXR1cm4gYnVpbGQKCX0KCWJ1aWxkLmFkZFdoZXJl
cyh3aGVyZXN7dHA6ICJyYXciLCBzcWw6IHNxbCwgYm9vbGVhbjogImFuZCIsIHZh
bHVlczogYmluZGluZ3N9KQoJcmV0dXJuIGJ1aWxkCn0KCmZ1bmMgKGJ1aWxkICpC
dWlsZCkgV2hlcmVOdWxsKGNvbHVtbiBzdHJpbmcpICpCdWlsZCB7CglyZXR1cm4g
YnVpbGQuV2hlcmVSYXcoY29sdW1uICsgIiBpcyBOVUxMIikKfQoKZnVuYyAoYnVp
bGQgKkJ1aWxkKSBXaGVyZU5vdE51bGwoY29sdW1uIHN0cmluZykgKkJ1aWxkIHsK
CXJldHVybiBidWlsZC5XaGVyZVJhdyhjb2x1bW4gKyAiIElTIE5PVCBOVUxMIikK
fQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBPcldoZXJlTnVsbChjb2x1bW4gc3RyaW5n
KSAqQnVpbGQgewoJYnVpbGQuYWRkV2hlcmVzKHdoZXJlc3t0cDogInJhdyIsIHNx
bDogY29sdW1uICsgIiBpcyBOVUxMIiwgYm9vbGVhbjogIm9yIn0pCglyZXR1cm4g
YnVpbGQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBPcldoZXJlTm90TnVsbChjb2x1
bW4gc3RyaW5nKSAqQnVpbGQgewoJYnVpbGQuYWRkV2hlcmVzKHdoZXJlc3t0cDog
InJhdyIsIHNxbDogY29sdW1uICsgIiBpcyBOT1QgTlVMTCIsIGJvb2xlYW46ICJv
ciJ9KQoJcmV0dXJuIGJ1aWxkCn0KCmZ1bmMgKGJ1aWxkICpCdWlsZCkgV2hlcmVJ
bihjb2x1bW4gc3RyaW5nLCB2YWx1ZXMgaW50ZXJmYWNle30pICpCdWlsZCB7Cgli
dWlsZC5hZGRXaGVyZXMod2hlcmVze3RwOiAiSW4iLCB2YWx1ZTogdmFsdWVzLCBj
b2x1bW46IGNvbHVtbiwgYm9vbGVhbjogImFuZCJ9KQoJcmV0dXJuIGJ1aWxkCn0K
CmZ1bmMgKGJ1aWxkICpCdWlsZCkgV2hlcmVOb3RJbihjb2x1bW4gc3RyaW5nLCB2
YWx1ZXMgaW50ZXJmYWNle30pICpCdWlsZCB7CglidWlsZC5hZGRXaGVyZXMod2hl
cmVze3RwOiAiTm90SW4iLCB2YWx1ZTogdmFsdWVzLCBjb2x1bW46IGNvbHVtbiwg
Ym9vbGVhbjogImFuZCJ9KQoJcmV0dXJuIGJ1aWxkCn0KCmZ1bmMgKGJ1aWxkICpC
dWlsZCkgam9pbih0YWJsZSBzdHJpbmcsIG9uZSBzdHJpbmcsIG9wIHN0cmluZywg
dHdvIHN0cmluZywgam9pblR5cGUgc3RyaW5nKSAqQnVpbGQgewoJYnVpbGQuam9p
bnMgPSBhcHBlbmQoYnVpbGQuam9pbnMsIGpvaW5ze3RwOiBqb2luVHlwZSwgdGFi
bGU6IHRhYmxlLCBvbmU6IG9uZSwgb3A6IG9wLCB0d286IHR3b30pCglyZXR1cm4g
YnVpbGQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBMZWZ0Sm9pbih0YWJsZSBzdHJp
bmcsIG9uZSBzdHJpbmcsIGFyZ3MgLi4uaW50ZXJmYWNle30pICpCdWlsZCB7Cgl2
YXIgb3AsIHR3byBzdHJpbmcKCWlmIGxlbihhcmdzKSA9PSAwIHsKCQlvcCA9ICI9
IgoJCXR3byA9IG9uZQoJfSBlbHNlIGlmIGxlbihhcmdzKSA9PSAxIHsKCQlvcCA9
ICI9IgoJCXR3byA9IGFyZ3NbMF0uKHN0cmluZykKCX0gZWxzZSB7CgkJb3AgPSBh
cmdzWzBdLihzdHJpbmcpCgkJdHdvID0gYXJnc1sxXS4oc3RyaW5nKQoJfQoJYnVp
bGQuam9pbih0YWJsZSwgb25lLCBvcCwgdHdvLCAibGVmdCIpCglyZXR1cm4gYnVp
bGQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBSaWdodEpvaW4odGFibGUgc3RyaW5n
LCBvbmUgc3RyaW5nLCBhcmdzIC4uLmludGVyZmFjZXt9KSAqQnVpbGQgewoJdmFy
IG9wLCB0d28gc3RyaW5nCglpZiBsZW4oYXJncykgPT0gMCB7CgkJb3AgPSAiPSIK
CQl0d28gPSBvbmUKCX0gZWxzZSBpZiBsZW4oYXJncykgPT0gMSB7CgkJb3AgPSAi
PSIKCQl0d28gPSBhcmdzWzBdLihzdHJpbmcpCgl9IGVsc2UgewoJCW9wID0gYXJn
c1swXS4oc3RyaW5nKQoJCXR3byA9IGFyZ3NbMV0uKHN0cmluZykKCX0KCWJ1aWxk
LmpvaW4odGFibGUsIG9uZSwgb3AsIHR3bywgInJpZ2h0IikKCXJldHVybiBidWls
ZAp9CgpmdW5jIChidWlsZCAqQnVpbGQpIElubmVySm9pbih0YWJsZSBzdHJpbmcs
IG9uZSBzdHJpbmcsIGFyZ3MgLi4uaW50ZXJmYWNle30pICpCdWlsZCB7Cgl2YXIg
b3AsIHR3byBzdHJpbmcKCWlmIGxlbihhcmdzKSA9PSAwIHsKCQlvcCA9ICI9IgoJ
CXR3byA9IG9uZQoJfSBlbHNlIGlmIGxlbihhcmdzKSA9PSAxIHsKCQlvcCA9ICI9
IgoJCXR3byA9IGFyZ3NbMF0uKHN0cmluZykKCX0gZWxzZSB7CgkJb3AgPSBhcmdz
WzBdLihzdHJpbmcpCgkJdHdvID0gYXJnc1sxXS4oc3RyaW5nKQoJfQoJYnVpbGQu
am9pbih0YWJsZSwgb25lLCBvcCwgdHdvLCAiaW5uZXIiKQoJcmV0dXJuIGJ1aWxk
Cn0KCmZ1bmMgKGJ1aWxkICpCdWlsZCkgT3JkZXJCeShjb2x1bW4gc3RyaW5nLCBk
aXJlY3Rpb24gc3RyaW5nKSAqQnVpbGQgewoJYnVpbGQub3JkZXJzID0gYXBwZW5k
KGJ1aWxkLm9yZGVycywgb3JkZXJ7Y29sdW1uOiBjb2x1bW4sIGRpcmVjdGlvbjog
c3RyaW5ncy5Ub0xvd2VyKGRpcmVjdGlvbil9KQoJcmV0dXJuIGJ1aWxkCn0KCmZ1
bmMgKGJ1aWxkICpCdWlsZCkgT3JkZXJEZXNjQnkoY29sdW1uIHN0cmluZykgKkJ1
aWxkIHsKCXJldHVybiBidWlsZC5PcmRlckJ5KGNvbHVtbiwgImRlc2MiKQp9Cgpm
dW5jIChidWlsZCAqQnVpbGQpIE9yZGVyQXNjQnkoY29sdW1uIHN0cmluZykgKkJ1
aWxkIHsKCXJldHVybiBidWlsZC5PcmRlckJ5KGNvbHVtbiwgImFzYyIpCn0KCmZ1
bmMgKGJ1aWxkICpCdWlsZCkgR3JvdXAoY29sdW1uIHN0cmluZykgKkJ1aWxkIHsK
CWJ1aWxkLmdyb3VwID0gY29sdW1uCglyZXR1cm4gYnVpbGQKfQoKZnVuYyAoYnVp
bGQgKkJ1aWxkKSBJbmMoY29sIHN0cmluZywgbnVtIGludCkgKkJ1aWxkIHsKCWJ1
aWxkLm9wZXJhdGUuY3VtdWxhdGVJbmMoY29sLCBudW0pCglyZXR1cm4gYnVpbGQK
fQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBTZXQoY29sIHN0cmluZywgdmFsIGludGVy
ZmFjZXt9KSAqQnVpbGQgewoJYnVpbGQub3BlcmF0ZS5jdW11bGF0ZVNldChjb2ws
IHZhbCkKCXJldHVybiBidWlsZAp9CgpmdW5jIChidWlsZCAqQnVpbGQpIFNraXAo
bGluZXMgaW50KSAqQnVpbGQgewoJYnVpbGQub2Zmc2V0ID0gbGluZXMKCXJldHVy
biBidWlsZAp9CgpmdW5jIChidWlsZCAqQnVpbGQpIE9mZnNldChsaW5lcyBpbnQp
ICpCdWlsZCB7CglyZXR1cm4gYnVpbGQuU2tpcChsaW5lcykKfQoKZnVuYyAoYnVp
bGQgKkJ1aWxkKSBMaW1pdChsaW5lcyBpbnQpICpCdWlsZCB7CglidWlsZC5saW1p
dCA9IGxpbmVzCglyZXR1cm4gYnVpbGQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBU
YWtlKGxpbmVzIGludCkgKkJ1aWxkIHsKCXJldHVybiBidWlsZC5MaW1pdChsaW5l
cykKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBGaXJzdCgpIGJvb2wgewoJZyA6PSBi
dWlsZC5uZXdHb3JtKCkuV2hlcmUoYnVpbGQud2hlcmVRdWVyeSgpLCBidWlsZC5i
aW5kaW5ncy4uLikKCWcgPSBidWlsZC5qb2luUXVlcnkoZykKCWcgPSBidWlsZC5m
aW5kUXVlcnkoZykKCWcgPSBnLkZpcnN0KGJ1aWxkLm1vZGVsKQoJYnVpbGQuZGVh
bEVycm9yKGcpCglpZiBnLlJvd3NBZmZlY3RlZCA+IDAgewoJCXJldHVybiB0cnVl
Cgl9IGVsc2UgewoJCXJldHVybiBmYWxzZQoJfQp9CgpmdW5jIChidWlsZCAqQnVp
bGQpIEdldCgpIGludDY0IHsKCWcgOj0gYnVpbGQubmV3R29ybSgpLldoZXJlKGJ1
aWxkLndoZXJlUXVlcnkoKSwgYnVpbGQuYmluZGluZ3MuLi4pCglnID0gYnVpbGQu
am9pblF1ZXJ5KGcpCglnID0gYnVpbGQuZmluZFF1ZXJ5KGcpCglnID0gZy5GaW5k
KGJ1aWxkLm1vZGVsKQoJYnVpbGQuZGVhbEVycm9yKGcpCglyZXR1cm4gZy5Sb3dz
QWZmZWN0ZWQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBDb3VudCgpIGludDY0IHsK
CXJldHVybiBpbnQ2NChidWlsZC5tYXRoKCIqIiwgIkNvdW50IikpCn0KCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgU3VtKGNvbCBzdHJpbmcpIGZsb2F0NjQgewoJcmV0dXJu
IGJ1aWxkLm1hdGgoY29sLCAiU3VtIikKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBN
YXgoY29sIHN0cmluZykgZmxvYXQ2NCB7CglyZXR1cm4gYnVpbGQubWF0aChjb2ws
ICJNQVgiKQp9CgpmdW5jIChidWlsZCAqQnVpbGQpIG1hdGgoY29sIHN0cmluZywg
bWV0aG9kIHN0cmluZykgZmxvYXQ2NCB7Cgl0eXBlIHNjYW4gc3RydWN0IHsKCQlS
ZXQgZmxvYXQ2NAoJfQoJcyA6PSBuZXcoc2NhbikKCWcgOj0gYnVpbGQubmV3R29y
bSgpLlNlbGVjdChtZXRob2QrIigiK2NvbCsiKSBhcyByZXQiKS5XaGVyZShidWls
ZC53aGVyZVF1ZXJ5KCksIGJ1aWxkLmJpbmRpbmdzLi4uKQoJZyA9IGJ1aWxkLmpv
aW5RdWVyeShnKQoJZyA9IGJ1aWxkLmZpbmRRdWVyeShnKQoJZyA9IGcuU2Nhbihz
KQoJYnVpbGQuZGVhbEVycm9yKGcpCglyZXR1cm4gcy5SZXQKfQoKZnVuYyAoYnVp
bGQgKkJ1aWxkKSBVcGRhdGUoaCBpdXRpbHMuSCkgaW50NjQgewoJZyA6PSBidWls
ZC5uZXdHb3JtKCkuV2hlcmUoYnVpbGQud2hlcmVRdWVyeSgpLCBidWlsZC5iaW5k
aW5ncy4uLikuVXBkYXRlcyhoKQoJYnVpbGQuZGVhbEVycm9yKGcpCglyZXR1cm4g
Zy5Sb3dzQWZmZWN0ZWQKfQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBTYXZlKCkgewoJ
ZyA6PSBidWlsZC5uZXdHb3JtKCkuU2F2ZShidWlsZC5tb2RlbCkKCWJ1aWxkLmRl
YWxFcnJvcihnKQp9CgovLyDmlK/mjIHliIbooagKZnVuYyAoYnVpbGQgKkJ1aWxk
KSBJbnNlcnQoYXJndSBpbnRlcmZhY2V7fSkgewoJaCwgb2sgOj0gYXJndS4oW11t
YXBbc3RyaW5nXWludGVyZmFjZXt9KQoJaWYgIW9rIHsKCQl4LCBvayA6PSBhcmd1
LihtYXBbc3RyaW5nXWludGVyZmFjZXt9KQoJCWlmICFvayB7CgkJCXBhbmljKCLm
lbDmja7nsbvlnovplJnor68iKQoJCX0KCQloID0gW11tYXBbc3RyaW5nXWludGVy
ZmFjZXt9e3h9Cgl9CglpZiBsZW4oaCkgPT0gMCB7CgkJcmV0dXJuCgl9CgllbGUg
Oj0gaFswXQoJa2V5cyA6PSBtYWtlKFtdc3RyaW5nLCBsZW4oZWxlKSkKCWkgOj0g
MAoJZm9yIGsgOj0gcmFuZ2UgZWxlIHsKCQlrZXlzW2ldID0gawoJCWkrKwoJfQoJ
aWYgYnVpbGQudGFibGUgIT0gIiIgewoJCS8vIOW3sue7j+aMh+WumuihqOWQjeS6
hgoJCWJpbmRpbmdzIDo9IGZvcm1hdEJpbmRpbmdzKGgsIGtleXMpCgkJYnVpbGQu
RXhlYyhidWlsZC5pbnNlcnRTcWwoa2V5cywgbGVuKGgpKSwgYmluZGluZ3MuLi4p
CgkJcmV0dXJuCgl9CglzcGxpdCwgdGFiIDo9IGJ1aWxkLmlzU3BsaXQoKQoJc29y
dC5TdHJpbmdzKGtleXMpCglpZiB0YWIgIT0gbmlsIHsKCQkvLyDpnZ7liIbooagK
CQliaW5kaW5ncyA6PSBmb3JtYXRCaW5kaW5ncyhoLCBrZXlzKQoJCWJ1aWxkLkV4
ZWMoYnVpbGQuaW5zZXJ0U3FsKGtleXMsIGxlbihoKSksIGJpbmRpbmdzLi4uKQoJ
CXJldHVybgoJfQoJc3BsaXQuQmFzZU5hbWUoKQoJc3BsaXQubG9hZCgpCgoJLy8g
5YiG5ouj5pWw5o2uCglzcGxpdERhdGEgOj0gbWFrZShtYXBbaW50NjRdW11tYXBb
c3RyaW5nXWludGVyZmFjZXt9KQoJZm9yIF8sIGl0ZW0gOj0gcmFuZ2UgaCB7CgkJ
diwgb2sgOj0gaXRlbVtzcGxpdC5nZXRTcGxpdENvbCgpXQoJCWlmICFvayB7CgkJ
CXBhbmljKCLmj5LlhaXmlbDmja7nvLrlsJHliIbooajlgLwiKQoJCX0KCQlpZHgg
Oj0gc3BsaXQuc3BsaXRJZHgoaXV0aWxzLkFzSW50NjQodikpCgkJXywgb2sgPSBz
cGxpdERhdGFbaWR4XQoJCWlmICFvayB7CgkJCXNwbGl0RGF0YVtpZHhdID0gbWFr
ZShbXW1hcFtzdHJpbmddaW50ZXJmYWNle30sIDApCgkJfQoJCXNwbGl0RGF0YVtp
ZHhdID0gYXBwZW5kKHNwbGl0RGF0YVtpZHhdLCBpdGVtKQoJfQoJZm9yIGlkeCwg
bGlzdCA6PSByYW5nZSBzcGxpdERhdGEgewoJCWJpbmRpbmdzIDo9IGZvcm1hdEJp
bmRpbmdzKGxpc3QsIGtleXMpCgkJYiA6PSBOZXdCdWlsZChidWlsZC5tb2RlbCku
U3BsaXRCeShpZHgpCgkJYi5FeGVjKGIuaW5zZXJ0U3FsKGtleXMsIGxlbihsaXN0
KSksIGJpbmRpbmdzLi4uKQoJfQp9CgpmdW5jIChidWlsZCAqQnVpbGQpIERlbGV0
ZSgpIGludDY0IHsKCXNvZnRlciA6PSBidWlsZC5zb2Z0RGVsZXRlZCgpCglpZiBz
b2Z0ZXIgPT0gbmlsIHsKCQlnIDo9IGJ1aWxkLm5ld0dvcm0oKS5XaGVyZShidWls
ZC53aGVyZVF1ZXJ5KCksIGJ1aWxkLmJpbmRpbmdzLi4uKS5EZWxldGUoJmJ1aWxk
Lm1vZGVsKQoJCWJ1aWxkLmRlYWxFcnJvcihnKQoJCXJldHVybiBnLlJvd3NBZmZl
Y3RlZAoJfQoJY29sLCBfLCBkZWxWYWwgOj0gc29mdGVyLlNvZnREZWxldGVkKCkK
CXJldHVybiBidWlsZC5VcGRhdGUoaXV0aWxzLkh7Y29sOiBkZWxWYWx9KQp9Cgov
LyDlsIbliJrliJrnp6/ntK/nmoTmk43kvZzvvIzkuIDmrKHmgKfmiafooYzliLDm
lbDmja7lupMKZnVuYyAoYnVpbGQgKkJ1aWxkKSBEb25lT3BlcmF0ZSgpIGludDY0
IHsKCW9wIDo9IGJ1aWxkLm9wZXJhdGUucmF3KCkKCWlmIGxlbihvcCkgPT0gMCB7
CgkJcmV0dXJuIDAKCX0KCXJldHVybiBidWlsZC5VcGRhdGUob3ApCn0KCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgSW5jcmVtZW50KGNvbHVtbiBzdHJpbmcsIGFtb3VudCBp
bnQpIGludDY0IHsKCWcgOj0gYnVpbGQubmV3R29ybSgpLldoZXJlKGJ1aWxkLndo
ZXJlUXVlcnkoKSwgYnVpbGQuYmluZGluZ3MuLi4pLgoJCVVwZGF0ZShjb2x1bW4s
IGdvcm0uRXhwcihjb2x1bW4rIiArID8iLCBhbW91bnQpKQoJYnVpbGQuZGVhbEVy
cm9yKGcpCglyZXR1cm4gZy5Sb3dzQWZmZWN0ZWQKfQoKZnVuYyAoYnVpbGQgKkJ1
aWxkKSBSYXcoc3FsIHN0cmluZywgdmFsdWVzIC4uLmludGVyZmFjZXt9KSAqZ29y
bS5EQiB7CglnIDo9IGJ1aWxkLm5ld0NvbigpLk1vZGVsKGJ1aWxkLm1vZGVsKS5S
YXcoc3FsLCB2YWx1ZXMuLi4pLlNjYW4oYnVpbGQubW9kZWwpCglidWlsZC5kZWFs
RXJyb3IoZykKCXJldHVybiBnCn0KCmZ1bmMgKGJ1aWxkICpCdWlsZCkgRXhlYyhz
cWwgc3RyaW5nLCB2YWx1ZXMgLi4uaW50ZXJmYWNle30pICpnb3JtLkRCIHsKCWcg
Oj0gYnVpbGQubmV3Q29uKCkuTW9kZWwoYnVpbGQubW9kZWwpLkV4ZWMoc3FsLCB2
YWx1ZXMuLi4pCglidWlsZC5kZWFsRXJyb3IoZykKCXJldHVybiBnCn0KCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgRm9yVXBkYXRlKCkgKkJ1aWxkIHsKCWJ1aWxkLnNldHNb
Imdvcm06cXVlcnlfb3B0aW9uIl0gPSAiRk9SIFVQREFURSIKCXJldHVybiBidWls
ZAp9CgovLyDlvZPpnIDopoHliIbooajvvIzkvYbmmK/lnKh3aGVyZeS4reW5tuay
oeacieiwg+eUqOWIhuihqGtleeeahOafpeivou+8jOWImemcgOimgeaJi+WKqOaJ
p+ihjOivpeWHveaVsOWIpOaWreWIhuihqApmdW5jIChidWlsZCAqQnVpbGQpIFNw
bGl0QnkodmFsIGludDY0KSAqQnVpbGQgewoJYnVpbGQuaXNTZXRTcGxpdFZhbCA9
IHRydWUKCWJ1aWxkLnNwbGl0VmFsID0gdmFsCglyZXR1cm4gYnVpbGQKfQoKLy8g
5oyH5a6a6KGo5ZCNCmZ1bmMgKGJ1aWxkICpCdWlsZCkgVGFibGVOYW1lKG5hbWUg
c3RyaW5nKSAqQnVpbGQgewoJYnVpbGQudGFibGUgPSBuYW1lCglyZXR1cm4gYnVp
bGQKfQoKLy8g5oyH5a6abW9kZWznsbsKZnVuYyAoYnVpbGQgKkJ1aWxkKSBNb2Rl
bFR5cGUodiBpbnRlcmZhY2V7fSkgKkJ1aWxkIHsKCWJ1aWxkLm1vZGVsID0gdgoJ
cmV0dXJuIGJ1aWxkCn0KCi8vIOiwg+eUqOatpOWHveaVsOWQju+8jOaVsOaNruW6
k+eahOaJgOaciemUmeivr+mDveS4jeWGjeaKm+WHuu+8jOmcgOimgeS9v+eUqOiA
heS4u+WKqOiwg+eUqEVycm9y6I635Y+W6ZSZ6K+v5aSE55CGCmZ1bmMgKGJ1aWxk
ICpCdWlsZCkgRGlzYWJsZVBhbmljKCkgKkJ1aWxkIHsKCWJ1aWxkLm5vdFBhbmlj
ID0gdHJ1ZQoJcmV0dXJuIGJ1aWxkCn0KCi8vIOiOt+WPlnNxbOaJp+ihjOeahOmU
meivrywg6ZyA6KaB5YWI6LCD55SoRGlzYWJsZVBhbmljCmZ1bmMgKGJ1aWxkICpC
dWlsZCkgRXJyb3IoKSBlcnJvciB7CglyZXR1cm4gYnVpbGQuZXJyCn0KCi8vIOeU
qOe7meWumueahOWAvOaehOW7uuWHunNxbCwg5ZKM57uR5a6a5YC8CmZ1bmMgKGJ1
aWxkICpCdWlsZCkgQnVpbGRJbnNlcnRTcWwoYXJndSBpbnRlcmZhY2V7fSkgKHN0
cmluZywgW11pbnRlcmZhY2V7fSkgewoJaCwgb2sgOj0gYXJndS4oW11tYXBbc3Ry
aW5nXWludGVyZmFjZXt9KQoJaWYgIW9rIHsKCQl4LCBvayA6PSBhcmd1LihtYXBb
c3RyaW5nXWludGVyZmFjZXt9KQoJCWlmICFvayB7CgkJCXBhbmljKCLmlbDmja7n
sbvlnovplJnor68iKQoJCX0KCQloID0gW11tYXBbc3RyaW5nXWludGVyZmFjZXt9
e3h9Cgl9CglpZiBsZW4oaCkgPT0gMCB7CgkJcmV0dXJuICIiLCBbXWludGVyZmFj
ZXt9e30KCX0KCWVsZSA6PSBoWzBdCglrZXlzIDo9IG1ha2UoW11zdHJpbmcsIGxl
bihlbGUpKQoJaSA6PSAwCglmb3IgayA6PSByYW5nZSBlbGUgewoJCWtleXNbaV0g
PSBrCgkJaSsrCgl9CgliaW5kaW5ncyA6PSBmb3JtYXRCaW5kaW5ncyhoLCBrZXlz
KQoJc3FsIDo9IGJ1aWxkLmluc2VydFNxbChrZXlzLCBsZW4oaCkpCglyZXR1cm4g
c3FsLCBiaW5kaW5ncwp9CgovLyDkvb/nlKjphY3nva7kuK1Jc1dyaXRl55qE6L+e
5o6l77yM6Iul5LiN5a2Y5Zyo77yM5YiZ5L2/55SocHJveHnov57mjqUKZnVuYyAo
YnVpbGQgKkJ1aWxkKSBPbldyaXRlKCkgKkJ1aWxkIHsKCWJ1aWxkLmlzT25Xcml0
ZSA9IHRydWUKCXJldHVybiBidWlsZAp9CgovLyDojrflj5bkuIDkuKrpgJrnlKjm
lbDmja7lupPlj6Xmn4QKZnVuYyAoYnVpbGQgKkJ1aWxkKSBEQigpICpzcWwuREIg
ewoJcmV0dXJuIGJ1aWxkLm5ld0NvbigpLkRCKCkKfQoKLy8gUHJpdmF0ZSBmdW5j
IOS7peS4i+S4uuengeacieaOpeWPowoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBhZGRX
aGVyZXMod2hlIHdoZXJlcykgewoJYnVpbGQud2hlcmVzID0gYXBwZW5kKGJ1aWxk
LndoZXJlcywgd2hlKQp9CgpmdW5jIChidWlsZCAqQnVpbGQpIGFkZEJpbmRpbmco
dmFsdWVzIGludGVyZmFjZXt9KSB7CglidWlsZC5iaW5kaW5ncyA9IGFwcGVuZChi
dWlsZC5iaW5kaW5ncywgdmFsdWVzKQp9CgpmdW5jIChidWlsZCAqQnVpbGQpIG5l
d0NvbigpICpnb3JtLkRCIHsKCXZhciBnICpnb3JtLkRCCglnaWQgOj0gZ2V0U2xv
dygpCgljb24gOj0gYnVpbGQuZ2V0Q29uTmFtZSgpCgloIDo9IGdldEhhbmRsZShn
aWQsIGNvbikudHJhbnMKCWlmIGggPT0gbmlsIHsKCQlpZiBidWlsZC5pc09uV3Jp
dGUgewoJCQlpZiBfLCBvayA6PSBtb2RlbHNbY29uXVtXcml0ZV07ICFvayB7CgkJ
CQlnID0gbW9kZWxzW2Nvbl1bV3JpdGVdCgkJCX0gZWxzZSB7CgkJCQlnID0gbW9k
ZWxzW2Nvbl1bUHJveHldCgkJCX0KCQl9IGVsc2UgewoJCQlnID0gbW9kZWxzW2Nv
bl1bUHJveHldCgkJfQoJfSBlbHNlIHsKCQlnID0gaC5kYgoJfQoJcmV0dXJuIGcK
fQoKZnVuYyAoYnVpbGQgKkJ1aWxkKSBuZXdHb3JtKCkgKmdvcm0uREIgewoJZyA6
PSBidWlsZC5uZXdDb24oKS5Nb2RlbChidWlsZC5tb2RlbCkuVGFibGUoYnVpbGQu
dGFibGVOYW1lKCkpLlVuc2NvcGVkKCkKCWZvciBrZXksIHNldCA6PSByYW5nZSBi
dWlsZC5zZXRzIHsKCQlnID0gZy5TZXQoa2V5LCBzZXQpCgl9CgoJLy8g5byA5Y+R
546v5aKD5rOo5YaMIEFuYWx5emVyIOaPkuS7tgoJLy8gaWYgY29uZmlnLkRldmVs
b3BtZW50IHsKCS8vICAgIGcuQ2FsbGJhY2soKS5RdWVyeSgpLlJlZ2lzdGVyKCJB
bmFseXplciIsIEFuYWx5emVyQ2FsbGJhY2spCgkvLyAgICAvLyBnLkNhbGxiYWNr
KCkuUXVlcnkoKS5CZWZvcmUoImdvcm06cXVlcnkiKS5SZWdpc3RlcigiQW5hbHl6
ZXIiLCBBbmFseXplckNhbGxiYWNrKQoJLy8gfQoKCXJldHVybiBnCn0KCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgd2hlcmVTcWwod2hlIHdoZXJlcykgKHN0cmluZywgaW50
ZXJmYWNle30sIGJvb2wpIHsKCXN3aXRjaCB3aGUudHAgewoJY2FzZSAiQmFzaWMi
OgoJCXJldHVybiAiYCIgKyB3aGUuY29sdW1uICsgImAgIiArIHdoZS5vcGVyYXRv
ciArICIgPyIsIHdoZS52YWx1ZSwgZmFsc2UKCWNhc2UgInJhdyI6CgkJcmV0dXJu
IHdoZS5zcWwsIHdoZS52YWx1ZXMsIHRydWUKCWNhc2UgIkluIjoKCQlyZXR1cm4g
ImAiICsgd2hlLmNvbHVtbiArICJgIGluICg/KSIsIHdoZS52YWx1ZSwgZmFsc2UK
CWNhc2UgIk5vdEluIjoKCQlyZXR1cm4gImAiICsgd2hlLmNvbHVtbiArICJgIG5v
dCBpbiAoPykiLCB3aGUudmFsdWUsIGZhbHNlCgljYXNlICJjYiI6CgkJYiA6PSBO
ZXdCdWlsZChidWlsZC5tb2RlbCkKCQl3aGUudmFsdWUuKFdoZXJlQ2IpKGIpCgkJ
cmV0dXJuIGIud2hlcmVRdWVyeSgpLCBiLmJpbmRpbmdzLCB0cnVlCgl9CglwYW5p
Yygi5pON5L2c57G75Z6L6ZSZ6K+vIikKfQoKLy8g5LiN5aSq5aW955SoCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgam9pblF1ZXJ5KGcgKmdvcm0uREIpICpnb3JtLkRCIHsK
CW5hbWUgOj0gYnVpbGQudGFibGVOYW1lKCkKCWZvciBfLCBqb2luIDo9IHJhbmdl
IGJ1aWxkLmpvaW5zIHsKCQl2YXIgbywgdCwgcXVlcnkgc3RyaW5nCgkJaWYgIXN0
cmluZ3MuQ29udGFpbnMoam9pbi5vbmUsICIuIikgJiYgIXN0cmluZ3MuQ29udGFp
bnMoc3RyaW5ncy5Ub0xvd2VyKG5hbWUpLCAiIGFzICIpIHsKCQkJbyA9IG5hbWUg
KyAiLiIgKyBqb2luLm9uZQoJCX0gZWxzZSB7CgkJCW8gPSBqb2luLm9uZQoJCX0K
CQlpZiAhc3RyaW5ncy5Db250YWlucyhqb2luLnR3bywgIi4iKSAmJiAhc3RyaW5n
cy5Db250YWlucyhzdHJpbmdzLlRvTG93ZXIoam9pbi50YWJsZSksICIgYXMgIikg
ewoJCQl0ID0gam9pbi50YWJsZSArICIuIiArIGpvaW4udHdvCgkJfSBlbHNlIHsK
CQkJdCA9IGpvaW4udHdvCgkJfQoJCXF1ZXJ5ICs9IGpvaW4udHAgKyAiIGpvaW4g
IiArIGpvaW4udGFibGUgKyAiIG9uICIgKyBvICsgIiAiICsgam9pbi5vcCArICIg
IiArIHQKCQlnID0gZy5Kb2lucyhxdWVyeSkKCX0KCXJldHVybiBnCn0KCmZ1bmMg
KGJ1aWxkICpCdWlsZCkgd2hlcmVRdWVyeSgpIHN0cmluZyB7CglxdWVyeSA6PSAi
IgoKCS8vIOa3u+WKoOi9r+WIoOmZpOeahOafpeivouadoeS7tgoJaWYgIWJ1aWxk
LnVuc2NvcGVkICYmIGJ1aWxkLnNvZnREZWxldGVkKCkgIT0gbmlsIHsKCQlzb2Z0
ZXIgOj0gYnVpbGQuc29mdERlbGV0ZWQoKQoJCWNvbCwgbGl2ZVZhbCwgXyA6PSBz
b2Z0ZXIuU29mdERlbGV0ZWQoKQoJCWlmIGxpdmVWYWwgPT0gbmlsIHsKCQkJYnVp
bGQuV2hlcmVOdWxsKGNvbCkKCQl9IGVsc2UgewoJCQlidWlsZC5XaGVyZShjb2ws
IGxpdmVWYWwpCgkJfQoKCX0KCglmb3IgaWR4LCB3aGUgOj0gcmFuZ2UgYnVpbGQu
d2hlcmVzIHsKCQlzcWwsIHZhbCwgaXNNdWx0aUJpbmQgOj0gYnVpbGQud2hlcmVT
cWwod2hlKQoJCWJvb2xlYW4gOj0gd2hlLmJvb2xlYW4KCQlpZiBpZHggPT0gMCB7
CgkJCWJvb2xlYW4gPSAiIgoJCX0KCQlxdWVyeSArPSBib29sZWFuICsgIiAoIiAr
IHNxbCArICIpICIKCQlpZiB2YWwgIT0gbmlsIHsKCQkJaWYgaXNNdWx0aUJpbmQg
ewoJCQkJZm9yIF8sIHYgOj0gcmFuZ2UgdmFsLihbXWludGVyZmFjZXt9KSB7CgkJ
CQkJYnVpbGQuYWRkQmluZGluZyh2KQoJCQkJfQoJCQl9IGVsc2UgewoJCQkJYnVp
bGQuYWRkQmluZGluZyh2YWwpCgkJCX0KCQl9Cgl9CglyZXR1cm4gcXVlcnkKfQoK
ZnVuYyAoYnVpbGQgKkJ1aWxkKSBmaW5kUXVlcnkoZyAqZ29ybS5EQikgKmdvcm0u
REIgewoJaWYgbGVuKGJ1aWxkLnNlbGVjdHMpICE9IDAgewoJCWcgPSBnLlNlbGVj
dChidWlsZC5zZWxlY3RzKQoJfQoJaWYgYnVpbGQubGltaXQgIT0gMCB7CgkJZyA9
IGcuTGltaXQoYnVpbGQubGltaXQpCgl9CgoJaWYgYnVpbGQub2Zmc2V0ICE9IDAg
ewoJCWcgPSBnLk9mZnNldChidWlsZC5vZmZzZXQpCgl9CglpZiBidWlsZC5ncm91
cCAhPSAiIiB7CgkJZyA9IGcuR3JvdXAoYnVpbGQuZ3JvdXApCgl9Cglmb3IgXywg
b3JkZXIgOj0gcmFuZ2UgYnVpbGQub3JkZXJzIHsKCQlnID0gZy5PcmRlcihvcmRl
ci5jb2x1bW4gKyAiICIgKyBvcmRlci5kaXJlY3Rpb24pCgl9CglyZXR1cm4gZwp9
CgpmdW5jIChidWlsZCAqQnVpbGQpIGluc2VydFNxbChrZXlzIFtdc3RyaW5nLCBs
aW5lcyBpbnQpIHN0cmluZyB7CglrZXlzU3FsIDo9IHN0cmluZ3MuSm9pbihrZXlz
LCAiYCxgIikKCXZhbHVlQ2hhciA6PSBtYWtlKFtdc3RyaW5nLCBsZW4oa2V5cykp
Cglmb3IgaSA6PSAwOyBpIDwgbGVuKGtleXMpOyBpKysgewoJCXZhbHVlQ2hhcltp
XSA9ICI/IgoJfQoJdmFsdWVTcWwgOj0gc3RyaW5ncy5Kb2luKHZhbHVlQ2hhciwg
IiwiKQoJdmFsdWVzU3RyIDo9IG1ha2UoW11zdHJpbmcsIGxpbmVzKQoJZm9yIGsg
Oj0gMDsgayA8IGxpbmVzOyBrKysgewoJCXZhbHVlc1N0cltrXSA9ICIoIiArIHZh
bHVlU3FsICsgIikiCgl9Cgl2YWx1ZXNTcWwgOj0gc3RyaW5ncy5Kb2luKHZhbHVl
c1N0ciwgIiwiKQoJcmV0dXJuICJJTlNFUlQgIiArIGJ1aWxkLnRhYmxlTmFtZSgp
ICsgIiAoYCIgKyBrZXlzU3FsICsgImApIFZBTFVFUyAiICsgdmFsdWVzU3FsCn0K
Ci8vIOagueaNrm1vZGVs5Yik5pat5piv5ZCm5piv5YiG6KGoLCDkuKTkuKrov5Tl
m57lgLzkuK3vvIzlj6rmnInkuIDkuKrmnInlgLwKZnVuYyAoYnVpbGQgKkJ1aWxk
KSBpc1NwbGl0KCkgKHNwbGl0SW50ZXJmYWNlLCB0YWJsZXIpIHsKCXJlZmxlY3RU
eXBlIDo9IHJlZmxlY3QuVmFsdWVPZihidWlsZC5tb2RlbCkuVHlwZSgpCglmb3Ig
cmVmbGVjdFR5cGUuS2luZCgpID09IHJlZmxlY3QuU2xpY2UgfHwgcmVmbGVjdFR5
cGUuS2luZCgpID09IHJlZmxlY3QuUHRyIHsKCQlyZWZsZWN0VHlwZSA9IHJlZmxl
Y3RUeXBlLkVsZW0oKQoJfQoJdGFiLCBvayA6PSByZWZsZWN0Lk5ldyhyZWZsZWN0
VHlwZSkuSW50ZXJmYWNlKCkuKHNwbGl0SW50ZXJmYWNlKQoJaWYgIW9rIHsKCQly
ZXR1cm4gbmlsLCByZWZsZWN0Lk5ldyhyZWZsZWN0VHlwZSkuSW50ZXJmYWNlKCku
KHRhYmxlcikKCX0KCXJldHVybiB0YWIsIG5pbAp9CgovLyB0YWJsZU5hbWUg6L+U
5Zue6KGo5ZCN77yM5aaC5p6c5p+l6K+i55qE5piv5YiG6KGo77yM5L2G5piv5Y20
5Y+I5rKh5pyJ6K6+572u5YiG6KGoa2V555qE5YC877yM5YiZ5oqb5Ye65byC5bi4
CmZ1bmMgKGJ1aWxkICpCdWlsZCkgdGFibGVOYW1lKCkgc3RyaW5nIHsKCWlmIGJ1
aWxkLnRhYmxlICE9ICIiIHsKCQlyZXR1cm4gYnVpbGQudGFibGUKCX0KCXRhYiwg
b3RoZXIgOj0gYnVpbGQuaXNTcGxpdCgpCglpZiBvdGhlciAhPSBuaWwgewoJCXJl
dHVybiBvdGhlci5UYWJsZU5hbWUoKQoJfQoJdGFiLkJhc2VOYW1lKCkKCXRhYi5s
b2FkKCkKCgkvLyDlpoLmnpzlpJbpg6jlt7Lnu4/miYvliqjorr7nva7ov4fliIbo
oajlgLzvvIzliJnnm7TmjqXorr7nva4KCWlmIGJ1aWxkLmlzU2V0U3BsaXRWYWwg
ewoJCXRhYi5zZXRTcGxpdFZhbChidWlsZC5zcGxpdFZhbCkKCQlyZXR1cm4gdGFi
LlRhYmxlTmFtZSgpCgl9Cglmb3IgXywgd2hlIDo9IHJhbmdlIGJ1aWxkLndoZXJl
cyB7CgkJaWYgd2hlLnRwICE9ICJCYXNpYyIgewoJCQljb250aW51ZQoJCX0KCQlp
ZiB3aGUub3BlcmF0b3IgIT0gIj0iIHsKCQkJY29udGludWUKCQl9CgkJaWYgd2hl
LmNvbHVtbiAhPSB0YWIuZ2V0U3BsaXRDb2woKSB7CgkJCWNvbnRpbnVlCgkJfQoJ
CXZhciB2IGludDY0CgkJc3dpdGNoIHdoZS52YWx1ZS4odHlwZSkgewoJCWNhc2Ug
aW50OgoJCQl2ID0gaW50NjQod2hlLnZhbHVlLihpbnQpKQoJCWNhc2UgaW50NjQ6
CgkJCXYgPSB3aGUudmFsdWUuKGludDY0KQoJCWNhc2UgZmxvYXQ2NDoKCQkJdiA9
IGludDY0KHdoZS52YWx1ZS4oZmxvYXQ2NCkpCgkJY2FzZSBzdHJpbmc6CgkJCWlk
LCBlcnIgOj0gc3RyY29udi5QYXJzZUludCh3aGUudmFsdWUuKHN0cmluZyksIDEw
LCA2NCkKCQkJaWYgZXJyICE9IG5pbCB7CgkJCQlwYW5pYygi5YiG6KGo5YC857G7
5Z6L6ZSZ6K+vLCIgKyBlcnIuRXJyb3IoKSkKCQkJfQoJCQl2ID0gaWQKCQlkZWZh
dWx0OgoJCQlwYW5pYygi5YiG6KGo5YC857G75Z6L6ZSZ6K+vIikKCQl9CgkJdGFi
LnNldFNwbGl0VmFsKHYpCgkJcmV0dXJuIHRhYi5UYWJsZU5hbWUoKQoJfQoJcGFu
aWMoIuWIhuihqOafpeivouacquiuvue9ruWIhuihqOWAvCIpCn0KCmZ1bmMgKGJ1
aWxkICpCdWlsZCkgZGVhbEVycm9yKGcgKmdvcm0uREIpIHsKCWlmIGcuRXJyb3Ig
PT0gbmlsIHsKCQlyZXR1cm4KCX0KCWlmIGcuRXJyb3IuRXJyb3IoKSA9PSAicmVj
b3JkIG5vdCBmb3VuZCIgewoJCS8vIOW/veeVpeayoeacieaJvuWIsOaVsOaNruea
hOmUmeivrwoJCXJldHVybgoJfQoJaWYgIWJ1aWxkLm5vdFBhbmljIHsKCQlwYW5p
YyhnLkVycm9yLkVycm9yKCkpCgl9IGVsc2UgewoJCWJ1aWxkLmVyciA9IGcuRXJy
b3IKCX0KfQoKLy8g5bel5YW35Ye95pWwCmZ1bmMgZm9ybWF0QmluZGluZ3MobGlz
dCBbXW1hcFtzdHJpbmddaW50ZXJmYWNle30sIGtleXMgW11zdHJpbmcpIFtdaW50
ZXJmYWNle30gewoJYmluZGluZ3MgOj0gbWFrZShbXWludGVyZmFjZXt9LCBsZW4o
bGlzdCkqbGVuKGtleXMpKQoJbSA6PSAwCglmb3IgXywgaXRlbSA6PSByYW5nZSBs
aXN0IHsKCQlmb3IgXywgayA6PSByYW5nZSBrZXlzIHsKCQkJdiwgb2sgOj0gaXRl
bVtrXQoJCQlpZiAhb2sgewoJCQkJcGFuaWMoIuaPkuWFpeaVsOaNrue8uuWwkWtl
eToiICsgaykKCQkJfQoJCQliaW5kaW5nc1ttXSA9IHYKCQkJbSsrCgkJfQoJfQoJ
cmV0dXJuIGJpbmRpbmdzCn0KCi8vIOiOt+WPlui/nuaOpeWQjQpmdW5jIChidWls
ZCAqQnVpbGQpIGdldENvbk5hbWUoKSBzdHJpbmcgewoJdHlwZSBjb25JbnRlcmZh
Y2UgaW50ZXJmYWNlIHsKCQlDb25uZWN0KCkgc3RyaW5nCgl9CgoJaWYgYnVpbGQu
Y29uICE9ICIiIHsKCQlyZXR1cm4gYnVpbGQuY29uCgl9CgoJaWYgYnVpbGQubW9k
ZWwgPT0gbmlsIHsKCQlyZXR1cm4gRGVmYXVsdENvbgoJfQoJcmVmbGVjdFR5cGUg
Oj0gcmVmbGVjdC5WYWx1ZU9mKGJ1aWxkLm1vZGVsKS5UeXBlKCkKCWZvciByZWZs
ZWN0VHlwZS5LaW5kKCkgPT0gcmVmbGVjdC5TbGljZSB8fCByZWZsZWN0VHlwZS5L
aW5kKCkgPT0gcmVmbGVjdC5QdHIgewoJCXJlZmxlY3RUeXBlID0gcmVmbGVjdFR5
cGUuRWxlbSgpCgl9CglfLCBvayA6PSByZWZsZWN0Lk5ldyhyZWZsZWN0VHlwZSku
SW50ZXJmYWNlKCkuKGNvbkludGVyZmFjZSkKCWlmIG9rIHsKCQlyZXR1cm4gcmVm
bGVjdC5OZXcocmVmbGVjdFR5cGUpLkludGVyZmFjZSgpLihjb25JbnRlcmZhY2Up
LkNvbm5lY3QoKQoJfQoJcmV0dXJuIERlZmF1bHRDb24KfQoKdHlwZSBzb2Z0ZXIg
aW50ZXJmYWNlIHsKCVNvZnREZWxldGVkKCkgKHN0cmluZywgaW50ZXJmYWNle30s
IGludGVyZmFjZXt9KQp9CgovLyDojrflj5bova/liKDpmaTnmoTorr7nva4KZnVu
YyAoYnVpbGQgKkJ1aWxkKSBzb2Z0RGVsZXRlZCgpIHNvZnRlciB7CglpZiAhcmVm
bGVjdC5WYWx1ZU9mKGJ1aWxkLm1vZGVsKS5Jc1ZhbGlkKCkgewoJCXJldHVybiBu
aWwKCX0KCXJlZmxlY3RUeXBlIDo9IHJlZmxlY3QuVmFsdWVPZihidWlsZC5tb2Rl
bCkuVHlwZSgpCglmb3IgcmVmbGVjdFR5cGUuS2luZCgpID09IHJlZmxlY3QuU2xp
Y2UgfHwgcmVmbGVjdFR5cGUuS2luZCgpID09IHJlZmxlY3QuUHRyIHsKCQlyZWZs
ZWN0VHlwZSA9IHJlZmxlY3RUeXBlLkVsZW0oKQoJfQoJcywgb2sgOj0gcmVmbGVj
dC5OZXcocmVmbGVjdFR5cGUpLkludGVyZmFjZSgpLihzb2Z0ZXIpCglpZiBvayB7
CgkJcmV0dXJuIHMKCX0KCXJldHVybiBuaWwKfQo=
`



	r, _ := base64.StdEncoding.DecodeString(b)
	return string(r)
}




