package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")
var ErrInvalidBase = errors.New("invalid base")

func FromCodon(codon string) (string, error) {
	var protein string

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCG", "UCU", "UCC", "UCA":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		return protein, ErrStop
	default:
		return protein, ErrInvalidBase
	}

	return protein, nil
}

func FromRNA(rna string) ([]string, error) {
	var aminoAcids []string

	for i := 0; i < len(rna); i += 3 {
		aminoAcid, err := FromCodon(rna[i : i+3])

		if err == ErrStop {
			break
		}

		if err != nil {
			return aminoAcids, err
		}

		aminoAcids = append(aminoAcids, aminoAcid)
	}

	return aminoAcids, nil
}
