package flags

import (
	"github.com/iancoleman/strcase"
	"github.com/spf13/pflag"
	"strconv"
	"strings"
)

func ReadStringFlag(name string) string {
	name = strings.ReplaceAll(strcase.ToSnake(name), "_", "-")
	return pflag.Lookup(name).Value.String()
}

func ReadStringOptionalFlag(name string) *string {
	name = strings.ReplaceAll(strcase.ToSnake(name), "_", "-")
	if v := pflag.Lookup(name).Value.String(); len(v) > 0 {
		return &v
	}
	return nil
}

func ReadInt64Flag(name string) int64 {
	str := ReadStringFlag(name)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func ReadInt64OptionalFlag(name string) *int64 {
	str := ReadStringOptionalFlag(name)
	if str != nil {
		i, _ := strconv.ParseInt(*str, 10, 64)
		return &i
	}
	return nil
}

func ReadBooleanFlag(name string) bool {
	str := ReadStringFlag(name)
	i, _ := strconv.ParseBool(str)
	return i
}

func ReadBooleanOptionalFlag(name string) *bool {
	str := ReadStringOptionalFlag(name)
	if str != nil {
		i, _ := strconv.ParseBool(*str)
		return &i
	}
	return nil
}

func ReadStringArrayFlag(name string) []string {
	str := ReadStringOptionalFlag(name)
	if str != nil {
		return []string{*str}
	}
	return nil
}

func ReadEnumArrayFlag[T ~string](name string) []T {
	str := ReadStringOptionalFlag(name)
	if str != nil {
		var s interface{} = *str
		if v, ok := s.(T); ok {
			return []T{v} //TODO
		}
	}
	return nil
}

func ReadMapStringFlag(name string) map[string]string {
	return nil //TODO
}

func ReadMapStringArrayFlag(name string) map[string][]string {
	return nil //TODO
}
