package shared

// You shouldn't modify `InitAll` and `InitNoCfg`. You can only modify `InitEssential` and `InitNonEssential`.
func InitAll(path string) {
	InitCfg(path)
	InitNoCfg()
}

func InitNoCfg() {
	InitEssential()
	InitNonEssential()
}

func InitEssential() {
	InitLogger()
}

func InitNonEssential() {
	InitGin()
}
