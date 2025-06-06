package valueobject

type PaymentMethodType string 


const (
	PaymentMethodCard PaymentMethodType = "CARD"
	PaymentMethodPayPal PaymentMethodType = "PAYPAL"
)

type PaymentMethod struct{
	Typ PaymentMethodType
	Details string
}

func NewPaymentMethod(typ PaymentMethodType, detail string){
	
}