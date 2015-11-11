package mg
import (
"github.com/cloudfoundry/cli/plugin"
"github.com/ArthurHlt/plugin-cf-manifest-generator/mg/manifest"
"fmt"
"bufio"
"os"
)

type Step<?php echo ucfirst($stepName); ?> struct {
Step
}
const DEFAULT_<?php echo strtoupper($stepName); ?>= "<?php echo $defaultValue ?>"
func NewStep<?php echo ucfirst($stepName); ?>(cliConnection plugin.CliConnection, appManifest *manifest.Manifest) *Step<?php echo ucfirst($stepName); ?>{
step<?php echo ucfirst($stepName); ?> := new(Step<?php echo ucfirst($stepName); ?>)
step<?php echo ucfirst($stepName); ?>.appManifest = appManifest
step<?php echo ucfirst($stepName); ?>.cliConnection = cliConnection
return step<?php echo ucfirst($stepName); ?>
}

func (s *Step<?php echo ucfirst($stepName); ?>) Run() error {
reader := bufio.NewReader(os.Stdin)
for true {
fmt.Print(fmt.Sprintf("How much instances do you want <%s> ? ", DEFAULT_<?php echo strtoupper($stepName); ?>))
<?php echo $stepName; ?>Bytes, _, err := reader.ReadLine()
if err != nil {
return err
}
<?php echo $stepName; ?> := string(<?php echo $stepName; ?>Bytes)
if (<?php echo $stepName; ?> == "") {
<?php echo $stepName; ?> = DEFAULT_<?php echo strtoupper($stepName); ?>
}
<?php echo $stepName; ?>, err = s.get<?php echo ucfirst($stepName); ?>(<?php echo $stepName; ?>)
if err != nil {
showError(err)
continue
}
s.appManifest.<?php echo ucfirst($stepName); ?>(session.AppName, <?php echo $stepName; ?>)
break
}
return nil
}
func (s *Step<?php echo ucfirst($stepName); ?>) get<?php echo ucfirst($stepName); ?>(<?php echo $stepName; ?> string) (string, error) {
return <?php echo $stepName; ?>, nil
}