package flags

func ReadStringFlag(name string) string {

}

func ReadStringOptionalFlag(name string) *string {

}

func ReadInt64Flag(name string) int64 {

}

func ReadInt64OptionalFlag(name string) *int64 {

}

func ReadBooleanFlag(name string) bool {

}

func ReadBooleanOptionalFlag(name string) *bool {

}

func ReadStringArrayFlag(name string) []string {

}

func ReadEnumArrayFlag[T string](name string) []T {

}

func ReadMapStringFlag(name string) map[string]string {

}

func ReadMapStringArrayFlag(name string) map[string][]string {

}
