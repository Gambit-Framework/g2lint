package sprofile

type ServerProfile struct {
	Profile struct {
		Name        string `kdl:"name"`
		Description string `kdl:"description"`
	} `kdl:"profile"`

	Server struct {
		BindHost string `kdl:"bind-host"`
		BindPort int    `kdl:"bind-port"`
	} `kdl:"server"`

	Operators struct {
		RootPassword string         `kdl:"root-password"`
		Users        map[string]any `kdl:"user,multiple"`
	} `kdl:"operators"`

	Listeners struct {
		HttpListeners map[string]any `kdl:"http,multiple"`
	} `kdl:"listeners"`
}
