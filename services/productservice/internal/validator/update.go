package productvalidator

import (
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/param"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) ValidateUpdateRequest(req param.UpdateProductRequest) (map[string]string, error) {
	const op = "productvalidator.ValidateUpdateRequest"

	if err := validation.ValidateStruct(&req,
		validation.Field(&req.ID,
			validation.Required,
			validation.Min(uint(1))),

		validation.Field(&req.StoreID,
			validation.Required,
			validation.Min(uint(1))),

		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 100)),

		validation.Field(&req.Description,
			validation.Length(0, 1000)),

		validation.Field(&req.Category,
			validation.Required),

		validation.Field(&req.Price,
			validation.Required,
			validation.Min(0.0)),

		validation.Field(&req.Stock,
			validation.Required,
			validation.Min(0)),

		validation.Field(&req.SKU,
			validation.Required,
			validation.Length(3, 50),
			validation.Match(regexp.MustCompile(skuRegex)).
				Error(errmsg.ErrorMsgSKUIsNotValid)),

		validation.Field(&req.ImageURL,
			validation.Match(regexp.MustCompile(urlRegex)).
				Error(errmsg.ErrorMsgImageURLIsNotValid)),
	); err != nil {
		fieldErrors := make(map[string]string)

		errV, ok := err.(validation.Errors)
		if ok {
			for key, value := range errV {
				if value != nil {
					fieldErrors[key] = value.Error()
				}
			}
		}

		return fieldErrors, richerror.New(op).WithMessage(errmsg.ErrorMsgInvalidInput).
			WithKind(richerror.KindInvalid).
			WithMeta(map[string]interface{}{"req": req}).WithErr(err)
	}

	return nil, nil
}
