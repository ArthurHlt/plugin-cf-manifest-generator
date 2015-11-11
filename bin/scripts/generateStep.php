<?php
$name = $argv[0];
if (count($argv) < 3) {
    echo "Usage: generate_step.sh <step name> <default value>\n";
    exit(1);
}
$stepName = $argv[1];
$defaultValue = $argv[2];
$file = sprintf(__DIR__ . '/../../mg/step_%s.go', strtolower($stepName));
ob_start();
include_once __DIR__ . '/../template/template.php';
$stepFile = ob_get_contents();
ob_end_clean();

file_put_contents($file, $stepFile);
echo "Step has been generated !\n";
exit(0);