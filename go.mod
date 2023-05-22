module puppy

go 1.19

require github.com/go-sql-driver/mysql v1.7.0 // direct

require (
	github.com/denisenkom/go-mssqldb v0.12.3 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/microsoft/go-mssqldb v0.21.0 // indirect
	golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// Uitleg Go.Mod:Een go.mod bestand is een bestand dat wordt gebruikt door Go-modules om afhankelijkheden te beheren en te versiebeheer
// van Go-pakketten en hun bijbehorende modules. Het go.mod bestand bevat informatie over de naam van het project, de vereiste versie
// van Go, de modules waar het project van afhankelijk is en hun versies. Dit bestand wordt gegenereerd en bijgewerkt door de go tool
// wanneer u Go-modules in uw project gebruikt.
