package q1

func CalculateDiscount(currentPurchase float64, purchaseHistory []float64) (float64, error) {
	var desconto float64
	var valortotal float64

	for _, x := range purchaseHistory {
		valortotal += x

	}

	if currentPurchase <= 0 {
		fmt.Print("valor da compra inválido")
		return 0, nil
	}

	if len(purchaseHistory) == 0 {
		desconto = currentPurchase * 0.1
	} else if valortotal > 1000 {
		desconto = currentPurchase * 0.1
	} else if purchaseHistory > 500 {
		desconto = currentPurchase * 0.05
	} else {
		desconto = currentPurchase * 0.02
	}
	if valortotal/float64(len(purchaseHistory)) > 1000 {
		desconto = currentPurchase * 0.2
	}
	return 0, nil
}
