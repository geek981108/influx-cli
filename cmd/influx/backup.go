package main

import (
	"errors"

	"github.com/influxdata/influx-cli/v2/clients/backup"
	br "github.com/influxdata/influx-cli/v2/internal/backup_restore"
	"github.com/influxdata/influx-cli/v2/pkg/cli/middleware"
	"github.com/urfave/cli"
)

func newBackupCmd() cli.Command {
	var params backup.Params
	// Default to gzipping local files.
	params.Compression = br.GzipCompression

	return cli.Command{
		Name:  "backup",
		Usage: "Backup database",
		Description: `Backs up InfluxDB to a directory

Examples:
	# backup all data
	influx backup /path/to/backup
`,
		ArgsUsage: "path",
		Before:    middleware.WithBeforeFns(withCli(), withApi(true)),
		Flags: append(
			commonFlagsNoPrint(),
			&cli.StringFlag{
				Name:        "org-id",
				Usage:       "The ID of the organization",
				EnvVar:      "INFLUX_ORG_ID",
				Destination: &params.OrgID,
			},
			&cli.StringFlag{
				Name:        "org, o",
				Usage:       "The name of the organization",
				EnvVar:      "INFLUX_ORG",
				Destination: &params.Org,
			},
			&cli.StringFlag{
				Name:        "bucket-id",
				Usage:       "The ID of the bucket to backup",
				Destination: &params.BucketID,
			},
			&cli.StringFlag{
				Name:        "bucket, b",
				Usage:       "The name of the bucket to backup",
				Destination: &params.Bucket,
			},
			&cli.GenericFlag{
				Name:  "compression",
				Usage: "Compression to use for local backup files, either 'none' or 'gzip'",
				Value: &params.Compression,
			},
		),
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() != 1 {
				return errors.New("backup path must be specified as a single positional argument")
			}
			params.Path = ctx.Args().Get(0)

			api := getAPI(ctx)
			client := backup.Client{
				CLI:       getCLI(ctx),
				BackupApi: api.BackupApi.OnlyOSS(),
				HealthApi: api.HealthApi.OnlyOSS(),
			}
			return client.Backup(getContext(ctx), &params)
		},
	}
}
