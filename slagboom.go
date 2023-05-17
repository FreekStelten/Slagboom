package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// err := GetConfig("config.yaml"): Deze regel roept de functie GetConfig aan en geeft het argument
// "config.yaml" mee, naam van het configuratiebestand dat moet worden opgehaald.
// Het resultaat van de functie wordt toegewezen aan de variabele err.
func main() {
	err := GetConfig("config.yaml")
	if err != nil {
		errMsg := fmt.Sprintf(" configfile not found: %s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	//plate is een argument om het kenteken te ontvangen zodat klantgegevens moet opgehaald kan worden.
	//flag wordt gebruikt om om dit argument uit te lezen. als er geen argument kan gelezen worden dan wordt
	// de usage functie aangeroepen om te laten zien het programma correct gebruikt kan worden. en dan wordt de code met exit code afgesloten.
	plate := flag.String("plate", "", "Code1: er moet een geldig kenteken opgegeven worden!")
	flag.Parse()
	if !flag.Parsed() || *plate == "" {
		flag.Usage()
		log.Println("Geen kenteken opgegeven, Probeer het opnieuw.")
		logError("Geen kenteken opgegeven, probeer het opnieuw.")
		os.Exit(1)
	}

	// Create data source name (DSN)
	//connectie met database gemaakt, datasourcename gegenereerd met de gegevens van de db conn parameters.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", Configuration.Database.DbUser, Configuration.Database.DbPass, Configuration.Database.DbAddress, Configuration.Database.DbName)

	//nieuwe DB conn te openen met de opgegeven dsn, als dit niet lukt wordt er een error geschreven naar de console en errorfile.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("code2: db gegevens niet vindbaar: %s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	//defer dbclose wordt gebruikt om de conn met de db af te sluiten na de functie als er een error voorkomt.
	defer db.Close()

	//Hier wordt de functie Ping() op de database uitgevoerd om te controleren of de connectie geldig is. Als de functie een error teruggeeft,
	//wordt deze opgeslagen in een variable genaamd errMsg, waarna deze wordt geprint naar de console en naar een logbestand via de functie
	//logError(). Vervolgens wordt de functie gestopt via de return statement.
	err = db.Ping()
	if err != nil {
		errMsg := fmt.Sprintf("code3: Er kan geen connectie gemaakt worden met de database: %s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}

	// Connection successful, als de vorige stappen succesvol doorlopen zijn geeft die in de terminal aan"connected to database".
	// worden er 2 variablen aangemaakt genaamd: name en licenseplate. die ook allebei strings zijn.
	// er wordt een query toegepast uit de bovengenoemde db, om alle rijen te selecteren in de tabel klant waarvan de waarde van het veld licenseplate gelijk is aan *plate.
	//Het resultaat wordt opgeslagen in de variabele rows. Als er een fout optreedt tijdens het uitvoeren van de query, wordt de fout opgeslagen in de variabele err.
	//bij if err!=... controleert of er een fout is opgetreden bij uitvoeren query. zoja wordt de fout opgeslagen in de variabele errMsg, dat wordt laten zien in de terminal en de errologsfile.
	fmt.Println("Connected to database!")

	var name, licenseplate, begindatum, vertrekdatum string
	rows, err := db.Query("SELECT name, licenseplate, begindatum, vertrekdatum FROM klant WHERE licenseplate = ?", *plate)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	//Dit zorgt ervoor dat de rijen die zijn geretourneerd door de query, worden gesloten wanneer de functie is voltooid.
	defer rows.Close()

	//for rows.next itereert door elke rij die is geretourneerd door de query.
	//dit is het geval als er een kenteken wordt opgegeven en die niet in de database staat.
	if !rows.Next() {
		log.Println("voer een geldig kenteken op, deze is niet op uw reservering gezet, probeer het opnieuw.")
		logError("voer een geldig kenteken op, deze is niet op uw reservering gezet, probeer het opnieuw.")
		os.Exit(1)
	}

	//err = rows.Scan(&name, &licenseplate): Dit scant de waarden van de kolommen name en licenseplate in
	//de huidige rij en slaat ze op in de variabelen name en licenseplate.
	//Als er een fout optreedt tijdens het scannen, wordt de fout in de terminal getoond en errorlogfile.
	err = rows.Scan(&name, &licenseplate, &vertrekdatum, &begindatum)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}

	//Dit wordt getoond in de terminal, als alles woordt doorgelopen wordt de eerste %s (string;Name) en daarna de 2e %s (stringl;licenseplate) getoond
	//dus= welkom: Name , Jouw kenteken is Kenteken.
	fmt.Printf("Welkom: %s, Jouw kenteken is %s, je begindatum is %s, je vertrekdatum is %s. \n", name, licenseplate, vertrekdatum, begindatum)
}

// logerror slaat een foutmleding op in een bestand genaamd errorlogs.txt. als daar een fout mee is dan wordt de fout naar de terminal gestuurd en wordt de functie verlaten.
func logError(errMsg string) {
	file, err := os.OpenFile("errorlogs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Failed to open errorlogs.txt:", err.Error())
		return
	}
	defer file.Close()

	//stelt uitvoer van logpakket in op het geopende bestand, zodat de logboekmelding wordt geschreven naar het bestand ipv naar de standaartuitvoering.
	log.SetOutput(file)
	//print de errormessage uit als de errorlogfile niet gevonden kan worden.
	log.Println(errMsg)
}

// Deze code leest een YAML-configuratiebestand in en slaat de inhoud ervan op in een Configuration-variabele van het type Config.
// De functie GetConfig leest het bestand en gebruikt yaml.Unmarshal om de YAML-gegevens om te zetten naar de Configuration-variabele.
// Als er een fout optreedt, wordt de fout geretourneerd.
var Configuration Config

type Config struct {
	Database struct {
		DbUser    string `yaml:"dbUser"`
		DbPass    string `yaml:"dbPass"`
		DbName    string `yaml:"dbName"`
		DbAddress string `yaml:"dbAddress"`
	} `yaml:"database"`
}

func GetConfig(fileLocation string) error {

	// Read the file
	data, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Unmarshal the YAML data into the struct
	err = yaml.Unmarshal(data, &Configuration)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
