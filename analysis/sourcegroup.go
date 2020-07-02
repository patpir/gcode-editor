package analysis

type SourceGroup map[string]*Source

func (g SourceGroup) DeriveAll() {
	for key, source := range g {
		derivedKey := key + "'"
		if _, ok := g[derivedKey]; !ok {
			g[derivedKey] = Derivative(source)
		}
	}
}
