package main

// Import-blokk
/*
  Biblioteker:
  - "strings", "fmt", "log"
  - "os", "path/filepath", "flag"
  - "context"
  - "github.com/google/generative-ai-go/genai"
  - "google.golang.org/api/option"
*/

func main() {
	// ---------------------------------------------------------
	// 1. OPPSETT OG KONFIGURASJON
	// ---------------------------------------------------------
	// flag.String("path", ...) -> Hvor er app-koden?
	// flag.String("model", ...)
	// flag.Parse()

	// Hent API Key fra os.Getenv.
	// Hvis tom -> log.Fatal.

	// ---------------------------------------------------------
	// 2. INITIALISER GEMINI KLIENT
	// ---------------------------------------------------------
	// context.Background()
	// genai.NewClient(...)
	// defer client.Close()

	// ---------------------------------------------------------
	// 3. LES INN MALENE (TEMPLATES)
	// ---------------------------------------------------------
	// os.ReadFile("templates/system_instruction.md")
	// os.ReadFile("templates/output_template.md")

	// ---------------------------------------------------------
	// 3.5. LES DESIGNBESLUTNINGER (NY! - ADR)
	// ---------------------------------------------------------
	// Vi ser etter mappen "docs/adr" inne i app-repoet.
	// adrPath := filepath.Join(targetDirectory, "docs", "adr")

	// Kaller ny hjelpefunksjon (se bunnen av filen).
	// adrContext := collectDesignDecisions(adrPath)

	// ---------------------------------------------------------
	// 4. SAMLE KODEBASEN (SCANNER)
	// ---------------------------------------------------------
	// codeContext, err := scanFiles(targetDirectory)
	// Sjekk feil.

	// ---------------------------------------------------------
	// 5. BYGG PROMPTEN (OPPDATERT SANDWICH)
	// ---------------------------------------------------------
	// var sb strings.Builder

	// A. System Instruction (Rollen)
	// sb.Write(systemInstruction)
	// sb.WriteString("\n---\n")

	// B. Template (Formatet)
	// sb.WriteString("Please use this template:\n")
	// sb.Write(outputTemplate)
	// sb.WriteString("\n---\n")

	// C. Design Decisions (Kontekst) - NYTT!
	// Her primer vi AI-en med hvorfor ting er som de er.
	// sb.WriteString("Here are the Architecture Decision Records (ADR):\n")
	// sb.WriteString(adrContext)
	// sb.WriteString("\n---\n")

	// D. Koden (Fakta)
	// sb.WriteString("Here is the source code:\n")
	// sb.WriteString(codeContext)

	// fullPrompt := sb.String()

	// ---------------------------------------------------------
	// 6. SEND TIL AI (API KALL)
	// ---------------------------------------------------------
	// model := client.GenerativeModel(modelName)
	// resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	// Hent tekst fra responsen.

	// ---------------------------------------------------------
	// 7. LAGRE RESULTATET
	// ---------------------------------------------------------
	// os.MkdirAll(filepath.Join(targetDirectory, "docs"), ...)
	// os.WriteFile(..., data, 0644)
	// Print "Done".
}

// ---------------------------------------------------------
// HJELPEFUNKSJON 1: FIL-SCANNER (Rekursiv)
// ---------------------------------------------------------
func scanFiles(rootPath string) (string, error) {
	// strings.Builder...
	// filepath.WalkDir(rootPath, func(...) {
	// 1. Ignorer .git, node_modules, etc (return filepath.SkipDir)
	// 2. Sjekk filendelser (.go, .tf, .yaml)
	// 3. Les fil
	// 4. Formater: "--- FILE: [navn] ---\n [innhold]"
	// })
	// return builder.String()
}

// ---------------------------------------------------------
// HJELPEFUNKSJON 2: LES DESIGNBESLUTNINGER (Flat liste)
// ---------------------------------------------------------
func collectDesignDecisions(adrPath string) string {
	// 1. SJEKK OM MAPPEN FINNES
	// GO-LÆRDOM: Bruk os.Stat() for å sjekke eksistens.
	// if _, err := os.Stat(adrPath); os.IsNotExist(err) {
	//     return "No ADRs found." // Helt ok, vi bare returnerer tomt.
	// }

	// 2. LES MAPPEN (IKKE REKURSIVT)
	// GO-LÆRDOM: Bruk os.ReadDir() når du bare vil ha filene i én mappe.
	// entries, err := os.ReadDir(adrPath)

	// strings.Builder...

	// 3. LOOP GJENNOM FILENE
	// for _, entry := range entries {
	//     if entry.IsDir() { continue } // Hopp over undermapper
	//     if !strings.HasSuffix(entry.Name(), ".md") { continue } // Kun markdown

	//     fullPath := filepath.Join(adrPath, entry.Name())
	//     content, _ := os.ReadFile(fullPath)

	//     sb.WriteString("\n--- DECISION RECORD: " + entry.Name() + " ---\n")
	//     sb.Write(content)
	// }

	// return sb.String()
}
