package strand

func ToRNA(dna string) string {
	var rna string

	dnaToRna := getDnaToRnaTable()

	for _, n := range dna {
		rna += dnaToRna[n]
	}

	return rna
}

func getDnaToRnaTable() map[rune]string {
	return map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
}
