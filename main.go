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
    Nilai int
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

    for i := 0; i < 30; i++ {
        n := PAP{
            Nilai: nilai[i],
            Bersih: nilaiBersih[i],
            Skala10: nilaiSkala10[i],
            Skala100: nilaiSkala100[i],
            Skala5CR1: nilaiSkala5Cara1[i],
            Skala5CR2: nilaiSkala5Cara2[i],
            Skala5CR3: nilaiSkala5Cara3[i],
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


