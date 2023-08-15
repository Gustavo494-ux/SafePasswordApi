package routines

import (
	"os/exec"
	"strings"
)

// ListModulesToUpgrade : lista os múdlos que precisma ser atualizados
func ListModulesToUpgrade() {
	cmd := "go list -m -u all"
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		// log.Error().Err(err).Msgf("Erro ao tentar analizar o código para gerar uma documentação do swagger.")
	}

	title := "Lista dos pacotes para upgrade"
	title = strings.Replace(strings.TrimSpace(title), " ", "-", -1)
	// content := string(out)
	// log.Info().Msgf("\n\n%v%v \n\n%v%v \n\n", logger.Colors["info"], title, logger.Colors["normal"], content)
	// _ = bucket.Write(fmt.Sprintf("./data/environment-go/%v.txt", title), content)
}
