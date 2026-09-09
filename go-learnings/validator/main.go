package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

// tag = `json:"key" validate:"asd,asd,asd" proto:"key"`
// space seperated & inside the value coma seperated
type Person struct {
	Name    string `json:"name" validate:"required"`
	Persona string `json:"persona" validate:"required,len=10"`                 //for multiple validation use ","
	Address string `json:"address" validate:"required,min=10,max=20,alphanum"` //min=10, max=20 and should be alpha-numeric
	Email   string `json:"email" validate:"email"`
	Age     int    `json:"age" validate:"required,gte=18,neq=100"` //min, max, gt, gte, lt, lte, eq, ne(not equal) -> works for int, float both

	Colour     string `json:"colour" validate:"oneof=green purple"`
	Nominees   string `json:"nominees" validate:"oneof='shiva chandra' 'surya teja'"` //oneof which has values with space
	ColourCase string `json:"colour_case" validate:"oneofci=green purple"`

	Number   int     `json:"number" validate:"oneof=5 7 9"`
	FloatVal float64 `json:"float_val" validate:"oneof=5.1 7.1 9.1"`
}

/*
   strings:
       alpha
       alphanum
       boolean -> make sure that "true/false" or "TRUE/FALSE" -> basically strconv.ParseBool it will check
       number -> checks if its int, float
       json -> check if its json
       oneof
       oneofci -> case in-sensitive
       max -> max len of string
       min
       len -> len of string
       eq
       ne -> not equal

   int/float:
       eq
       ne
       min
       max
       len
       oneof
*/

type Config struct {
	Email    string `json:"email" validate:"required,email"`
	EmailBkp string `json:"email_bkp" validate:"email"` //here, "required" is not present still it validates for email and fails
	Name     string `json:"name" validate:"max=10"`
	Address  string `json:"address" validate:"required"` //this will make that it does not have default value
	Age      int    `json:"age" validate:"required"`     //this will make sure that it does not have defuault value
}

func main() {
	// config := Config{Email: "shiva@gmail.com"}
	config := Config{}
	configValidate(config)
	print(config)
}

func print(v any) {
	b, _ := json.MarshalIndent(v, " ", "  ")
	fmt.Println(string(b))
}

func configValidate(config Config) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(config)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				// fmt.Println(e.Namespace())
				// fmt.Println(e.Field())
				// fmt.Println(e.StructNamespace())
				// fmt.Println(e.StructField())
				// fmt.Println(e.Tag())
				// fmt.Println(e.ActualTag())
				// fmt.Println(e.Kind())
				// fmt.Println(e.Type())
				// fmt.Println(e.Value())
				// fmt.Println(e.Param())
				// fmt.Println()
				log.Println("config validation error", "err", e)
			}
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
}
