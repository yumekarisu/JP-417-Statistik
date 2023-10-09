package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"text/template"
)

const nilaiMax = 20.0
const pilihan = 4

type PAP struct {
    Nilai, Skala9, Skala11 int
    Bersih, Skala10, Skala100 float64
    Skala5CR1, Skala5CR2, Skala5CR3 string
}

var PAPS []PAP

var tpl *template.Template 

var meanIdeal float64
var standarDeviasiIdeal float64

var nilai []int
var nilaiBersih []float64
var nilaiSkala10 []float64
var nilaiSkala100 []float64
var nilaiSkala5Cara1 []string
var nilaiSkala5Cara2 []string
var nilaiSkala5Cara3 []string
var nilaiSkala9 []int
var nilaiSkala11 []int

func init() {
 tpl = template.Must(template.ParseGlob("./*.gotpl"))
}

func main() {

    file, err := os.Open("nilai.txt")
    if err != nil {
        log.Fatalln("Error Openning File")
    }
    defer file.Close()

    data := bufio.NewScanner(file)

    for data.Scan() {
        n, err := strconv.Atoi(data.Text())
        if err != nil {
            log.Fatalln("Error reading line")
        }
        nilai = append(nilai, n)
    }
    
    meanIdeal = MI(nilaiMax)
    standarDeviasiIdeal = SDI(meanIdeal)

    for _, v := range nilai {
        n := skorBersih(v)
        nilaiBersih = append(nilaiBersih, n)
    }

    for _, v := range nilaiBersih {
        n := skala10(v)
        nilaiSkala10 = append(nilaiSkala10, n)
    }

    for _, v := range nilaiBersih {
        n := skala100(v)
        nilaiSkala100 = append(nilaiSkala100, n)
    }

    for _, v := range nilaiBersih {
        n:= skala5cara1(meanIdeal, standarDeviasiIdeal, v)
        nilaiSkala5Cara1 = append(nilaiSkala5Cara1, n)
    }

    for _, v := range nilaiBersih {
        n:= skala5cara2(meanIdeal, standarDeviasiIdeal, v)
        nilaiSkala5Cara2 = append(nilaiSkala5Cara2, n)
    }

    for _, v := range nilaiBersih {
        n:= skala5cara3(meanIdeal, standarDeviasiIdeal, v)
        nilaiSkala5Cara3 = append(nilaiSkala5Cara3, n)
    }

    for _, v := range nilaiBersih {
        n:= skala9(meanIdeal, standarDeviasiIdeal, v)
        nilaiSkala9 = append(nilaiSkala9, n)
    }

    for _, v := range nilaiBersih {
        n:= skala11(meanIdeal, standarDeviasiIdeal, v)
        nilaiSkala11 = append(nilaiSkala11, n)
    }

    for i := 0; i < 30; i++ {
        n := PAP{
            Nilai: nilai[i],
            Bersih: nilaiBersih[i],
            Skala10: nilaiSkala10[i],
            Skala100: nilaiSkala100[i],
            Skala5CR1: nilaiSkala5Cara1[i],
            Skala5CR2: nilaiSkala5Cara2[i],
            Skala5CR3: nilaiSkala5Cara3[i],
            Skala9: nilaiSkala9[i],
            Skala11: nilaiSkala11[i],
        }
        PAPS = append(PAPS, n)
    }

    err = tpl.ExecuteTemplate(os.Stdout, "index.gotpl", PAPS)
    if err != nil {
        panic(err)
    }

}

// Menghitung Skor Bersih dari Skor Kotor
func skorBersih (nilai int) (skorFinal float64) {
    n := float64(nilai)
    skor := n - ((nilaiMax - n) / (pilihan - 1))
    skorFinal = math.Ceil(skor*100)/100
    return
}

// Menghitung Nilai dengan cara skala10
func skala10 (skorBersih float64) (skorFinal float64) {
    skor := (skorBersih / nilaiMax) * 10.0
    skorFinal = math.Ceil(skor*100)/100
    return
}

// Menghitung Nilai dengan cara skala100
func skala100 (skorBersih float64) (skorFinal float64) {
    skor := (skorBersih / nilaiMax) * 100.0
    skorFinal = math.Ceil(skor*100)/100
    return
}

// Mencari Mean Ideal
func MI (nilaiMax float64) (mi float64) {
    n := 0.5 * nilaiMax
    mi = math.Ceil(n*100)/100
    return
}

// Mencari Stnadar Deviasi Ideal
func SDI (mi float64) (sdi float64) {
    n := mi / 3.0
    sdi = math.Ceil(n*100)/100
    return
}

// Menghitung Nilai dengan metode skala 5 cara 1
func skala5cara1 (mi float64, sdi float64, skorBersih float64) (skala5 string) {
    var lim1, lim2, lim3, lim4 float64
    lim1 = mi + 2*sdi
    lim2 = mi + 1*sdi
    lim3 = mi - 1*sdi
    lim4 = mi - 2*sdi
    switch {
        case skorBersih > lim1:
            skala5 = "A"
        case skorBersih > lim2:
            skala5 = "B"
        case skorBersih > lim3:
            skala5 = "C"
        case skorBersih > lim4:
            skala5 = "D"
        default:
            skala5 = "E"
    }
    return
}

// Menghitung Nilai dengan metode skala 5 cara 2
func skala5cara2 (mi float64, sdi float64, skorBersih float64) (skala5 string) {
    var lim1, lim2, lim3, lim4 float64
    lim1 = mi + 1.5*sdi
    lim2 = mi + 0.5*sdi
    lim3 = mi - 0.5*sdi
    lim4 = mi - 1.5*sdi
    switch {
        case skorBersih > lim1:
            skala5 = "A"
        case skorBersih > lim2:
            skala5 = "B"
        case skorBersih > lim3:
            skala5 = "C"
        case skorBersih > lim4:
            skala5 = "D"
        default:
            skala5 = "E"
    }
    return
}

// Menghitung Nilai dengan metode skala 5 cara 3
func skala5cara3 (mi float64, sdi float64, skorBersih float64) (skala5 string) {
    var lim1, lim2, lim3, lim4 float64
    lim1 = mi + 1.8*sdi
    lim2 = mi + 0.6*sdi
    lim3 = mi - 0.6*sdi
    lim4 = mi - 1.8*sdi
    switch {
        case skorBersih > lim1:
            skala5 = "A"
        case skorBersih > lim2:
            skala5 = "B"
        case skorBersih > lim3:
            skala5 = "C"
        case skorBersih > lim4:
            skala5 = "D"
        default:
            skala5 = "E"
    }
    return
}

// Menghitung Nilai dengan metode skala 9
func skala9 (mi float64, sdi float64, skorBersih float64) (skala9 int) {
    var lim1, lim2, lim3, lim4, lim5, lim6, lim7, lim8 float64
    lim1 = mi + 1.75*sdi
    lim2 = mi + 1.25*sdi
    lim3 = mi + 0.75*sdi
    lim4 = mi + 0.25*sdi
    lim5 = mi - 0.25*sdi
    lim6 = mi - 0.75*sdi
    lim7 = mi - 1.25*sdi
    lim8 = mi - 1.75*sdi
    switch {
        case skorBersih > lim1:
            skala9 = 9
        case skorBersih > lim2:
            skala9 = 8
        case skorBersih > lim3:
            skala9 = 7
        case skorBersih > lim4:
            skala9 = 6
        case skorBersih > lim5:
            skala9 = 5
        case skorBersih > lim6:
            skala9 = 4
        case skorBersih > lim7:
            skala9 = 3
        case skorBersih > lim8:
            skala9 = 2
        default:
            skala9 = 1
    }
    return
}

// Menghitung Nilai dengan metode skala 11
func skala11 (mi float64, sdi float64, skorBersih float64) (skala11 int) {
    var lim1, lim2, lim3, lim4, lim5, lim6, lim7, lim8, lim9, lim10 float64
    lim1 = mi + 2.75*sdi
    lim2 = mi + 1.75*sdi
    lim3 = mi + 1.25*sdi
    lim4 = mi + 0.75*sdi
    lim5 = mi + 0.25*sdi
    lim6 = mi - 0.25*sdi
    lim7 = mi - 0.75*sdi
    lim8 = mi - 1.25*sdi
    lim9 = mi - 1.75*sdi
    lim10 = mi - 2.75*sdi
    switch {
        case skorBersih > lim1:
            skala11 = 10
        case skorBersih > lim2:
            skala11 = 9
        case skorBersih > lim3:
            skala11 = 8
        case skorBersih > lim4:
            skala11 = 7
        case skorBersih > lim5:
            skala11 = 6
        case skorBersih > lim6:
            skala11 = 5
        case skorBersih > lim7:
            skala11 = 4
        case skorBersih > lim8:
            skala11 = 3
        case skorBersih > lim9:
            skala11 = 3
        case skorBersih > lim10:
            skala11 = 3
       default:
            skala11 = 0
    }
    return
}
