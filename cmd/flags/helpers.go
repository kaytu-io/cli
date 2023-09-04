package flags

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func ReadStringFlag(cmd *cobra.Command, name string) string {
	name = strings.ReplaceAll(strcase.ToSnake(name), "_", "-")
	if cmd.Flags().Lookup(name) == nil {
		fmt.Println("cant find", name)
	}
	return cmd.Flags().Lookup(name).Value.String()
}

func ReadStringOptionalFlag(cmd *cobra.Command, name string) *string {
	name = strings.ReplaceAll(strcase.ToSnake(name), "_", "-")
	if v := cmd.Flags().Lookup(name).Value.String(); len(v) > 0 && (cmd.Flags().Lookup(name).Changed == true) {
		return &v
	}
	return nil
}

func ReadInt64Flag(cmd *cobra.Command, name string) int64 {
	str := ReadStringFlag(cmd, name)
	i, _ := strconv.ParseInt(str, 10, 64)
	return i
}

func ReadInt64OptionalFlag(cmd *cobra.Command, name string) *int64 {
	str := ReadStringOptionalFlag(cmd, name)
	if str != nil {
		i, _ := strconv.ParseInt(*str, 10, 64)
		return &i
	}
	return nil
}

func ReadTimeOptionalFlag(cmd *cobra.Command, name string) *int64 {
	str := ReadStringOptionalFlag(cmd, name)
	if str != nil {
		i, err := strconv.ParseInt(*str, 10, 64)
		if err != nil {
			layout := "2006-01-02"
			t, err := time.Parse(layout, *str)
			if err != nil {
				panic(err)
			}
			if name == "EndTime" {
				t = t.AddDate(0, 0, 1).Add((-1) * time.Duration(1) * time.Second)
			}
			epochTime := t.Unix()
			return &epochTime
		} else {
			return &i
		}
	}
	return nil
}

func ReadTimeFlag(cmd *cobra.Command, name string) int64 {
	str := ReadStringFlag(cmd, name)
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		layout := "2006-01-02"
		t, err := time.Parse(layout, str)
		if err != nil {
			panic(err)
		}
		if name == "EndTime" {
			t = t.AddDate(0, 0, 1).Add((-1) * time.Duration(1) * time.Second)
		}
		epochTime := t.Unix()
		return epochTime
	} else {
		return i
	}
}

func ReadBooleanFlag(cmd *cobra.Command, name string) bool {
	str := ReadStringFlag(cmd, name)
	i, _ := strconv.ParseBool(str)
	return i
}

func ReadBooleanOptionalFlag(cmd *cobra.Command, name string) *bool {
	str := ReadStringOptionalFlag(cmd, name)
	if str != nil {
		i, _ := strconv.ParseBool(*str)
		return &i
	}
	return nil
}

func ReadStringArrayFlag(cmd *cobra.Command, name string) []string {
	str := ReadStringOptionalFlag(cmd, name)
	if str != nil {
		str := *str
		str = strings.ReplaceAll(str, "[", "")
		str = strings.ReplaceAll(str, "]", "")
		if str == "" {
			return nil
		}
		return strings.Split(str, ",")
	}
	return nil
}

func ReadEnumArrayFlag[T ~string](cmd *cobra.Command, name string) []T {
	str := ReadStringOptionalFlag(cmd, name)
	if str != nil {
		var s interface{} = *str
		if v, ok := s.(T); ok {
			return []T{v} //TODO
		}
	}
	return nil
}

func ReadMapStringFlag(cmd *cobra.Command, name string) map[string]string {
	return nil //TODO
}

func ReadMapStringArrayFlag(cmd *cobra.Command, name string) map[string][]string {
	return nil //TODO
}

func Name(n string) string {
	return strings.ReplaceAll(strcase.ToSnake(n), "_", "-")
}

func ReadIntArrayFlag(cmd *cobra.Command, name string) []int64 {
	strArr := ReadStringArrayFlag(cmd, name)
	intArr := make([]int64, len(strArr))

	for i, s := range strArr {
		num, _ := strconv.ParseInt(s, 10, 64)
		intArr[i] = num
	}

	return intArr
}
