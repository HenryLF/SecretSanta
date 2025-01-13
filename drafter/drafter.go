package drafter

import (
	"errors"
	"log"
)

type Result struct {
	Draft        map[Person]Person
	MaxRecursion int
}

func NewResult() Result {
	var out Result
	out.Draft = map[Person]Person{}
	out.MaxRecursion = 0
	return out

}

type outJSON struct {
	Gifter   []Person
	Receiver []Person
}

func (R Result) toJSON() outJSON {
	var out = outJSON{}
	for g, r := range R.Draft {
		out.Gifter = append(out.Gifter, g)
		out.Receiver = append(out.Receiver, r)
	}
	return out
}

type Person struct {
	Name, Family string
	Mutex        int
}

var ErrNoGifterError = errors.New("no gifters")
var ErrNoReceiverError = errors.New("no receiver")
var ErrEndRecursion = errors.New("end recursion")
var ErrUnstricModeOK = errors.New("/ unstrict mode")
var ErrRecursionFail = errors.New("recursion failed")
var ErrUnstricModeFail = errors.New("unstrict mode failed")

func Process(req Resquest) (Result, error) {
	var Gifters []Person
	var Receivers []Person
	var err error
	var out = Result{make(map[Person]Person), 0}

	Gifters, _ = shuffle(req.Gifters)
	Receivers, err = shuffle(req.Gifters)
	if err != nil {
		return Result{}, ErrNoGifterError
	}

	out, err = recursiveFuck(out, Gifters, Receivers, 0)

	if errors.Is(err, ErrRecursionFail) && req.Unstrict {
		log.Println("Unstrict Mode @ ", out.MaxRecursion)
		out, err = recursiveFuck(Result{make(map[Person]Person), 0}, Gifters, Receivers, out.MaxRecursion)
	}
	if err != nil {
		log.Print(err)
	}

	return out, nil

}

func recursiveFuck(out Result, Gifters, Receivers []Person, maxRec int) (Result, error) {
	var err error
	var validReceiver []Person
	log.Println(out)
	if len(Gifters) == 0 && len(Receivers) == 0 {
		return out, ErrEndRecursion
	}
	Gifters, err = shuffle(Gifters)
	if err != nil {
		return out, ErrNoGifterError
	}
	Receivers, err = shuffle(Receivers)
	if err != nil {
		return out, ErrNoReceiverError
	}
	out.MaxRecursion++
	A := Gifters[len(Gifters)-1]

	//UnstrictMode
	if maxRec != 0 && out.MaxRecursion > maxRec {
		return justPick(out, Gifters, Receivers)
	}

	//Not in same family
	validReceiver = filter(Receivers, func(B Person) bool {
		return B.Family != A.Family
	})
	//Not in same mutex
	validReceiver = filter(validReceiver, func(B Person) bool {
		return A.Mutex != B.Mutex || A.Mutex == 0 || B.Mutex == 0
	})
	for _, B := range validReceiver {
		//Not Same persone
		if A == B {
			continue
		}
		newGifters := remove(Gifters, A)
		newReceiver := remove(Receivers, B)
		newOut := out
		newOut.Draft[A] = B
		newOut, err = recursiveFuck(newOut, newGifters, newReceiver, maxRec)
		if err != nil {
			if errors.Is(err, ErrEndRecursion) {
				return newOut, ErrEndRecursion
			}
		}

	}
	out.Draft = map[Person]Person{}
	return out, ErrRecursionFail
}

func justPick(out Result, Gifter, Receiver []Person) (Result, error) {
	if len(Gifter) != len(Receiver) {
		return Result{}, ErrUnstricModeFail
	}
	for _, A := range Gifter {
		B, err := randomOf(Receiver)
		if err != nil {
			return out, errors.Join(ErrUnstricModeFail, err)
		}
		Receiver = remove(Receiver, B)
		out.Draft[A] = B
	}
	return out, errors.Join(ErrEndRecursion, ErrUnstricModeOK)
}
