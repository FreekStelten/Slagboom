package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() { 
	//Database connectie gemaakt/ staat als hardcode nu erin. dit is niet veilig.
	dbUser := "Admin"
	dbPass := "Fonteyn@DB"
	dbName := "slagboom_db"
	dbAddress := "127.0.0.1"

	//plate is een argument om het kenteken te ontvangen zodat klantgegevens moet opgehaald kan worden.
	//flag wordt gebruikt om om dit argument uit te lezen. als er geen argument kan gelezen worden dan wordt
	// de usage functie aangeroepen om te laten zien het programma correct gebruikt kan worden. en dan wordt de code met exit code afgesloten.
	plate := flag.String("plate", "", "er moet een kenteken opgegeven worden!")
	flag.Parse()
	if !flag.Parsed() || *plate == "" {
		flag.Usage()
		log.Println("Geen kenteken opgegeven, probeer het opnieuw.")
		logError("Geen kenteken opgegeven, probeer het opnieuw.")
		os.Exit(1)
	}

	// Create data source name (DSN)
	//connectie met database gemaakt, datasourcename gegenereerd met de gegevens van de db conn parameters.
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbAddress, dbName)

	//nieuwe DB conn te openen met de opgegeven dsn, als dit niet lukt wordt er een error geschreven naar de console en errorfile.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errMsg := fmt.Sprintf("er kan geen connecting naar de database gemaakt worden: %s", err.Error())
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
		errMsg := fmt.Sprintf("Er kan niet gepinged worden naar de database: %s", err.Error())
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
	var name, licenseplate string
	rows, err := db.Query("SELECT name,licenseplate FROM klant WHERE licenseplate = ?", *plate)
	if err != nil {
		errMsg := fmt.Sprintf("%s", err.Error())
		log.Println(errMsg)
		logError(errMsg)
		return
	}
	//Dit zorgt ervoor dat de rijen die zijn geretourneerd door de query, worden gesloten wanneer de functie is voltooid.
	defer rows.Close()

	//for rows.next itereert door elke rij die is geretourneerd door de query.
	//err = rows.Scan(&name, &licenseplate): Dit scant de waarden van de kolommen name en licenseplate in de huidige rij en slaat ze op in de variabelen name en licenseplate.
	//Als er een fout optreedt tijdens het scannen, wordt de fout in de terminal getoond en errorlogfile.
	for rows.Next() {
		err = rows.Scan(&name, &licenseplate)
		if err != nil {
			errMsg := fmt.Sprintf("%s", err.Error())
			log.Println(errMsg)
			logError(errMsg)
			return
		}
		//Dit wordt getoond in de terminal, als alles woordt doorgelopen wordt de eerste %s (string;Name) en daarna de 2e %s (stringl;licenseplate) getoond
		//dus= welkom: Name , Jouw kenteken is Kenteken.
		fmt.Printf("Welkom: %s, Jouw kenteken is %s.\n", name, licenseplate)
	}
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
