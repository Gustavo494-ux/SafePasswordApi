package routines

import (
	"fmt"
	"os/exec"
	"strings"
	// "github.com/lucasfrct/environment-go/pkg/modules/logger"
)

// var log = logger.New()

// AnalyzingCode : analiza o codigo fonte
func AnalyzingCode() {
	go func() {
		ImportAnalysis()
		SecurityAnalysis()
		HeuristicAnalysis()
		StyleAnalysis()
		DeepHeiristicalAnalysis()
		WritingAnalysisAndCorrection()
		PerformanceAnalysisAndDiagnosis()
	}()
}

// ImportAnalysis : corrige as importações e também formata o código no mesmo estilo do gofmt
func ImportAnalysis() {
	cmd := "goimports"
	// out, err := exec.Command("bash", "-c", cmd).Output()
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro nas importações do código. (goimports)")
	}
	print("GOIMPORTS Import Analysis", string(out))
}

// SecurityAnalysis : verifica segurança
func SecurityAnalysis() {
	cmd := "gosec ./..." // No windows para escrever o exec dele. exec.Command ("gosec","./...")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro de segurança no código. (gosec)")
	}

	print("GOSEC security Analysis", string(out))
}

// HeuristicAnalysis : análise eurística
func HeuristicAnalysis() {
	cmd := "go vet ."
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro ao tentar analizar o código (go vet)")
	}

	print("GO VET Heuristic Analysis", string(out))
}

// StyleAnalysis : analise de estilo
func StyleAnalysis() {
	cmd := "golint ./pkg/..."
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro do Linter na estrutura do código. (golint .)")
	}

	print("GOLINT Style Analysis", string(out))
}

// DeepHeiristicalAnalysis : analise profunda
func DeepHeiristicalAnalysis() {
	cmd := "staticcheck ."
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro erro ")
	}

	print("GO STATICCHECK Deep Heuristical Analysis", string(out))
}

// WritingAnalysisAndCorrection : análise e correçao de escrita
func WritingAnalysisAndCorrection() {
	cmd := "golangci-lint run ./..."
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro de escrita no código. (golangci-lint run ./...)")
	}

	print("GOLANGCI-LINT Writing Analysis AND Correction", string(out))
}

// PerformanceAnalysisAndDiagnosis : diagnóstico de performance
func PerformanceAnalysisAndDiagnosis() {
	cmd := "gocritic check ."
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(fmt.Errorf("erro %s: ", err.Error()))
		// log.Error().Err(err).Msgf("Erro cricito no código. (gocritic check .)")
	}

	print("GOCRITIC Performance Analysis AND Diagnosis", string(out))
}

// print : print a mensagem
func print(title, content string) {
	title = strings.Replace(strings.TrimSpace(title), " ", "-", -1)
	fmt.Printf("\n\n%v \n\n%v \n\n", title, content)
	// log.Info().Msgf("\n\n%s%v \n\n%v%s \n\n", logger.Colors["info"], title, logger.Colors["normal"], content)
	// _ = bucket.Write(fmt.Sprintf("./data/environment-go/%v.txt", title), content)
}
