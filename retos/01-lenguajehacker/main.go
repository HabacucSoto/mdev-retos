/*
 * Escribe un programa que reciba un texto y transforme lenguaje natural a
 * "lenguaje hacker" (conocido realmente como "leet" o "1337"). Este lenguaje
 *  se caracteriza por sustituir caracteres alfanuméricos.
 * - Utiliza esta tabla (https://www.gamehouse.com/blog/leet-speak-cheat-sheet/)
 *   con el alfabeto y los números en "leet".
 *   (Usa la primera opción de cada transformación. Por ejemplo "4" para la "a")
 */
package main

import (
	"fmt"
	"strings"
)

func identifyIndices(str string) []int {
	indices := []int{}
	letters := strings.Split(str, "")
	alphabet := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for _, l := range letters {
		for i, a := range alphabet {
			if l == a {
				indices = append(indices, i)
			}
		}
	}
	return indices
}

func indexLeet(str string) string {
	indices := identifyIndices(str)
	conversion := []string{}
	leet := []string{"4","I3","[",")","3","|=","&","#","1",",_|",">|","1","JVI","^/","0","|*","(_,)","I2","5","7","(_)","|/","'//","><","j","2"}
	for _, i := range indices{
		conversion = append(conversion, leet[i])
	}
	return strings.Join(conversion, "")
}

func main(){
	fmt.Println(indexLeet("hola"))
}