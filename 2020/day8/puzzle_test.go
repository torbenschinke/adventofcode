package day8

import (
	"fmt"
	"testing"
)

const (
	set0 = "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	set1 = "acc +9\nacc -2\nacc -12\nacc +33\njmp +301\nnop +508\njmp +216\nacc +27\nacc +35\nacc +43\nacc +31\njmp +309\nacc +18\nacc -19\nacc +7\njmp +44\nacc -13\nacc -17\nacc +31\njmp +311\nnop +612\njmp +143\nacc +22\nnop +85\njmp +458\nacc -3\njmp +13\nacc -19\nacc +27\nacc +12\njmp +483\nacc +40\nacc +6\njmp +128\njmp +10\nacc +0\nacc -3\nacc -2\njmp -11\nacc +43\nacc -12\njmp +158\nacc +0\njmp +240\njmp +1\nacc +5\nacc +15\njmp +187\nnop +563\njmp +51\nacc -16\njmp +158\njmp +322\nacc +47\nnop -1\njmp +299\nacc +26\nacc +25\njmp +232\njmp -9\nacc +15\njmp +54\njmp +558\nacc +7\nacc -7\njmp +399\nnop +447\njmp +71\nacc +26\nacc +46\njmp +145\nacc +38\nacc +30\nacc +21\njmp +263\nacc +10\njmp +168\nacc +22\nnop +561\njmp -26\njmp +1\nacc -7\njmp -5\nacc +28\nacc -6\njmp +370\njmp +94\nacc +50\nacc +42\nacc -9\nacc +30\njmp +70\nacc +29\njmp +166\nacc -5\nacc -18\nnop +84\nacc +2\njmp +366\njmp -40\nacc -4\nacc -15\nacc -1\njmp +169\njmp +1\nacc -4\nacc +0\njmp -45\nnop -21\nnop +241\nacc -18\nacc +19\njmp +26\nnop -51\njmp +260\nacc +17\njmp +428\nacc +6\njmp +405\nacc +22\nacc +10\nnop +471\njmp +352\nacc -6\nacc +48\nacc +7\nacc +3\njmp +57\nacc -10\nacc +16\nacc +16\nacc +43\njmp +432\nacc -5\nacc +0\nnop +339\nacc +49\njmp +17\nacc +33\nnop +166\nacc -5\njmp +392\nnop +246\nacc -7\nacc +21\nacc +30\njmp +398\nacc +36\nacc +24\nacc -15\nacc -9\njmp +114\nacc +19\njmp +11\nacc +43\nnop +182\njmp -129\nnop -29\nacc -6\nacc +2\njmp +398\njmp +78\nacc +36\njmp +393\nacc +15\nnop -11\nacc -7\nacc -9\njmp +76\nacc +0\nacc +27\njmp +25\nacc +27\nnop -54\njmp +458\nacc +3\nacc +29\nacc -4\nacc +43\njmp +413\nacc +33\nacc +13\njmp +382\njmp -83\nacc +42\nacc +24\njmp +64\nacc +23\nacc -13\nnop +110\nacc -5\njmp +114\njmp +113\nnop +112\nacc +26\njmp -133\njmp -12\njmp +1\njmp +330\nacc +25\nacc -1\nacc +30\nacc +42\njmp -187\njmp +1\nacc +20\nacc +35\nacc +36\njmp -125\njmp +165\nacc +28\nacc -17\nacc -12\njmp +1\njmp -120\nnop +1\nacc +2\nacc +26\njmp +398\nacc +20\nacc -1\njmp -127\nacc +36\nacc +14\njmp +1\njmp +331\nacc +50\nacc +1\nacc -10\nnop +159\njmp -83\njmp +374\nacc +17\njmp +372\nacc +44\nnop -39\njmp +228\nacc +17\njmp +74\nacc +16\nacc +33\nacc -2\njmp +152\njmp +29\nacc +8\nacc +27\nnop +59\njmp -32\nacc +28\njmp -227\nnop -35\njmp -168\nacc +13\nnop +390\njmp -204\nacc +16\nacc +44\njmp -230\njmp +25\nacc +30\njmp +383\nacc -11\nacc +38\nacc +11\njmp +341\nacc +35\nacc +46\nacc -1\njmp +94\nacc -4\nacc +12\njmp +111\njmp +133\nnop +283\nacc +13\nacc +37\njmp +74\nnop -218\njmp -178\nacc +46\nacc +25\nacc -5\njmp -174\nacc +28\nacc +39\nacc +36\nacc +22\njmp -172\nacc +19\njmp -250\nnop +62\nacc +44\nnop +347\nacc +40\njmp +345\nacc -3\nacc -13\nacc -11\njmp +56\njmp -180\nacc +17\nacc -4\nacc +46\nnop -165\njmp +321\nacc -4\njmp +1\nacc +9\nacc -12\njmp -155\nacc +5\njmp -96\nacc +0\nacc -2\nacc +38\njmp +67\nacc -4\nnop -283\nacc +28\njmp +324\nacc -9\nacc +43\nacc -1\nacc +9\njmp -290\nacc +3\nacc +22\nnop +84\nacc -17\njmp -210\nacc +7\njmp -260\nnop -232\nnop +87\nacc +43\nacc +36\njmp +96\njmp +238\nacc +13\nacc -14\nacc +32\nacc +11\njmp -146\nacc +13\nacc +37\nacc -10\njmp +187\nacc +49\nacc +15\njmp -234\njmp -328\njmp -136\njmp +143\njmp +1\nacc +27\nacc +22\njmp +1\njmp -5\nacc +30\nnop -7\nacc -6\njmp -71\nacc -17\nacc +15\njmp -52\njmp -126\nacc -4\njmp +151\njmp +52\nnop -86\nacc +25\njmp +187\nnop -22\njmp -219\nacc +33\nnop -120\nacc +0\njmp +215\nacc +46\nacc +38\njmp +1\njmp -262\njmp +157\nacc -15\nacc +48\nacc +39\nacc +10\njmp -137\nacc +47\nacc +50\njmp -324\nnop +214\nacc +39\njmp -178\nacc +49\nacc -10\njmp -268\njmp +50\nacc -14\nnop -100\njmp +20\nacc +45\nacc -12\nacc -4\njmp -208\nacc -19\njmp -340\nacc +36\nnop -358\nacc +5\njmp -348\nacc +47\nnop -18\nacc -12\njmp -131\nacc +19\nacc +10\nacc +19\nacc +31\njmp -164\nnop +162\nnop -260\njmp +146\nacc +32\nacc -1\nnop -14\njmp -192\nacc +3\nacc +31\nnop -185\njmp -208\njmp -69\nacc +43\nacc +43\njmp -68\nacc -16\nacc +5\nacc -9\njmp +126\nacc +33\nacc +2\nacc +34\nacc -9\njmp -16\nacc +34\nacc -19\njmp -266\nnop +135\nnop -389\nacc +33\njmp -195\nacc +48\njmp +1\nacc -12\njmp +143\nnop -317\nacc -14\nnop -127\nacc +32\njmp -372\nacc +24\nnop -41\nnop -42\njmp -344\nacc +23\nnop +117\nnop +92\nacc +42\njmp +143\nacc +48\nacc -6\nnop -272\nacc -13\njmp -379\nacc -2\nacc +44\nacc +9\njmp -369\nacc +6\nacc +25\nacc +34\njmp -301\nnop -227\nacc +43\njmp -141\nacc +12\nacc +41\nacc +17\nacc -11\njmp +29\njmp -121\nacc +6\nacc +7\nacc +7\njmp +131\nnop +144\nnop -142\nacc -13\nacc -18\njmp +149\nacc +14\nacc +49\nacc +25\nacc -17\njmp -9\nacc +26\nacc -4\njmp -230\nacc -18\nacc +36\nacc +27\nnop -142\njmp +21\nacc +34\nnop +54\njmp -476\nacc +10\njmp -174\nnop -354\nacc +1\njmp -324\nacc +40\njmp +94\nacc -12\njmp -136\nnop -454\nacc -14\njmp +116\nacc +12\nacc -1\nnop -453\njmp -241\njmp -479\nacc -19\njmp -87\nacc +27\nacc +48\nacc +0\njmp -476\nacc +16\nacc +46\njmp -534\nacc +0\njmp -344\nacc +0\nacc +28\njmp +10\njmp -248\nnop -186\njmp +1\nacc +26\njmp -153\nacc +14\nacc -8\nnop -416\njmp -91\njmp -409\njmp -326\nacc +2\nacc +8\nacc -18\nacc +33\njmp -468\njmp -175\nacc -7\nacc +45\njmp -18\njmp -375\nacc -8\njmp +28\nacc -16\nnop -38\nacc +37\nacc +48\njmp -343\nacc +10\nacc +26\nacc -9\nacc -16\njmp -348\nacc +37\njmp -453\nacc -2\nacc +27\nacc +17\nacc +28\njmp -406\nacc +25\nacc +24\nacc +44\nacc +44\njmp -532\nacc +10\njmp -531\nacc +39\nacc +40\njmp -284\nacc +19\nacc +3\nnop -533\nacc -3\njmp -162\nnop -438\nacc -5\njmp -114\nacc +45\nacc +1\nacc +28\nacc +9\njmp -550\njmp -222\njmp -106\nacc -7\nnop -263\nnop -375\njmp -381\nacc -4\nnop -223\njmp -171\njmp -465\nacc -2\nnop -562\njmp -190\nacc +40\njmp -4\nacc +30\nacc +21\njmp -435\nacc +1\nacc +10\njmp +1\njmp -157\nacc -7\nacc +18\nacc -3\nacc +24\njmp -113\nacc +21\njmp -339\nacc +34\njmp -563\nacc +27\njmp -589\njmp -61\nacc +35\nacc +50\nacc +8\njmp -553\nacc +48\nacc -15\nacc +29\nacc +24\njmp +1"
)

func TestInterpreter_Execute(t *testing.T) {
	ins, err := ParseInstructions(set1)
	if err != nil {
		t.Fatal(err)
	}

	ip := Interpreter{
		Ins: ins,
	}

	acc, err := ip.Execute()
	if err != errInfiniteLoop {
		t.Fatal(err)
	}

	if acc != 1723 {
		t.Fatal(acc)
	}

	ip.Reset()
	fmt.Println(ip.FixLoop())
}

func TestInterpreter_ExecuteTest(t *testing.T) {
	ins, err := ParseInstructions(set0)
	if err != nil {
		t.Fatal(err)
	}

	ip := Interpreter{
		Ins: ins,
	}

	acc, err := ip.Execute()
	if err != errInfiniteLoop {
		t.Fatal(err)
	}

	if acc != 5 {
		t.Fatal(acc)
	}

	ip.Reset()
	fmt.Println(ip.FixLoop())
}
