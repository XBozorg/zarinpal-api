package zarinpal

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	ValidatePaymentRequest             func(req PaymentRequest) error
	ValidateGatewayResponse            func(resp GatewayResponse) error
	ValidatePaymentVerificationRequest func(req PaymentVerificationRequest) error
)

func statusCheck(str string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != "OK" && s != "NOK" {
			return errors.New("invalid status")
		}
		return nil
	}
}

func ValidatePayment() ValidatePaymentRequest {
	return func(req PaymentRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.MerchantID, validation.Required, validation.Length(36, 36)),
			validation.Field(&req.Amount, validation.Required, validation.Min(1000)),
			validation.Field(&req.Description, validation.Required),
			validation.Field(&req.CallbackURL, validation.Required, is.URL),
			validation.Field(&req.Metadata, validation.Map(
				validation.Key("mobile", validation.Length(11, 11)),
				validation.Key("email", is.Email),
			)),
		)
	}
}

func ValidateGateway() ValidateGatewayResponse {
	return func(req GatewayResponse) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.Status, validation.Required, validation.By(statusCheck(req.Status))),
			validation.Field(&req.Authority, validation.Required, is.Digit, validation.Length(36, 36)),
		)
	}
}

func ValidatePaymentVerification() ValidatePaymentVerificationRequest {
	return func(req PaymentVerificationRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.MerchantID, validation.Required, validation.Length(36, 36)),
			validation.Field(&req.Amount, validation.Required, validation.Min(1000)),
			validation.Field(&req.Authority, validation.Required, is.Digit, validation.Length(36, 36)),
		)
	}
}
