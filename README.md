# Slagboom
POP uitwerking
Deze applicatie is een prototype om een poort te openen op een vakantiepark. De app neemt een kentekenplaat als invoer en beslist op basis van reserveringsgegevens of de auto met de opgegeven kentekenplaat toegang heeft tot de parkeerplaats.
Bij het betreden van het vakantiepark kan de app worden gebruikt om de poort te openen door het invoeren van het kenteken van de auto. Vervolgens zal de app controleren of de auto is gereserveerd en of deze toegang heeft tot de parkeerplaats.
## Exit Codes
0. Success
1. er is geen kenteken opgegeven.
2. er is geen database gevonden met de naam slagboom_db
3. er kan geen connecting naar de database gemaakt worden

## Log Configuration
Om ervoor te zorgen dat de applicatie fouten op een effectieve manier kan opsporen en oplossen, moet er een logbestand worden gemaakt. Dit logbestand moet worden genoemd errorlogs.txt en moet zich bevinden in de hoofdmap van de applicatie. Hierdoor kan de applicatie alle relevante gebeurtenissen en fouten bijhouden en kunnen ontwikkelaars deze informatie gebruiken om eventuele problemen op te lossen.

**File**: The file is the path to the file to write the logs
C:\fontys\jaar 1\S2\Code\Slagboom\Slagboom\errorlogs.txt

**LogLevel**: 
in de logfile worden alle foutmeldingen die voor kunnen komen getoond. er staat de datum en tijd bij dus dan is goed te achterhalen
wanneer er iets is voorgekomen.
