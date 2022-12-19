package validate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/diag"
)

// uniqueItemsValidator validates that an Attribute's list items have unique values.
type uniqueItemsValidator struct{}

// Description describes the validation in plain text formatting.
func (v uniqueItemsValidator) Description(_ context.Context) string {
	return "list items must have unique values"
}

// MarkdownDescription describes the validation in Markdown formatting.
func (v uniqueItemsValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

// Validate performs the validation.
func (v uniqueItemsValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	for i1, val1 := range request.ConfigValue.Elements() {
		val1, err := val1.ToTerraformValue(ctx)

		if err != nil {
			response.Diagnostics.Append(diag.NewUnableToObtainValueAttributeError(
				request.Path,
				err,
			))

			return
		}

		if !val1.IsFullyKnown() {
			continue
		}

		for i2, val2 := range request.ConfigValue.Elements() {
			if i2 == i1 {
				continue
			}

			val2, err := val2.ToTerraformValue(ctx)

			if err != nil {
				response.Diagnostics.Append(diag.NewUnableToObtainValueAttributeError(
					request.Path,
					err,
				))

				return
			}

			if !val2.IsFullyKnown() {
				continue
			}

			if val1.Equal(val2) {
				response.Diagnostics.AddAttributeError(
					request.Path,
					"Duplicate value",
					"duplicate values",
				)

				return
			}
		}
	}
}

// UniqueItems returns a new unique items validator.
func UniqueItems() validator.List {
	return uniqueItemsValidator{}
}
