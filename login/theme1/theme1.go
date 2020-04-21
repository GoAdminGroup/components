package theme1

type Theme1 struct{}

func (*Theme1) GetAssetList() []string {
	return AssetsList
}

func (*Theme1) GetAsset(name string) ([]byte, error) {
	return Asset(name)
}

func (*Theme1) GetHTML() string {
	return List["login"]
}
