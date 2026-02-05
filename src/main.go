// Pakke-definisjon (alltid 'package main' for kjørbare programmer)
package main

// Import-blokk (konseptuell)
/*
  Du trenger biblioteker for:
  - Tekstmanipulasjon (strings, fmt)
  - Fil- og OS-operasjoner (os, path/filepath, flag)
  - Google Gemini SDK (github.com/google/generative-ai-go/genai)
  - Kontekst-håndtering (context)
*/

func main() {
	// ---------------------------------------------------------
	// 1. OPPSETT OG KONFIGURASJON
	// ---------------------------------------------------------
	// Bruk 'flag'-pakken til å lese input-argumenter fra kommandolinjen:
	//   --path (hvor ligger koden?)
	//   --model (hvilken modell?)
	// Parse flaggene.

	// Hent API-nøkkel fra miljøvariabler (os.Getenv).
	// VIKTIG GO-LÆRDOM: Sjekk alltid om variabler er tomme.
	// Hvis nøkkel mangler -> log.Fatal("Mangler nøkkel!")

	// ---------------------------------------------------------
	// 2. INITIALISER GEMINI KLIENT
	// ---------------------------------------------------------
	// Opprett en 'context' (ctx := context.Background()).
	// Opprett klienten (genai.NewClient).
	// HÅNDTER FEIL: I Go sjekker vi alltid `if err != nil`.

	// VIKTIG GO-LÆRDOM: Bruk 'defer client.Close()'.
	// Dette garanterer at forbindelsen lukkes når main() er ferdig,
	// selv om programmet krasjer lenger ned.

	// ---------------------------------------------------------
	// 3. LES INN MALENE (TEMPLATES)
	// ---------------------------------------------------------
	// Les filen "templates/system_instruction.md" -> lagre i variabel.
	// Les filen "templates/output_template.md" -> lagre i variabel.
	// (Bruk os.ReadFile)

	// ---------------------------------------------------------
	// 4. SAMLE KODEBASEN (SCANNER)
	// ---------------------------------------------------------
	// Her kaller du en hjelpefunksjon (se lenger ned) som går gjennom mappene.
	// codeContext, err := scanFiles(targetDirectory)
	// Hvis feil -> Stopp programmet.

	// ---------------------------------------------------------
	// 5. BYGG PROMPTEN (SAMMENSETNING)
	// ---------------------------------------------------------
	// VIKTIG GO-LÆRDOM: Bruk 'strings.Builder' her.
	// Ikke bruk "+" for å slå sammen store tekster (det er tregt i Go).

	// Builder.Write(SystemInstruction)
	// Builder.WriteString("\n---\n")
	// Builder.Write(OutputTemplate)
	// Builder.WriteString("\n---\n")
	// Builder.WriteString("Her er koden:\n")
	// Builder.Write(CodeContext)

	// ---------------------------------------------------------
	// 6. SEND TIL AI (API KALL)
	// ---------------------------------------------------------
	// Velg modell: model := client.GenerativeModel(modelName)
	// Send data: resp, err := model.GenerateContent(ctx, fullPrompt)

	// Hent ut teksten fra svaret (ligger ofte inni resp.Candidates[0].Content).

	// ---------------------------------------------------------
	// 7. LAGRE RESULTATET
	// ---------------------------------------------------------
	// Lag mappen "docs" hvis den ikke finnes (os.MkdirAll).
	// Skriv svaret til filen "docs/AI_GENERATED_DOCS.md" (os.WriteFile).

	// Print "Suksess!" til terminalen.
}

// ---------------------------------------------------------
// HJELPEFUNKSJON: FIL-SCANNER
// ---------------------------------------------------------
func scanFiles(rootPath string) (string, error) {
	// Opprett en ny strings.Builder for å samle all koden.

	// VIKTIG GO-LÆRDOM: Bruk 'filepath.WalkDir'.
	// Den går gjennom filsystemet rekursivt (mapper inni mapper).
	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {

		// A. FILTERING AV MAPPER (STØY)
		// Hvis 'd' er en mappe:
		//   Sjekk om den heter ".git", "node_modules", "vendor", etc.
		//   Hvis ja -> return filepath.SkipDir (Go hopper over hele mappen).

		// B. FILTRERING AV FILER (HVA SKAL MED?)
		// Sjekk filendelse (filepath.Ext).
		// Er det .tf? .go? .yaml? .md?
		// Hvis nei -> return nil (gå til neste fil).

		// C. LES FILEN
		// Les innholdet (os.ReadFile).

		// D. FORMATER TIL MARKDOWN
		// Skriv til builderen:
		// "--- FILNAVN: [sti] ---\n"
		// "```[filtype]\n"
		// [filinnhold]
		// "\n```\n"

		return nil // Alt gikk bra, fortsett til neste.
	})

	return builder.String(), err
}
