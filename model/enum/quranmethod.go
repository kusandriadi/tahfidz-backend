package enum

func QuranMethodEnum() *quranMethod {
	return &quranMethod{
		SABAQ:  "Sabaq",
		MANZIL: "Manzil",
		SABAQI: "Sabaqi",
		EMPTY:  "",
	}
}

type quranMethod struct {
	SABAQ  string
	MANZIL string
	SABAQI string
	EMPTY  string
}
