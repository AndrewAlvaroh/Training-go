package main

import "fmt"

type Ban interface {
    Jenis() string
}

type BanKaret struct{}

func (bk BanKaret) Jenis() string {
    return "Karet"
}

type BanKayu struct{}

func (by BanKayu) Jenis() string {
    return "Kayu"
}

type BanBesi struct{}

func (bb BanBesi) Jenis() string {
    return "Besi"
}

type Pintu struct {
    Sifat string
}

func (p *Pintu) Ketuk() string {
    if p.Sifat == "kanan" {
        return "Knock"
    }
    return "Beep"
}

func (p *Pintu) Buka() string {
    if p.Sifat == "kanan" {
        return "Beep"
    }
    return "Knock"
}

type Mobil struct {
    Roda [4]Ban
    PintuKiri, PintuKanan Pintu
}

func main() {
    mobil := Mobil{
        Roda: [4]Ban{BanKaret{}, BanKaret{}, BanKaret{}, BanKaret{}},
        PintuKiri: Pintu{Sifat: "kiri"},
        PintuKanan: Pintu{Sifat: "kanan"},
    }

    fmt.Println("Pintu Kanan - Ketuk:", mobil.PintuKanan.Ketuk())
    fmt.Println("Pintu Kanan - Buka:", mobil.PintuKanan.Buka())

    fmt.Println("Pintu Kiri - Ketuk:", mobil.PintuKiri.Ketuk())
    fmt.Println("Pintu Kiri - Buka:", mobil.PintuKiri.Buka())

    // Penggantian roda
    mobil.Roda[0] = BanBesi{}
    mobil.Roda[1] = BanBesi{}

    fmt.Println("Setelah penggantian roda:")
    fmt.Println("Roda 1 jenis:", mobil.Roda[0].Jenis())
    fmt.Println("Roda 2 jenis:", mobil.Roda[1].Jenis())
}
