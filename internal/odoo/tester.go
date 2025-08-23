package odoo

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/imotif-tools/pkg/text"
)

type Tester struct {
	args []string
}

func NewTester(args []string) *Tester {
	return &Tester{
		args: args,
	}
}

func (t *Tester) RunTest() error {
	parser := text.NewParser(t.args)
	addon, err := parser.Parse("addon name is required")
	if err != nil {
		return err
	}
	if err := t.runDockerTest(addon); err != nil {
		return err
	}
	return nil
}

func (t *Tester) runDockerTest(addons string) error {
	addonList := parseAddons(addons)

	addonInstall := strings.Join(addonList, ",")
	addonCov := buildCovArgs(addonList)
	addonTestPaths := buildTestPaths(addonList)

	script := fmt.Sprintf(`OPTIONS="bash -lc 'odoo -c /etc/odoo/odoo.conf -d odoo_test -i %s --stop-after-init && pytest -c /mnt/imbase/pytest.ini --odoo-database=odoo_test --odoo-addons-path=/mnt/imbase/addons,/mnt/imbase/additional-addons %s --cov-config=/mnt/imbase/.coveragerc --cov-report=term-missing %s -v'" docker compose -f docker-compose.local-test.yml up --build --abort-on-container-exit tests`, addonInstall, addonCov, addonTestPaths)

	cmd := exec.Command("sh", "-c", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func parseAddons(addons string) []string {
	items := strings.Split(addons, ",")
	var result []string
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func buildCovArgs(addons []string) string {
	var covArgs []string
	for _, addon := range addons {
		arg := fmt.Sprintf("--cov=/mnt/imbase/addons/%s", addon)
		covArgs = append(covArgs, arg)
	}
	return strings.Join(covArgs, " ")
}

func buildTestPaths(addons []string) string {
	var paths []string
	for _, addon := range addons {
		paths = append(paths, fmt.Sprintf("/mnt/imbase/addons/%s/tests", addon))
	}
	return strings.Join(paths, " ")
}
