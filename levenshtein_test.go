package chainletter_test

import (
	"strings"
	"testing"

	"github.com/joshuanario/chainletter"
)

func Test_ComputeLevenshteinDistanceForSameDocument(t *testing.T) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 0 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForSameDocument(t *testing.T) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 1.0 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinDistanceForModifiedDocument(t *testing.T) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	b[0] = '!'
	b[1] = '!'
	b[2] = '!'
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 3 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForModifiedDocument(t *testing.T) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	b[0] = '!'
	b[1] = '!'
	b[2] = '!'
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 0.9990586758707248 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinDistanceForDifferentDocument(t *testing.T) {
	a := []rune(document)
	b := []rune("Love is pain.  The heart is a battleground.")
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 3145 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForDifferentDocument(t *testing.T) {
	a := []rune(document)
	b := []rune("Love is pain.  The heart is a battleground.")
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 0.013178537809852526 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinDistanceForEmptyDocument1(t *testing.T) {
	a := []rune(document)
	b := make([]rune, 0)
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 3187 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForEmptyDocument1(t *testing.T) {
	a := []rune(document)
	b := make([]rune, 0)
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 0.0 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinDistanceForEmptyDocument2(t *testing.T) {
	a := make([]rune, 0)
	b := []rune(document)
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 3187 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForEmptyDocument2(t *testing.T) {
	a := make([]rune, 0)
	b := []rune(document)
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 0.0 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinDistanceForEmptyDocument3(t *testing.T) {
	a := make([]rune, 0)
	b := make([]rune, 0)
	out := chainletter.ComputeLevenshteinDistance(a, b)
	if out != 0 {
		t.Fail()
	}
}

func Test_ComputeLevenshteinRatioForEmptyDocument3(t *testing.T) {
	a := make([]rune, 0)
	b := make([]rune, 0)
	out := chainletter.ComputeLevenshteinRatio(a, b)
	if out != 1.0 {
		t.Fail()
	}
}

func BenchmarkComputeLevenshteinDistance(benchmark *testing.B) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	b[0] = '!'
	b[1] = '!'
	b[2] = '!'
	for n := 0; n < benchmark.N; n++ {
		chainletter.ComputeLevenshteinDistance(a, b)
	}
}

func BenchmarkComputeLevenshteinRatio(benchmark *testing.B) {
	a := []rune(document)
	b := []rune(strings.Clone(document))
	b[0] = '!'
	b[1] = '!'
	b[2] = '!'
	for n := 0; n < benchmark.N; n++ {
		chainletter.ComputeLevenshteinRatio(a, b)
	}
}

const document = `Drink UP! gulp, gulp, gulp, swallow, gulp..........................


         (ŻŻŻŻŻŻŻŻŻ)
          |       |
          |       |
          |       |
          |       |
         /         \
        /           \
       /             \
      /               \
     /^ v ^ v ^ v ^ v ^\
    /    ^ v     v ^    \
   /     v    ^    v     \
  /  v  ^      v    ^   v \
 (_________________________)
aaaaaaaaahhhhhhhhhhhhhhhhhhhhhhhh!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

You just drank an interactive love potion. The spell will take place in 15 seconds. As you scroll down the page, think VERY hard about the person you like.


   > 

   > 

   > 15.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 14......

   > 

   > 

   > 

   > 

   > 

   > 

   > 13....

   > 

   > 

   > 

   > 

   > 

   > 

   > 12...

   > 

   > 

   > 

   > 

   > 

   > 

   > 11.......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 10.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 9.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 8.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 7.......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 6.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 5.....

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 4......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 3.......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 2......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   > 1......

   > 

   > 

   > 

   > 

   > 

   > 

   > 

   >			You will be loved!!!!!! 

   > 

   > 

You are now under the love spell. It will only work depending on how many times you bless others with the interactive love potion.

Send it to 0 people -- the spell will work backwards and you will be cursed with bad luck in love !!

Send it to 1-5 people-- the person you thought about will become friends with you!!

Send it to 6-10 people--your crush will become VERY interested in you!!

Send it to 11-15 people--your crush will ask you out!!

Send it to 16-20 people-- you will make out with your crush!!!

HAPPY LOVING!!!!!

______________________________________

* Warning * if you do not pass this on, something bad or worse will happen to you :

CASE 1: Take Heather Dicks of Bufor, MN . She was in love with a man named Dennis Samson . Heather got this same exact letter and did not pass it on because she thought it was stupid. Dennis was her true love. Two days later Dennis died in a car accident.

You must send this on in 1 hour of reading this to 10 different people . If you do this you will receive luck in love. The person you are most attracted to will return your feelings. This is not a joke!!!!!!! You have read the warnings You must send on.

CASE 2: Eric Man sent this letter 45 minutes after receiving it. Not even 4 hours later he was walking along the street when he ran into Ann Heart, his secret crush for 5 years, Ann came up to him and told him of her passionate crush on him that she has had for 2 years. Eric and Ann are still married happy as ever.`
