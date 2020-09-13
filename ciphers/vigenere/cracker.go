package vigenere

import (
	"fmt"
	"sort"
	"strings"

	"github.com/khan745/cipher_cracker/language_tools"
)

type kv struct {
	Key   string
	Value float64
}

type ka struct {
	Key   string
	Value []int
}

func orderArr(data map[string][]int) []ka {
	var ss []ka
	for k, v := range data {
		ss = append(ss, ka{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return len(ss[i].Value) > len(ss[j].Value)
	})

	return ss
}

func orderdic(data map[string]float64) []kv {
	var ss []kv
	for k, v := range data {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func spacing(array []int) []int {
	var spacings []int

	for i := 1; i < len(array); i++ {
		diff := array[i] - array[i-1]
		spacings = append(spacings, diff)
		if len(spacings) > 1 {
			spacings = append(spacings, spacings[len(spacings)-2]+diff)
		}
	}
	sort.Ints(spacings[:])
	return spacings
}
func removeDuplicates(elements []int) []int {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	sort.Ints(result[:])
	return result
}

func calculateFactors(number, maxKeyLength int) []int {
	var factors []int

	if number < 2 {
		return factors
	}

	for i := 2; i < maxKeyLength+1; i++ {
		if number%i == 0 {
			factors = append(factors, i)
			otherFactor := number / i
			if otherFactor < maxKeyLength+1 && otherFactor != 1 {
				factors = append(factors, otherFactor)
			}
		}
	}

	factors = removeDuplicates(factors)
	return factors
}

func max(a []int) int {
	max := a[0]

	for _, x := range a {
		if x > max {
			max = x
		}
	}
	return max
}

func countOccurence(array []int) map[int]int {
	dic := make(map[int]int)
	for _, x := range array {
		dic[x]++
	}
	return dic
}

func getNstring(str string, n int) string {
	newString := ""
	str = strings.ReplaceAll(str, " ", "")

	for i, x := range str {
		if i != 0 && i%n == 0 {
			newString += string(x)
		}
	}
	return newString
}

func Crack(textSecret string, realQuadgrams map[string]float64, alphabetNormalProbability []language_tools.Alphabet, alphabets []string) {
	//TODO: crack viniger cipher
	// 1. find key length
	// 2. frequency analysis
	// textSecret = "PPQCAXQVEKGYBNKMAZUYBNGBALJONITSZMJYIMVRAGVOHTVRAUCTKSGDDWUOXITLAZUVAVVRAZCVKBQPIWPOU"

	keyLenths := []int{3, 4, 5, 6, 7}
	maxKeyLength := max(keyLenths)
	posibleKeyLengths := make(map[int]int)

	for _, k := range keyLenths {
		history := language_tools.RepetativeStrings(textSecret, k)
		historyOrder := orderArr(history)
		allValues := []int{}
		fmt.Println("key length: ", k)
		for _, x := range historyOrder[:] {
			if len(x.Value) > 2 {
				x.Value = spacing(x.Value)
				allValues = append(allValues, x.Value...)
			}
		}
		allValues = removeDuplicates(allValues)
		fmt.Println("Spacing: ", allValues)

		allfactors := []int{}
		for _, x := range allValues {
			factors := calculateFactors(x, maxKeyLength)
			allfactors = append(allfactors, factors...)
		}
		// allfactors = removeDuplicates(allfactors)
		// fmt.Println("Factors: ", allfactors)
		occurences := countOccurence(allfactors)
		// fmt.Println("Occur: ", occurences)

		for k, v := range occurences {
			posibleKeyLengths[k] += v
		}
	}

	fmt.Println(posibleKeyLengths)
	// TODO
	// for each key length calculate key

	t := getNstring(textSecret, 5)
	ic := language_tools.IndexOfCoincidence(t)
	fmt.Println(ic)

	fmt.Println(t)
	pt := language_tools.CalculateProbability(t)

	for k, v := range pt {
		fmt.Println(string(k), float64(v)/float64(len(t))*100)
	}
	alphabetNormalProb := language_tools.ReadFiles("./alphabets", "csv")

	for _, x := range alphabetNormalProb {
		fmt.Println(x)
	}
}