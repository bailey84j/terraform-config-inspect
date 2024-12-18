
# Module `testdata/variable-validation`

## Input Variables
* `variable_with_no_validation` (default `""`): This variable has no validation
* `variable_with_one_validation` (default `""`): This variable has one validation  
  Validation error messages
  * `var.variable_with_one_validation must be empty or 10 characters long.`
  Validation conditions
  * `(length(var.variable_with_one_validation) == 0 || length(var.variable_with_one_validation) == 10)`
* `variable_with_two_validations` (required): This variable has two validations  
  Validation error messages
  * `var.variable_with_two_validations must be 10 characters long.`
  * `var.variable_with_two_validations must start with 'magic'.`
  Validation conditions
  * `(length(var.variable_with_two_validations) == 10)`
  * `(startswith(var.variable_with_two_validations, "magic"))`

