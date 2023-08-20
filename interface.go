// interface methods
package main
import "fmt"

type NoiseMaker interface {
	Makesound(string)
}

// type Horn string
type Horn []string
type Whistle []string

func (h Horn) Makesound(sound string) {
	//fmt.Println(h)
	fmt.Println("play  ", sound)

}

func (w Whistle) Makesound(sound string) {
	//fmt.Println(w)
	fmt.Println("play  ", sound)

}

func Play(horn_whistle NoiseMaker, s []string) {

	//fmt.Println(horn_whistle)
	for _, sound := range s {
		horn_whistle.Makesound(sound)

	}

}
func main() {

	slcHorn := Horn{"horn1", "horn2"}

	var noisehorn NoiseMaker = slcHorn
	//var noisehorn NoiseMaker = Horn("horn")

	for _, slc := range slcHorn { //1st method
		noisehorn.Makesound(slc)
	}
	Play(noisehorn, slcHorn) // 2nd method using non interface Play method


	wh_whistle := Whistle{"whistle1"}
	var noisewhistle NoiseMaker = wh_whistle // Whistle("whistle")

	for _, wh := range wh_whistle { //1st method
		noisehorn.Makesound(wh)
	}
	Play(noisewhistle, wh_whistle) // 2nd method using non interface Play method

}
