package database

import (
	"embed"
)

//go:embed migration/*.sql
var EmbedMigration embed.FS
