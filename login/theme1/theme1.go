package theme1

import "github.com/GoAdminGroup/go-admin/template"

type Theme1 struct{
	*template.BaseComponent
}

func (*Theme1) GetAssetList() []string {
	return AssetsList
}

func (*Theme1) GetAsset(name string) ([]byte, error) {
	return Asset(name)
}

func (*Theme1) GetHTML() string {
	return List["login"]
}
