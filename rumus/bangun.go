package rumus

type Bangun struct {
	Panjang, Lebar, Tinggi float64	
}

func (bng Bangun) Luas() float64 {
	p := bng.Panjang
	l := bng.Lebar

	luas := p * l

	return luas

}

func (bng Bangun) Keliling() float64 {
	p := bng.Panjang
	l := bng.Lebar

	kll := 2*p + 2*l

	return kll
}

func (bng Bangun) Volume() float64 {
	p := bng.Panjang
	l := bng.Lebar
	t := bng.Tinggi

	volum := p * l * t

	return volum
}
