package client

import (
	"fmt"
	"os"
	"github.com/mmd-afegbua/web-crawler/crawl"
	"github.com/antoniou/go-crawler/sitemap"
	"github.com/antoniou/go-crawler/util"
	"github.com/urfave/cli"
)

// Client is a command line client

type Client struct {
	app *cli.App
}

// Run function starts the command line client
func (client *Client) Run(arguments []string) error {
	return client.app.Run(arguments)
}

// New is a constructor for Client
func New() (client *Client) {
	client = new(Client)
	app := cli.NewApp()
	app.Name = "go-crawler"
	app.Usage = "Crawl a site"
	ap.UsageText = "crawl [options url"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "o",
			Value: "result.out",
			Usage: "Output file",
		},
		cli.BooFlag{
			Name: "verbose",
			Usage: "Verbose mode",
		},
	}

	app.Action = func(c *cli.Context) error {
		_ = util.Logger(c.Bool("verbose"))
		if len(c.Args()) == 0 {
			client.app.Commands[0].Run(c)
			return fmt.Errorf("Needs at least one argument")
		}
		return client.crawl(c)
	}

	client.app = app
	return client
}

// crawl function initiates the crawling steps
func (client *Client) crawl(c *cli.Context) error {
	args := c.Args()

	seedURL, err := util.NormalizeStringUrl(args[0])
	if err != nil {
		return err
	}

	crawler := crawl.NewAsHTTPCrawler(seedURL)
	stmp, err := crawler.Crawl()
	if err != nil {
		return err
	}

	outfile := c.String("o")
	return client.export(outfile, stmp)
}

// export sitemap stmp to new file outfile

func (client *Client) export(outfile string, stmp sitemap.Sitemapper) error {
	f, err := os.Create(outfile)
	if err != nil {
		return err
	}

	err = sitemap.NewExporter(f).Export(stmp)
	if err != nil {
		return err
	}
	fmt.Printf("The sitemap is exported to %s\n", outfile)
	return nil
}
