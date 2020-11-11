// Package envcfg implements functions to load tagged structs the OS env.
//
// The structure's values can be arbitrary types that are supported in this
// library.  Currently supported types include: `float64`, `int` and `string`.
//
// Example structure:
//
// The only expected tag arguments are as follows (`envcfg:"tag0,tag1,tag2"`):
//  tag0 -  The first tag item is the key name to look for in the env file.
//  tag1 -  The second tag item is an optional command like "required".
//  tag2 -  The default value for the field, if speicfied.
package envcfg

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// parseTag takes a cfg tag and returns the Key we are looking for and a boolean
// to indicate if the key is a required value or not.
func parseTag(tag string) (string, bool, string) {
	key, req, def := "", false, ""
	items := strings.Split(tag, ",")
	if len(items) >= 1 {
		key = items[0]
	}
	if len(items) >= 2 {
		req = (strings.ToLower(items[1]) == "required")
	}
	if len(items) >= 3 {
		def = items[2]
	}
	return key, req, def
}

// Load populates a config structure from the OS env.
func Load(v interface{}) error {
	val := reflect.ValueOf(v).Elem()
	for i := 0; i < val.NumField(); i++ {
		vf := val.Field(i)
		tf := val.Type().Field(i)
		key, req, def := parseTag(tf.Tag.Get("envcfg"))
		if len(key) == 0 && req {
			key = tf.Name
		}

		// Look for the key in the json map if its a valid tag.
		if len(key) > 0 {
			value := os.Getenv(key)
			if value == "" && def != "" {
				value = def
			}
			if req && value == "" {
				return fmt.Errorf("missing required field in env (%v)", key)
			}

			switch tf.Type.Name() {
			case "string":
				vf.SetString(value)
			case "int":
				iv, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				vf.SetInt(iv)
			case "float64":
				fv, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				vf.SetFloat(fv)
			default:
				return fmt.Errorf("unhandled type (%v)", tf.Type.Name())
			}
		} else if req {
			return fmt.Errorf("field (%v) missing env prop (%v)", tf.Name, key)
		}
	}

	return nil
}
