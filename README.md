# Tempus

Een terminal-gebaseerde TUI app voor het bijhouden van vrije dagen en overuren, gebouwd met [Bubble Tea](https://github.com/charmbracelet/bubbletea) en [Lip Gloss](https://github.com/charmbracelet/lipgloss).

---

## Installatie

### Via Go (aanbevolen)

Vereist [Go](https://golang.org/) 1.21 of hoger.

```bash
go install github.com/axbrunn/tempus@latest
```

Na installatie is `tempus` beschikbaar als commando in je terminal.

### Via release (geen Go vereist)

Download de laatste binary voor jouw platform op de [releases pagina](https://github.com/axbrunn/tempus/releases).

| Platform | Bestand |
|----------|---------|
| Windows | `tempus_Windows_x86_64.zip` |
| macOS (Intel) | `tempus_Darwin_x86_64.tar.gz` |
| macOS (Apple Silicon) | `tempus_Darwin_arm64.tar.gz` |
| Linux | `tempus_Linux_x86_64.tar.gz` |

Pak het archief uit en zet het `tempus` bestand in een map die in je `PATH` staat, bijvoorbeeld `/usr/local/bin` op Mac/Linux. Op Windows kun je het `.exe` bestand direct uitvoeren.

#### Windows
Maak een map aan, bijv. `C:\Tools`
Zet `tempus.exe` daarin
Voeg die map toe aan je `PATH`:

Zoek op "omgevingsvariabelen" in het startmenu
Klik op "Omgevingsvariabelen bewerken voor uw account"
Selecteer Path → klik Bewerken
Klik Nieuw en vul `C:\Tools` in.
Open een nieuwe terminal en typ `tempus`

### Vanuit broncode

```bash
git clone https://github.com/axbrunn/tempus.git
cd tempus
go install .
```

---

## Gebruik

```bash
tempus
```

De app opent met een bestandskiezer. Hier kun je een bestaand bestand selecteren of een nieuw aanmaken.

---

## Navigatie

| Toets | Actie |
|-------|-------|
| `↑` / `k` | Omhoog navigeren |
| `↓` / `j` | Omlaag navigeren |
| `enter` | Selecteren / bevestigen |
| `tab` | Wisselen tussen invoervelden |
| `n` | Nieuw bestand aanmaken |
| `esc` | Terug / annuleren |
| `q` / `ctrl+c` | Afsluiten |

---

## Functies

- **Bestandsbeheer** — Maak meerdere JSON-bestanden aan om uren per project of periode bij te houden. Bestanden worden opgeslagen in `~/tempus/`.
- **Uren opbouwen** — Registreer overuren met een omschrijving en datum.
- **Uren opnemen** — Registreer opgenomen vrije uren.
- **Overzicht** — Bekijk alle geregistreerde uren in een overzichtstabel met het huidige saldo.
- **Rapport genereren** — Exporteer alle gegevens inclusief totalen naar een CSV-bestand in je `Downloads` map.

---

## Bestandsstructuur

```
~/tempus/
  werk-2026.json
  verlof-2026.json
  ...
```

Elk JSON-bestand bevat een lijst van entries:

```json
{
  "entries": [
    {
      "date": "2026-03-28T00:00:00Z",
      "hours": 8,
      "description": "Overwerk vrijdag",
      "type": "opbouw"
    }
  ]
}
```

---

## CSV Export

Het gegenereerde CSV-bestand wordt opgeslagen als `~/Downloads/tempus-export.csv` en bevat alle entries inclusief totaalrijen:

```
date,hours,description,type
2026-03-28,8.00,Overwerk vrijdag,opbouw
2026-03-29,2.50,Verlof,opnemen

Totaal opgebouwd,8.00,,
Totaal opgenomen,2.50,,
Saldo,5.50,,
```

---

## Technologie

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) — Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) — UI componenten (textinput)
- [Cobra](https://github.com/spf13/cobra) — CLI framework
