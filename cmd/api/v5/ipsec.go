package v5

import (
	"github.com/luscis/openlan/cmd/api"
	"github.com/luscis/openlan/pkg/schema"
	"github.com/urfave/cli/v2"
)

type IPSec struct {
	Cmd
}

func (o IPSec) Commands(app *api.App) {
	tunnel := IPSecTunnel{}
	app.Command(&cli.Command{
		Name:  "ipsec",
		Usage: "IPSec configuration",
		Subcommands: []*cli.Command{
			tunnel.Commands(),
		},
	})
}

type IPSecTunnel struct {
	Cmd
}

func (o IPSecTunnel) Url(prefix string) string {
	return prefix + "/api/network/ipsec/tunnel"
}

func (o IPSecTunnel) Add(c *cli.Context) error {
	output := &schema.IPSecTunnel{
		Right:     c.String("remote"),
		Secret:    c.String("secret"),
		Transport: c.String("transport"),
		LeftId:    c.String("localid"),
		RightId:   c.String("remoteid"),
		LeftPort:  c.Int("localport"),
		RightPort: c.Int("remoteport"),
	}
	url := o.Url(c.String("url"))
	clt := o.NewHttp(c.String("token"))
	if err := clt.PostJSON(url, output, nil); err != nil {
		return err
	}
	return nil
}

func (o IPSecTunnel) Remove(c *cli.Context) error {
	output := &schema.IPSecTunnel{
		Right:     c.String("remote"),
		Transport: c.String("transport"),
	}
	url := o.Url(c.String("url"))
	clt := o.NewHttp(c.String("token"))
	if err := clt.DeleteJSON(url, output, nil); err != nil {
		return err
	}
	return nil
}

func (o IPSecTunnel) Tmpl() string {
	return `# total {{ len . }}
{{ps -15 "Right"}} {{ps -15 "Transport"}} {{ps -15 "Secret"}} {{ps -15 "Port"}} {{ps -15 "Connection"}} 
{{- range . }}
{{ps -15 .Right}} {{ps -15 .Transport }} {{ps -15 .Secret}} {{.LeftPort}}-{{.RightPort}} {{.LeftId}}-{{.RightId}}
{{- end }}
`
}

func (o IPSecTunnel) List(c *cli.Context) error {
	url := o.Url(c.String("url"))
	clt := o.NewHttp(c.String("token"))
	var items []schema.IPSecTunnel
	if err := clt.GetJSON(url, &items); err != nil {
		return err
	}
	return o.Out(items, c.String("format"), o.Tmpl())
}

func (o IPSecTunnel) Commands() *cli.Command {
	return &cli.Command{
		Name:    "tunnel",
		Aliases: []string{"tun"},
		Usage:   "IPSec Tunnel configuration",
		Subcommands: []*cli.Command{
			{
				Name:  "add",
				Usage: "Add a tunnel for the network",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "remote", Required: true},
					&cli.StringFlag{Name: "remoteid"},
					&cli.IntFlag{Name: "remoteport"},
					&cli.StringFlag{Name: "transport", Required: true},
					&cli.StringFlag{Name: "secret", Required: true},
					&cli.StringFlag{Name: "localid"},
					&cli.IntFlag{Name: "localport"},
				},
				Action: o.Add,
			},
			{
				Name:    "remove",
				Usage:   "Remove a tunnel from the network",
				Aliases: []string{"rm"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "remote", Required: true},
					&cli.StringFlag{Name: "transport", Required: true},
				},
				Action: o.Remove,
			},
			{
				Name:    "list",
				Usage:   "Display all tunnel of the network",
				Aliases: []string{"ls"},
				Flags:   []cli.Flag{},
				Action:  o.List,
			},
		},
	}
}
